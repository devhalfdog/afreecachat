package afreecachat

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	dataUrl = "https://live.afreecatv.com/afreeca/player_live_api.php?bj_id=%s"
)

func (c *Client) setSocketData() error {
	data := url.Values{}
	data.Set("bid", c.Token.BJID)
	data.Set("player_type", "html5")

	resp, err := http.Post(
		fmt.Sprintf(dataUrl, c.Token.BJID),
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := gjson.GetBytes(body, "CHANNEL")
	domain := result.Get("CHDOMAIN").String()
	port := result.Get("CHPT").Int() + 1

	c.socketAddress = fmt.Sprintf("wss://%s:%d/Websocket", domain, port)
	c.Token.chatRoom = result.Get("CHATNO").String()

	return nil
}
