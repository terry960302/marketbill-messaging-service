package constants

type MessageType int

const (
	Default MessageType = 1 + iota
	Verification
	ApplyBizConnection
	ConfirmBizConnection
	RejectBizConnection
	IssueOrderSheetReceipt
)

var types = [...]string{
	"Default",
	"Verification",
	"ApplyBizConnection",
	"ConfirmBizConnection",
	"RejectBizConnection",
	"IssueOrderSheetReceipt",
}

var templates = [...]string{
	"[floway] %s",        // 메세지
	"[floway] 인증번호 : %s", // 인증번호
	"[floway] %s 님이 거래처 승인 요청을 보냈어요.\n마이페이지 > 거래처 승인에서 신청을 승인하면 거래처가 주문할 수 있어요.\n\nURL : %s", //소매상, URL
	"[floway] %s 님이 거래처 신청을 승인했어요.\n주문시, 거래처를 선택할때 “%s”을 선택할 수 있어요.\n\nURL : %s",             // 도매상, 도매상, URL
	"[floway] %s 님이 거래처 신청을 거절했어요.\n\n시장에서 %s 님께 직접 문의해주세요.",                                 // 도매상, 도매상
	"[floway] %s 님이 주문번호 %s에 영수증을 발행했어요.\n\n마이페이지 > 주문내역”에서 영수증을 다운로드 할 수 있어요.\n\nURL : %s",  // 도매상, 주문번호, URL
}

func (m MessageType) String() string { return types[(m-1)%12] }

func (m MessageType) Template() string {
	return templates[(m-1)%12]
}
