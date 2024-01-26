package afreecachat

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// NewClient 함수는 Client 구조체를
// 초기화하여 생성한다
func NewClient(token Token) *Client {
	return &Client{
		Token:     token,
		read:      make(chan []byte, 1024),
		handshake: make([][]byte, 2),
	}
}

// Connect 메서드는 채팅 서버 연결에 필요한
// 과정을 수행한다.
func (c *Client) Connect() error {
	err := c.createWebsocket()
	if err != nil {
		return err
	}

	return c.processSocket()
}

// setHandshake 메서드는 채팅 서버 연결에 필요한
// Handshake 과정을 수행한다.
// 이 때 2번째 단계를 수행했을 경우
// onConnect 콜백으로 값을 전달한다.
func (c *Client) setHandshake(svc int) error {
	err := c.socket.WriteMessage(websocket.BinaryMessage, c.handshake[svc-1])
	if err != nil {
		return err
	}

	if svc == 2 {
		if c.onConnect != nil {
			c.onConnect(true)
		}
	}

	return nil
}

// setLoginHandshake 메서드는 채팅 서버 연결에
// 필요한 Login Handshake 과정을 수행한다.
func (c *Client) setLoginHandshke() error {
	if c.Token.Flag == "" || c.Token.PdBoxTicket == "" {
		return errors.New("need token value")
	}

	var packet []string
	packet = append(packet, "\f", c.Token.PdBoxTicket, "\f", "\f", c.Token.Flag, "\f")

	return c.setHandshakeData(1, packet)
}

// setJoinHandshake 메서드는 채팅 서버 연결에
// 필요한 Join Handshake 과정을 수행한다.
func (c *Client) setJoinHandshake() error {
	if c.Token.ChatRoom == "" || c.Token.FanTicket == "" {
		return errors.New("need token value")
	}

	infoPacket := append(c.SetLogHandshake(DefaultLog()), c.SetInfoHandshake(DefaultInfo())...)
	var packet []string
	packet = append(
		packet,
		"\f",
		c.Token.ChatRoom,
		"\f",
		c.Token.FanTicket,
		"0",
		"\f",
		"",
		"\f",
		string(infoPacket),
		"\f",
	)

	return c.setHandshakeData(2, packet)
}

// setHandshakeData 메서드는 아프리카TV 채팅 서버에 연결할 때
// 필요한 데이터를 생성하는 과정을 수행한다.
func (c *Client) setHandshakeData(svc int, packet []string) error {
	bodyBuf := makeBuffer(packet)
	headerBuf := makeHeader(svc, len(bodyBuf), 0)
	p := append(headerBuf, bodyBuf...)

	c.handshake[svc-1] = p

	return nil
}

// SetLogHandshake 메서드는 Handshake 과정 중
// 필요한 Log 데이터를 가공한다.
func (c *Client) SetLogHandshake(log Log) []byte {
	result := append([]byte("log"), 17)
	result = append(result, c.setLogValue(log)...)
	result = append(result, 18)

	return result
}

// SetInfoHandshake 메서드는 Handshake 과정 중
// 필요한 Info 데이터를 가공한다.
func (c *Client) SetInfoHandshake(info Info) []byte {
	var result []byte
	infoValue := reflect.ValueOf(info)

	for i := 0; i < infoValue.NumField(); i++ {
		field := infoValue.Field(i)
		if !field.IsZero() {
			k := strings.ToLower(infoValue.Type().Field(i).Tag.Get("json"))
			v := fmt.Sprintf("%v", field.Interface())
			kv := append([]byte(k), 17)
			kv = append(kv, []byte(v)...)
			kv = append(kv, 18)
			result = append(result, kv...)
		}
	}

	return result
}

