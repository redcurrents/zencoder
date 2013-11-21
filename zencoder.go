package zencoder

import (
	"fmt"
	"io"
	"net/http"
)

type Zencoder struct {
	BaseUrl string
	Header  http.Header
	Client  *http.Client
}

func NewZencoder(apiKey string) *Zencoder {
	return &Zencoder{
		Client:  http.DefaultClient,
		BaseUrl: "https://app.zencoder.com/api/v2/",
		Header: http.Header{
			"Content-Type":     []string{"application/json"},
			"Accept":           []string{"application/json"},
			"Zencoder-Api-Key": []string{apiKey},
			"User-Agent":       []string{"gozencoder v1"},
		},
	}
}

func (z *Zencoder) call(method, path string, body io.ReadCloser) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", z.BaseUrl, path), body)
	if err != nil {
		return nil, err
	}

	req.Header = z.Header

	return z.Client.Do(req)
}
