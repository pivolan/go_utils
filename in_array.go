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
func IsProxyUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "" && u.Port() != "" && (u.Scheme == "socks5" || u.Scheme == "http")
}
func FilterValidProxies(proxies []string) (validProxies []string) {
	for _, proxy := range proxies {
		if IsProxyUrl(proxy) {
			validProxies = append(validProxies, proxy)
		}
	}
	return
}
func GetIpFromUrl(proxy string) string {
	u, err := url.Parse(proxy)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
