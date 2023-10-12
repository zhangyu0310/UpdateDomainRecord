package ipinfo

import "testing"

func TestGetMyIP(t *testing.T) {
	ip, err := GetMyIP()
	if err != nil {
		t.Error(err)
	}
	t.Log(ip)
}
