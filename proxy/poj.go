package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

/*
 	Login()
	Logout()
	SetCookie()
	GetProblem()
	Submit()
	QuerySubmitStatus()
*/
type POJ struct {
	username string
	password string
	cookies  []*http.Cookie
}

const pojBaseURL = "http://poj.org"
const pojLoginURL = "http://poj.org/login"
const pojLoginLog = "http://poj.org/loginlog"

func (r *POJ) Login() error {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	v := url.Values{}
	v.Set("user_id1", r.username)
	v.Set("password1", r.password)
	v.Set("B1", "login")
	v.Set("B1", "%2Fregister")

	_, err = client.PostForm(pojLoginURL, v)
	if err != nil {
		return err
	}

	resp, err := client.Get(pojLoginLog)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}