// setLogValue 메서드는 Handshake 과정 중
// Log 구조체를 []byte 로 변환한다.
func (c *Client) setLogValue(log Log) []byte {
	var result []byte
	logValue := reflect.ValueOf(log)

	for i := 0; i < logValue.NumField(); i++ {
		field := logValue.Field(i)
		if !field.IsZero() {
			k := strings.ToLower(logValue.Type().Field(i).Tag.Get("json"))
			v := fmt.Sprintf("%v", field.Interface())
			kv := append([]byte{6}, []byte(k)...)
			kv = append(kv, 61)
			kv = append(kv, []byte(v)...)
			kv = append(kv, 6)
			result = append(result, kv...)
		}
	}

	return append([]byte{6, 38, 61}, result...)
}

// processSocket 메서드는 웹소켓으로
// 들어오는 데이터를 처리한다.
func (c *Client) processSocket() error {
	defer c.socket.Close() // 함수가 종료되기 전에 소켓을 닫는다.

	wg := sync.WaitGroup{}
	wg.Add(1)

	go c.reader(&wg)
	c.pingpong()

	err := c.startParser(&wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

// reader 메서드는 웹소켓으로 들어오는 데이터를
// read 필드로 전달한다.
func (c *Client) reader(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			c.read <- []byte(fmt.Sprintf("error: %s", err.Error()))
			return
		}

		c.read <- msg
	}
}

// startParser 메서드는 Handshake 첫번째 과정을
// 수행하고, read 필드로 전달된 데이터를
// 처리하여 콜백 함수로 전달한다.
func (c *Client) startParser(wg *sync.WaitGroup) error {
	err := c.setLoginHandshke()
	if err != nil {
		wg.Done()
		return err
	}

	err = c.setHandshake(1)
	if err != nil {
		wg.Done()
		return err
	}

	for msg := range c.read {
		if strings.HasPrefix(string(msg), "error: ") {
			wg.Done()
			return errors.New(string(msg))
		}

		svc := getServiceCode(msg)
		switch svc {
		case 1: // Login, need JOIN handshake
			err = c.setJoinHandshake()
			if err != nil {
				return err
			}

			err = c.setHandshake(2) // join handshake
			if err != nil {
				return err
			}
		case 4: // 입장/퇴장
			m := c.parseUserJoin(msg)
			if c.onUserLists != nil {
				c.onUserLists(m)
			}
		case 5: // Chat
			m := c.parseChatMessage(msg)
			if c.onChatMessage != nil {
				c.onChatMessage(m)
			}
		case 18: // 별풍선
			m := c.parseBalloon(msg)
			if c.onBalloon != nil {
				c.onBalloon(m)
			}
		case 87: // 애드벌룬
			m := c.parseAdballoon(msg)
			if c.onAdballoon != nil {
				c.onAdballoon(m)
			}
		case 91, 93: // 신규 구독 / 연속 구독
			m := c.parseSubscription(msg, svc)
			if c.onSubscription != nil {
				c.onSubscription(m)
			}

			// case 109: // OGQ 이모티콘
			// case 12: // 플래그 변경 (ex, 열혈달았을 때)
		}
	}

	return nil
}

// pingpong 메서드는 매 1분마다 ping 데이터를
// 전송한다
func (c *Client) pingpong() {
	t := time.NewTicker(1 * time.Minute)
	go func() {
		for range t.C {
			p := "1b093030303030303030303130300c"
			h, _ := hex.DecodeString(p)
			c.socket.WriteMessage(websocket.BinaryMessage, h)
		}
	}()
}

// createWebsocket 메서드는 아프리카TV 채팅서버에
// 소켓을 연결한다.
func (c *Client) createWebsocket() error {
	if c.socket != nil {
		// 이미 존재하는 소켓이라면 반환
		return nil
	}

	dialer := websocket.DefaultDialer
	header := http.Header{}
	header.Set("Sec-WebSocket-Protocol", "chat")

	var err error
	c.socket, _, err = dialer.Dial(c.SocketAddress, header)
	if err != nil {
		return err
	}

	return nil
}
