package scrapper

import (
	"io"
	"net/http"
	"strings"
)

type Request struct {
	URL          string
	ResponseBody string
}

func NewRequest(URL string) *Request {
	return &Request{
		URL: URL,
	}
}

func (req *Request) GetResponse() (err error) {
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

func (req *Request) CountRepeatedStrInBody(subStr string) (occurrencesCount int) {
	return strings.Count(req.ResponseBody, subStr)
}
