package dispatch

import (
	"fmt"
	"testing"
)

var testLoginData = []struct {
	username string
	password string
	status   bool
}{
	{
		username: "ccccc",
		password: "cccccc",
		status:   false,
	},
	{
		username: "ccc",
		password: "ccc",
		status:   false,
	},
}

func TestPOJ_LoginStatus(t *testing.T) {
	for _, x := range testLoginData {
		poj := POJ{
			username: x.username,
			password: x.password,
			cookies:  nil,
		}
		err := poj.Login()
		if err != nil {
			t.Fail()
		}
		b, err := poj.LoginStatus()
		if err != nil {
			t.Fail()
		}
		if b != x.status {
			t.Fail()
		}
	}
}

func TestPOJ_GetProblem(t *testing.T) {
	var poj POJ
	problem, err := poj.GetProblem(1008)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(problem.SampleOutput)
}

var testSubmitArgs = []struct {
	args SubmitArgs
	err  error
}{
	{
		args: SubmitArgs{
			OJ: 1,
			Source: `#include<iostream>
						using namespace std;
						int main(){

						return 0;}`,
			ProblemID: 1009,
			Language:  1,
		},
		err: nil,
	},
	{
		args: SubmitArgs{
			OJ:        1,
			Source:    "####",
			ProblemID: 1010,
			Language:  1,
		},
		err: fmt.Errorf("%s", "Source code too long or too short,submit FAILED;if you really need submit this source please contact administrator"),
	},
}

func TestPOJ_Submit(t *testing.T) {
	poj := POJ{
		username: "ccc",
		password: "ccc",
		cookies:  nil,
	}
	err := poj.Login()
	if err != nil {
		t.Error(err)
	}
	for _, x := range testSubmitArgs {
		rid, err := poj.Submit(&x.args)
		wantStr := ""
		if x.err != nil {
			if err != nil {
				wantStr = x.err.Error()
			}
		}
		getStr := ""
		if err != nil {
			getStr = err.Error()
		}

		if wantStr != getStr {
			t.Error(err)
		}
		fmt.Println(rid)
	}
}

func TestPOJ_QuerySubmitStatus(t *testing.T) {
	poj := POJ{
		username: "",
		password: "",
		cookies:  nil,
	}
	resp, err := poj.QuerySubmitStatus(21168963)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Info)
}
