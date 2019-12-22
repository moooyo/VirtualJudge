package proxy

import "fmt"

type OnlineJudge interface {
	Login()
	Logout()
	SetCookie()
	GetProblem()
	Submit()
	QuerySubmitStatus()
	LoginStatus() bool
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

type ProblemInfo struct {
	OJ           int
	ProblemID    int
	ProblemName  string
	Description  string
	Source       string
	TimeLimit    string
	MemoryLimit  string
	Input        string
	Output       string
	SampleInput  []string
	SampleOutput []string
	Language     []struct {
		TypeName string
		ID       int
	}
}

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
