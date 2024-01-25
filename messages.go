package afreecachat

import (
	"strconv"
	"strings"
)

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
	flag, _ := strconv.Atoi(flags[0])
	cm := ChatMessage{
		User: User{
			ID:   removeParentheses(msg[2]),
			Name: msg[6],
			Flag: getFlag(flag),
		},
		Message: msg[1],
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
	// "\x1b\t009100004600\f8632\fmygomiee\fdaegoguryeo(5)\f대풍순대\f15\f" // 신규 구독?
	// "\x1b\t009300004100\fmygomiee\fangryyouth\f생선가시\f9\f8632\f" // 연속 9개월
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

//"\x1b\t001800006100\fmygomiee\fwhs741\f준치는준치\f100\f0\f0\f8632\f100\f0\f0\fjunwoo\f" 별풍
//"\x1b\t001800007200\fmygomiee\fiseung154\f잠없는사람Zz\f10\f516\f0\f8632\f10\f0\f0\fkor_custom05\f"

//"\x1b\t001800007100\fmygomiee\fmandeuk500\f[V]만두님♡\f10\f519\f0\f8632\f10\f0\f0\fkor_custom11\f" 10개 선물
//"\x1b\t001200006300\f589856|163840\fmandeuk500(2)\f[V]만두님♡\f0\f0\f589824|163840\f // 519번째 팬클럽
