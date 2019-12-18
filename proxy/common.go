package proxy

type OnlineJudge interface {
	Login()
	Logout()
	SetCookie()
	GetProblem()
	Submit()
	QuerySubmitStatus()
}
