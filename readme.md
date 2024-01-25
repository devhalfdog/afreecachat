# AfreecaTV Chat library
[아프리카TV](https://afreecatv.com) 채팅을 읽을 수 있는 [Go](https://go.dev) 라이브러리입니다.

공식 API가 없기 때문에 임시로 제작하였고, 언제든지 막히거나 수정될 가능성이 매우 높습니다.

## 사용 방법
**라이브러리 가져오기**

`go get https://github.com/devhalfdog/afreecachat`

```go
...

func main() {
    token := afreecachat.Token{
        PdBoxTicket: {PdBoxTicket},
        FanTicket: {FanTicket},
        ChatRoom: {ChatRoom},
        Flag: {Flag},
    }
    client := afreecachat.NewClient(token)
    client.SocketAddress = {WebSocket Address}

    client.OnChatMessage(func(message afreecachat.ChatMessage)) {
        fmt.Printf("ID: %s, NAME: %s, MESSAGE: %s\n", message.User.ID, message.User.Name, message.Message)
    })
}
```

### Callback
- `OnConnect(bool)`
  - 채널 입장 Handshake가 성공하면 `true`를 반환한다.
- `OnChatMessage(ChatMessage)`
  - 채팅 메시지가 있을 때마다 `ChatMessage` 구조체를 반환한다.
- `OnUserLists([]UserList)`
  - 유저 입장/퇴장 메시지가 있을 때마다 `[]UserList` 구조체를 반환한다.
- `OnBalloon(Balloon)`
  - 별풍선 메시지가 있을 때마다 `Balloon` 구조체를 반환한다.
- `OnAdballoon(Adballoon)`
  - 애드벌룬 메시지가 있을 때마다 `Adballoon` 구조체를 반환한다.
- `OnSubscription(Subscription)`
  - 구독 메시지가 있을 때마다 `Subscription` 구조체를 반환한다.