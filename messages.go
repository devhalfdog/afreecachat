package afreecachat

import (
	"strconv"
	"strings"
)

func (c *Client) parseJoinChannel(message []byte) bool {
	msg := strings.Split(string(message), "\f")
	return msg[1] != "비밀번호가 틀렸습니다."
}

// parseUserJoin 메서드는 전달된 데이터의
// 서비스 코드가 4일 때 이 데이터를 이용해
// User 구조체로 초기화하고 반환한다.
func (c *Client) parseUserJoin(message []byte) []UserList {
	// userlist := make([]User, 0)
	msg := strings.Split(string(message), "\f")
	var userlist []UserList

	if len(msg) > 10 {
		// 여러 명의 유저일 경우
		// 최초 실행이기 때문에 무조건 입장했다고 표시함.
		ul := parseMultiUserList(msg)
		return ul
	} else {
		// 단일 유저일 경우
		userlist = append(userlist, parseSingleUserList(msg))
	}

	return userlist
}

// parseChatMessage 메서드는 전달된 데이터의
// 서비스 코드가 5일 때 이 데이터를 이용해
// ChatMessage 구조체로 초기화하고 반환한다.
func (c *Client) parseChatMessage(message []byte) ChatMessage {
	msg := strings.Split(string(message), "\f")
	flags := strings.Split(msg[7], "|")
	userFlag := setFlag(flags)
	subMonth, _ := strconv.Atoi(msg[8])

	cm := ChatMessage{
		User: User{
			ID:             removeParentheses(strings.TrimSpace(msg[2])),
			Name:           strings.TrimSpace(msg[6]),
			SubscribeMonth: subMonth,
			Flag:           userFlag,
		},
		Message: strings.TrimSpace(msg[1]),
	}

	return cm
}

// parseBallon 메서드는 전달된 데이터의
// 서비스 코드가 18일 때 이 데이터를 이용해
// Ballon 구조체로 초기화하고 반환한다.
func (c *Client) parseBalloon(message []byte) Balloon {
	msg := strings.Split(string(message), "\f")
	balloonCount, _ := strconv.Atoi(msg[4])
	user := User{
		ID:   msg[2],
		Name: msg[3],
	}
	balloon := Balloon{
		User:  user,
		Count: balloonCount,
	}

	return balloon
}

// parseAdballoon 메서드는 전달된 데이터의
// 서비스 코드가 87일 때 이 데이터를 이용해
// Adballoon 구조체로 초기화하고 반환한다.
func (c *Client) parseAdballoon(message []byte) Adballoon {
	msg := strings.Split(string(message), "\f")
	adballoonCount, _ := strconv.Atoi(msg[10])
	user := User{
		ID:   msg[3],
		Name: msg[4],
	}
	adballoon := Adballoon{
		User:  user,
		Count: adballoonCount,
	}

	return adballoon
}

// parseSubscription 메서드는 전달된 데이터의
// 서비스 코드가 91, 93일 때 이 데이터를 이용해
// Subscription 구조체로 초기화하고 반환한다.
func (c *Client) parseSubscription(message []byte, svc int) Subscription {
	msg := strings.Split(string(message), "\f")
	var user User
	count := 1
	if svc == 91 {
		// 신규 구독
		user.ID = removeParentheses(msg[3])
		user.Name = msg[4]
	} else if svc == 93 {
		// 연속 구독
		user.ID = removeParentheses(msg[2])
		user.Name = msg[3]
		count, _ = strconv.Atoi(msg[4])
	}

	subscription := Subscription{
		User:  user,
		Count: count,
	}

	return subscription
}
