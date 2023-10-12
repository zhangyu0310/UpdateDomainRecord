package ipinfo

import (
	"errors"
	"fmt"
	"github.com/zhangyu0310/zlogger"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var testUrl = []string{
	"icanhazip.com",
	"ifconfig.me",
	"api.ipify.org",
	"ipinfo.io/ip",
	"ipecho.net/plain",
}

var (
	ErrGetIPDifferent = errors.New("get ip different")
)

func getIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			zlogger.Warn("Close response body failed, err:", err)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetMyIP() (string, error) {
	first := rand.Intn(len(testUrl))
	second := rand.Intn(len(testUrl))
	if first == second {
		second = (second + 1) % len(testUrl)
	}
	zlogger.Info("first url:", testUrl[first], " second url:", testUrl[second])

	ip1, err := getIP(fmt.Sprintf("http://%s", testUrl[first]))
	if err != nil {
		return "", err
	}
	ip1 = strings.TrimSpace(ip1)

	ip2, err := getIP(fmt.Sprintf("http://%s", testUrl[second]))
	if err != nil {
		return "", err
	}
	ip2 = strings.TrimSpace(ip2)

	if ip1 == ip2 && ip1 != "" {
		return ip1, nil
	} else {
		return "", ErrGetIPDifferent
	}
}
