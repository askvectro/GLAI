package adapters

import (
	"bytes"
	"net/http"
	"io/ioutil"
)

type CoreMLAdapter struct {
	endpoint string
}

func NewCoreMLAdapter(endpoint string) *CoreMLAdapter {
	return &CoreMLAdapter{endpoint: endpoint}
}

func (c *CoreMLAdapter) Generate(prompt string) (string, error) {
	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer([]byte(`{"prompt":"`+prompt+`"}`)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
