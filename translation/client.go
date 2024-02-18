package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

var _ HelloClient = &APIClient{}

type APIClient struct {
	endpoint string
}

// NewHelloClient creates instance of client with a given endpoint.
func NewHelloClient(endpoint string) *APIClient {
	return &APIClient{
		endpoint: endpoint,
	}
}

// Translate will call external translation service.
func (c *APIClient) Translate(word, language string) (string, error) {
	req := map[string]string{
		"word":     word,
		"language": language,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("unable to encode msg")
	}

	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return "", errors.New("unable to call service")
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	} else if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New("error in api")
	}

	b, _ = io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)

	var m map[string]any
	if err = json.Unmarshal(b, &m); err != nil {
		return "", errors.New("unable to decode msg")
	}

	return m["translation"].(string), nil
}
