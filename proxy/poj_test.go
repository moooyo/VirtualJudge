package proxy

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	poj := POJ{
		username: "ccc",
		password: "ccc",
		cookies:  nil,
	}
	err := poj.Login()
	if err != nil {
		fmt.Println(err)
	}
}
