package dispatch

import "fmt"

type OnlineJudge interface {
	Login() error
	//Logout()
	//SetCookie()
	GetProblem(problemID int) (*ProblemInfo, error)
	Submit(args *SubmitArgs) (int, error)
	QuerySubmitStatus(int) (*StatusResp, error)
	LoginStatus() (bool, error)
}

func StatusError(status int, url string) error {
	return fmt.Errorf("url is %s. status code is %d", url, status)
}

// Online Judge Code
const (
	PojCode = 1
)

// Result Code
const (
	UnrecognizedError = -1
	Accepted          = iota
	WrongAnswer
	CompileError
	RuntimeError
	TimeLimitExceeded
	PresentationError
	MemoryLimitExceeded
	OutputLimitExceeded
)

type SubmitArgs struct {
	OJ        int
	Source    string
	ProblemID int
	Language  int
}

type StatusResp struct {
	Result     int
	Problem    int
	Memory     string
	Time       string
	Language   string
	Length     string
	SubmitTime string
	Info       string
}
