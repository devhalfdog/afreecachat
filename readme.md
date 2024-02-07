# AfreecaTV chat, Go library
공부용으로 제작하였습니다.

[아프리카TV](https://afreecatv.com) 채팅을 읽을 수 있는 [Go](https://go.dev) 라이브러리입니다.

공식 API가 없기 때문에 임시로 제작하였으므로 언제든지 막히거나 수정될 가능성이 매우 높습니다.

아직 실력이 미숙하여 코드에 대한 부분들은 언제든지 이슈 및 PR로 남겨주시면 감사하겠습니다.

## 사용 방법
**라이브러리 가져오기**

`go get -u https://github.com/devhalfdog/afreecachat`

```go
...

func main() {
    token := afreecachat.Token{
        PdBoxTicket: {PdBoxTicket},
        FanTicket: {FanTicket},
        ChatRoom: {ChatRoom},
        Flag: {Flag},
    }
    client, err := afreecachat.NewClient(token)
    if err != nil {
        panic(err)
    }

    client.SocketAddress = {WebSocket Address}

    client.OnChatMessage(func(message afreecachat.ChatMessage)) {
        fmt.Printf("ID: %s, NAME: %s, MESSAGE: %s\n", message.User.ID, message.User.Name, message.Message)
    })
}
```

### Token
- `BJID` (필수)
  - 자동으로 `SocketAddress` 및 `ChatRoom`을 가져오기 위한 BJ 아이디입니다.
- `PdBoxTicket`
  - 계정 연동에 필요한 `PdBoxTicket` 쿠키 값
  - 입력하지 않을 경우, 비로그인으로 접속합니다.
- `FanTicket`
  - 채팅 채널 연결에 필요한 `FanTicket` 값
  - 입력하지 않아도 됩니다.
- `Flag` (필수)
  - 채팅 채널 연결에 필요한 유저 플래그 값
  - example: `524304`

### Callback
- `OnConnect(bool)`
  - 채널 입장 Handshake가 성공하면 `true`를 반환합니다.
- `OnRawMessage(string)`
  - 원본 데이터 문자열을 반환합니다.
- `OnChatMessage(ChatMessage)`
  - 채팅 메시지가 있을 때마다 `ChatMessage` 구조체를 반환합니다.
- `OnUserLists([]UserList)`
  - 유저 입장/퇴장 메시지가 있을 때마다 `[]UserList` 구조체를 반환합니다.
- `OnBalloon(Balloon)`
  - 별풍선 메시지가 있을 때마다 `Balloon` 구조체를 반환합니다.
- `OnAdballoon(Adballoon)`
  - 애드벌룬 메시지가 있을 때마다 `Adballoon` 구조체를 반환합니다.
- `OnSubscription(Subscription)`
  - 구독 메시지가 있을 때마다 `Subscription` 구조체를 반환합니다.

### 예제
- 별풍선을 받을 때마다 [Warudo](https://warudo.app/) OSC 통신 예제 [문서](https://github.com/devhalfdog/afreeca-warudo)

## TODO
아래에 작성된 순서는 개발 순서가 아닌 생각난 대로 작성하였습니다.

- [ ] 에러 처리
  - `"\x1b\t000100005807\f시스템 에러가 발생 했습니다. (중복 세션)\f"`
- [ ] 비밀방 접속
- [ ] 스티커를 받았을 때 콜백
- [ ] 코드 최적화
- [ ] 테스트 파일 작성

## 레퍼런스
- [wakscord/afreeca](https://github.com/wakscord/afreeca)