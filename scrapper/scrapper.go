package scrapper

import (
	"io"
	"net/http"
	"strings"
)

type WebResource struct {
	URL          string
	ResponseBody string
}

func NewWebResource(URL string) *WebResource {
	return &WebResource{
		URL: URL,
	}
}

func (req *WebResource) GetResponse() (err error) {
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

func (req *WebResource) CountRepeatedStrInBody(subStr string) (occurrencesCount int) {
	return strings.Count(req.ResponseBody, subStr)
}
