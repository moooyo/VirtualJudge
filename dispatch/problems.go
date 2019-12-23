package dispatch

type ProblemArgs struct {
	OJ        int `uri:"OJ" binding:"required"`
	ProblemID int `uri:"ProblemID" binding:"required"`
}

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

func GetProblemInfo(args *ProblemArgs) (*ProblemInfo, error) {
	oj := GetOJWithoutUserInfo(args.OJ)
	return oj.GetProblem(args.ProblemID)
}
