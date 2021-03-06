package endpoint

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://cogs.10pearls.com/cogsapi/api"

type HttpError struct {
	Code    int
	Message string
}

func (h HttpError) Error() string {
	return h.Message
}

var (
	ErrUnauthorized   = HttpError{http.StatusUnauthorized, "authentication failed"}
	ErrBadRequest     = HttpError{http.StatusBadRequest, "bad request"}
	ErrInternalServer = HttpError{http.StatusInternalServerError, "server not responding"}
)

func Request(url, verb, data string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	var (
		req *http.Request
		err error
	)
	switch verb {
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case http.MethodPost:
		if data == "" {
			req, err = http.NewRequest(http.MethodPost, url, nil)
		} else {
			req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(data)))
		}
	default:
		return nil, fmt.Errorf("unsupported http method: %s", verb)
	}
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusUnauthorized:
			return nil, ErrUnauthorized
		case http.StatusBadRequest:
			return nil, ErrBadRequest
		case http.StatusInternalServerError:
			return nil, ErrInternalServer
		default:
			return nil, HttpError{res.StatusCode, "something went wrong"}
		}
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
