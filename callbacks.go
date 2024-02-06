package afreecachat

/*************
 * callbacks *
 *************/

// OnConnect 메서드는 서버와 연결에 관한 데이터를 전달한다.
func (c *Client) OnConnect(callback func(connect bool)) {
	c.onConnect = callback
}

func (c *Client) OnJoinChannel(callback func(join bool)) {
	c.onJoinChannel = callback
}

// OnChatMessage 메서드는 채팅 메시지가 왔을 때 데이터를 전달한다.
func (c *Client) OnChatMessage(callback func(message ChatMessage)) {
	c.onChatMessage = callback
}

// OnUserLists 메서드는 유저 입장/퇴장 데이터를 전달한다.
func (c *Client) OnUserLists(callback func(userlist []UserList)) {
	c.onUserLists = callback
}

// OnBalloon 메서드는 별풍선 데이터를 전달한다.
func (c *Client) OnBalloon(callback func(balloon Balloon)) {
	c.onBalloon = callback
}

// OnAdballoon 메서드는 애드벌룬 데이터를 전달한다.
func (c *Client) OnAdballoon(callback func(adballoon Adballoon)) {
	c.onAdballoon = callback
}

// OnSubscription 메서드는 구독 데이터를 전달한다.
func (c *Client) OnSubscription(callback func(subscription Subscription)) {
	c.onSubscription = callback
}
