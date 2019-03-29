package cogs

import (
	"fmt"
	"net/http"
)

func Lunch(token string) ([]byte, error) {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	return request(fmt.Sprintf("%s/Lunches/Weekly", baseURL), http.MethodGet, "", headers)
}
