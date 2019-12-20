package tools

type Err int

const (
	OK              Err = 0
	PostDataInvalid Err = 1001
	ConnectOJError  Err = 2001
)
