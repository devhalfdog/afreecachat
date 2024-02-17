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
        BJID: {BJ ID},
        PdBoxTicket: {PdBoxTicket},
        FanTicket: {FanTicket},
        Flag: {Flag},
    }
    client, err := afreecachat.NewClient(token)
    if err != nil {
        // client 생성 중 에러가 발생할 경우
        // panic() 을 호출합니다.
        panic(err)
    }

    client.OnError(func(err error) {
        // 에러가 발생할 경우 에러를 출력합니다.
        fmt.Println(err)
    })

    client.OnChatMessage(func(message afreecachat.ChatMessage) {
        // 채팅 메시지가 있을 경우 ID, NAME, MESSAGE를 출력합니다.
        fmt.Printf("ID: %s, NAME: %s, MESSAGE: %s\n", message.User.ID, message.User.Name, message.Message)
    })

    // 채팅방 접속을 시도합니다.
    // 에러가 발생할 경우 panic()을 호출합니다.
    err := client.Connect()
    if err != nil {
        panic(err)
    }
}
```

### Token
- `BJID` (필수)
  - 자동으로 `SocketAddress` 및 `ChatRoom`을 가져오기 위한 BJ 아이디입니다.
- `Flag` (필수)
  - 채팅 채널 연결에 필요한 유저 플래그 값
  - example: `524304`
- `Identify`
  - 채팅 채널 연결을 할 때 로그인 데이터
  - `Identify.ID` 및 `Identify.Password`의 값이 있을 경우 자동으로 로그인을 진행합니다.
  - 입력하지 않을 경우 비로그인으로 채팅 채널에 연결합니다.

### Callback
- `OnError(error)`
  - 에러가 발생할 경우 에러를 반환합니다.
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
- `OnAdminNotice(string)`
  - 운영자 알림 메시지가 있을 때마다 문자열을 반환합니다.
  - example: "{BJ NAME}님의 방송이 별별랭킹의 '웃음이 끊이지 않는 방송' 1위에 등극!"

### 예제
- 별풍선을 받을 때마다 [Warudo](https://warudo.app/) OSC 통신 예제 [문서](https://github.com/devhalfdog/afreeca-warudo)

## TODO
아래에 작성된 순서는 개발 순서가 아닌 생각난 대로 작성하였습니다.

- [ ] 에러 처리
  - `"\x1b\t000100005807\f시스템 에러가 발생 했습니다. (중복 세션)\f"`
- [x] 비밀방 접속
- [ ] 스티커를 받았을 때 콜백
- [ ] 코드 최적화
- [ ] 테스트 파일 작성
- [ ] 회원일 경우 플래그 설정
- [x] 로그인
  - 연령제한 방송에 접근하기 위해서는 인증된 계정 토큰이 필요함.

## 레퍼런스
- [https://github.com/wakscord/afreeca](https://github.com/wakscord/afreeca)