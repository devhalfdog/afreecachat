package afreecachat

import "github.com/gorilla/websocket"

type Client struct {
	SocketAddress string // Socket Address
	Token         Token

	socket *websocket.Conn
	read   chan []byte

	handshake [][]byte

	// callback
	onConnect      func(connect bool)
	onChatMessage  func(message ChatMessage)
	onUserLists    func(userlist []UserList)
	onAdballoon    func(adballoon Adballoon)
	onBalloon      func(balloon Balloon)
	onSubscription func(subscrption Subscription)
}

type Token struct {
	PdBoxTicket string
	FanTicket   string
	ChatRoom    string
	Flag        string
}

type Log struct {
	SetBps          string `json:"set_bps"`
	ViewBps         string `json:"view_bps"`
	Quality         string `json:"quality"`
	GeoContryCode   string `json:"geo_cc"`
	GeoRegionCode   string `json:"geo_rc"`
	AcceptLanguage  string `json:"acpt_lang"`
	ServiceLanguage string `json:"svc_lang"`
	JoinContryCode  string `json:"join_cc"`
	Subscribe       string `json:"subscribe"`
}

type Info struct {
	Password     string `json:"pwd"`
	AuthInfo     string `json:"auth_info"`
	PVer         string `json:"pver"`
	AccessSystem string `json:"access_system"`
}

type User struct {
	ID   string // 유저 아이디
	Name string // 유저 닉네임
	Flag UserFlag
}

type UserFlag struct {
	Admin            bool // 관리자
	Hidden           bool // 아이디 숨김
	BJ               bool // 방장
	Dumb             bool // 벙어리
	Guest            bool // 비회원
	Fanclub          bool // 팬클럽 회원
	AutoManager      bool // 자동 매니저
	ManagerList      bool // 자동 매니저 리스트에 등록된 사람
	SubBJ            bool // 부방장, 매니저
	Female           bool // 여자, 거짓이면 남자
	AutoDumb         bool // 자동 벙어리
	DumbBlind        bool // 벙어리로 인한 블라인드
	PaperingBlind    bool // 도배로 인한 블라인드
	ExitUser         bool // 나간 사람
	Mobile           bool // 모바일 유저
	TopFan           bool // 열혈팬
	Realname         bool // 실명인증
	NoDirect         bool // 1:1 직접 채팅 금지
	GlobalApp        bool // 글로벌 모바일 앱 사용자
	QuickView        bool // 퀵 뷰 사용자
	StickerSupporter bool // 스티커 서포터
	Chromecast       bool // 크롬 캐스트 사용자
	Subscription     bool // 구독팬
}

type ChatMessage struct {
	User    User
	Message string // 채팅
}

type UserList struct {
	User   User
	Status bool // 입장 true, 퇴장 false
}

type Balloon struct {
	User  User
	Count int // 별풍선 갯수
}

type Adballoon struct {
	User  User
	Count int // 애드벌룬 갯수
}

type Subscription struct {
	User  User
	Count int // 구독 개월 수
}
