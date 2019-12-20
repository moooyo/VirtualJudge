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

const (
	PojCode = 1
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
