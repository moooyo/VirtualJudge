package dispatch

func GetOJWithoutUserInfo(oj int) OnlineJudge {
	switch oj {
	case PojCode:
		return new(POJ)
	default:
		return nil
	}
}
