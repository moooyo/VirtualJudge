package proxy

import (
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

/*
 	Login()
	Logout()
	SetCookie()
	GetProblem()
	Submit()
	QuerySubmitStatus()
	LoginStatus() bool
*/
type POJ struct {
	username string
	password string
	cookies  []*http.Cookie
}

const baseURL = "http://poj.org"
const loginURL = "http://poj.org/login"
const loginLog = "http://poj.org/loginlog"
const problemUrl = "http://poj.org/problem"
const submitUrl = "http://poj.org/submit"
const statusUrl = "http://poj.org/status"

var pojLanguage = []struct {
	TypeName string
	ID       int
}{
	{
		TypeName: "G++",
		ID:       0,
	},
	{
		TypeName: "GCC",
		ID:       1,
	},
	{
		TypeName: "Java",
		ID:       2,
	},
	{
		TypeName: "Pascal",
		ID:       3,
	},
	{
		TypeName: "C++",
		ID:       4,
	},
	{
		TypeName: "C",
		ID:       5,
	},
	{
		TypeName: "Fortran",
		ID:       6,
	},
}

func (r *POJ) Login() error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

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

	_, err = client.PostForm(loginURL, v)
	if err != nil {
		return err
	}

	r.cookies = jar.Cookies(u)

	return nil
}

func (r *POJ) LoginStatus() (bool, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return false, err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return false, err
	}

	jar.SetCookies(u, r.cookies)
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       0,
	}

	resp, err := client.Get(loginLog)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false, err
	}

	ret := doc.Find("td.h").Eq(4).Find("a").Eq(1).Text()
	if ret == "Log Out" {
		return true, nil
	}
	return false, nil
}

func (r *POJ) GetProblem(problemID int) (*ProblemInfo, error) {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	u := problemUrl + "?id=" + strconv.Itoa(problemID)
	resp, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, StatusError(resp.StatusCode, u)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	problem := &ProblemInfo{
		OJ:           PojCode,
		ProblemID:    problemID,
		ProblemName:  doc.Find("div.ptt").Eq(0).Text(),
		Description:  doc.Find("div.ptx").Eq(0).Text(),
		TimeLimit:    doc.Find("div.plm").Find("td").Eq(0).Text()[12:],
		MemoryLimit:  doc.Find("div.plm").Find("td").Eq(2).Text()[14:],
		Input:        doc.Find("div.ptx").Eq(1).Text(),
		Output:       doc.Find("div.ptx").Eq(2).Text(),
		SampleInput:  make([]string, 0),
		SampleOutput: make([]string, 0),
		Language:     pojLanguage,
		Source:       doc.Find("div.ptx").Eq(3).Text(),
	}
	idx := doc.Find("p.pst").Eq(3).Next()
	status := 0
	for {
		str := idx.Text()
		if str == "Sample Output" {
			status = 1
		} else if str == "Source" {
			break
		} else {
			if status == 0 {
				problem.SampleInput = append(problem.SampleInput, str)
			} else {
				problem.SampleOutput = append(problem.SampleOutput, str)
			}
		}
		idx = idx.Next()
	}
	return problem, err
}

// return run id
func (r *POJ) Submit(args *SubmitArgs) (int, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return 0, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return 0, err
	}
	jar.SetCookies(u, r.cookies)

	s := base64.StdEncoding.EncodeToString([]byte(args.Source))
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       0,
	}
	params := url.Values{}
	params.Set("problem_id", strconv.Itoa(args.ProblemID))
	params.Set("language", strconv.Itoa(args.Language))
	params.Set("source", s)
	params.Set("submit", "Submit")
	params.Set("encoded", "1")

	resp, err := client.PostForm(submitUrl, params)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	StatusCode := resp.StatusCode
	if StatusCode != http.StatusOK {
		return 0, StatusError(StatusCode, submitUrl)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, err
	}

	title := doc.Find("title").Text()
	if title == "Error" {
		return 0, fmt.Errorf(doc.Find("li").Text())
	}

	nickName := doc.Find("td.h").Eq(4).Find("b").Eq(0).Text()
	statusListUrl := statusUrl + "?user_id=" + nickName
	resp, err = client.Get(statusListUrl)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, StatusError(resp.StatusCode, statusListUrl)
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0, err
	}
	runID := doc.Find("tr.in").Next().Find("td").Eq(0).Text()
	rid, err := strconv.Atoi(runID)
	if err != nil {
		return 0, err
	}
	return rid, nil
}
