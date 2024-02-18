package translation

import (
	"fmt"
	"github.com/ktcf/hello-api/handlers/rest"
	"log"
	"strings"
)

var _ rest.Translator = &RemoteService{}

// RemoteService will allow for external call to existing translation service.
type RemoteService struct {
	client HelloClient
	cache  map[string]string
}

// HelloClient will call external translation service.
type HelloClient interface {
	Translate(word string, language string) (string, error)
}

// NewRemoteService will create a new RemoteService.
func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{
		client: client,
		cache:  make(map[string]string),
	}
}

// Translate will take a given word and try to find the result.
func (s *RemoteService) Translate(word string, language string) string {
	word = strings.ToLower(word)
	language = strings.ToLower(language)

	key := fmt.Sprintf("%s:%s", word, language)

	tr, ok := s.cache[key]
	if ok {
		return tr
	}

	resp, err := s.client.Translate(word, language)
	if err != nil {
		log.Println(err)
		return ""
	}

	s.cache[key] = resp
	return resp
}
