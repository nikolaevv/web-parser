package website

import (
	"io"
	"net/http"
	"strings"
)

type Website struct {
	URL          string
	ResponseBody string
}

func New(URL string) *Website {
	return &Website{
		URL: URL,
	}
}

func (req *Website) GetResponse() (err error) {
	resp, err := http.Get(req.URL)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	req.ResponseBody = string(body)

	if err = resp.Body.Close(); err != nil {
		return err
	}

	return nil
}

func (req *Website) Count(subStr string) (count int) {
	return strings.Count(req.ResponseBody, subStr)
}
