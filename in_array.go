package go_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/url"
	"sync"
	"time"
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
func EasyLocalDb(filename string, data *map[string]interface{}, lock *sync.RWMutex, saveInterval time.Duration) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.Wrapf(err, "read file err, path: %s, try to create...", filename)
		b, errEncode := json.Marshal(data)
		if errEncode != nil {
			err = errors.Wrapf(err, "encode failed: %s", errEncode)
			return err
		}
		errWrite := ioutil.WriteFile(filename, b, 0655)
		if errWrite != nil {
			err = errors.Wrapf(err, "create file failed: %s", errWrite)
			return err
		}
	}
	err = json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrapf(err, "cannot json decode, data: %s", string(b))
	}
	go func() {
		for {
			lastB := []byte{}
			time.Sleep(saveInterval)
			lock.RLock()
			b, err := json.Marshal(data)
			lock.RUnlock()
			if err != nil {
				log.Println(errors.Wrapf(err, "go routine err: cannot json decode, data: %s", string(b)))
				return
			}
			if bytes.Equal(lastB, b) {
				continue
			}
			err = ioutil.WriteFile(filename, b, 0655)
			if err != nil {
				log.Println(errors.Wrapf(err, "go routine err: cannot save to file"))
				return
			}
		}
	}()
	return nil
}
