package crawler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

const (
	URL_PREFIX = "https://atcoder.jp"
)

type Task struct {
	URL        string
	ScreenCode string
	Code       string
	Title      string
}

type SampleData struct {
	Input  string
	Output string
}

func GenDir(p Task, dt []SampleData) error {
	os.MkdirAll(p.Code, 0755)
	for i, d := range dt {
		input := filepath.Join(p.Code, fmt.Sprintf("in%d.txt", i+1))
		output := filepath.Join(p.Code, fmt.Sprintf("ans%d.txt", i+1))
		f1, _ := os.Create(input)
		defer f1.Close()
		f1.WriteString(d.Input)
		f1.WriteString("\n")

		f2, _ := os.Create(output)
		defer f2.Close()
		f2.WriteString(d.Output)
		f2.WriteString("\n")
	}
	return nil
}

func GetProblem(uri string) ([]SampleData, error) {
	var sams []SampleData
	url := fmt.Sprintf("%s%s", URL_PREFIX, uri)
	fmt.Println(url)
	resp, err := resty.New().SetRetryCount(50).SetRetryMaxWaitTime(time.Second * 10).SetTimeout(time.Second * 20).R().Get(url)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(resp.Body())
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	var sam SampleData
	doc.Find("section").Each(func(_ int, s *goquery.Selection) {
		if !strings.Contains(s.Find("h3").Text(), "Sample") {
			return
		}
		name := s.Find("h3").Text()
		data := strings.Trim(s.Find("pre").Text(), "\n \t")
		// fmt.Println(name, "\n", data)
		if strings.Contains(name, "Sample Input") {
			sam.Input = data
		} else if strings.Contains(name, "Sample Output") {
			sam.Output = data
			var tmp SampleData = sam
			sams = append(sams, tmp)
		}
	})
	return sams, nil
}

func GetTasks(contestCode string) ([]Task, error) {
	url := fmt.Sprintf("%s/contests/%s/tasks", URL_PREFIX, contestCode)
	fmt.Println(url)

	resp, err := resty.New().SetRetryCount(50).SetRetryMaxWaitTime(time.Second * 10).SetTimeout(time.Second * 20).R().Get(url)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(resp.Body())
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	var ps []Task
	doc.Find("tbody").Children().Each(func(i1 int, s1 *goquery.Selection) {
		var p Task
		s1.Find("td").Each(func(i2 int, s2 *goquery.Selection) {
			if i2 == 0 {
				problemUrl, _ := s2.Find("a").Attr("href")
				p.URL = problemUrl
				p.Code = problemUrl[len(problemUrl)-1:]
			} else if i2 == 1 {
				p.Title = s2.Text()
			}
		})
		if p.Code != "" {
			ps = append(ps, p)
		}
	})
	return ps, nil
}
