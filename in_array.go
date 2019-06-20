package go_utils

import (
	"fmt"
	"net/url"
)

func InArray(val string, array []string) bool {
	for _, row := range array {
		if val == row {
			return true
		}
	}

	return false
}
func IsUrlParse(str string) (*url.URL, error) {
	u, err := url.Parse(str)
	if err == nil && u.Scheme != "" && u.Host != "" && (u.Scheme == "http" || u.Scheme == "socks5") {
		return url.Parse(str)
	}
	return nil, fmt.Errorf("not valid url proxy format")
}
