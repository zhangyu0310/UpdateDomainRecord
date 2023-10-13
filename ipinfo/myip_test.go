package ipinfo

import "testing"

func TestGetMyIP(t *testing.T) {
	for i := 0; i < 10; i++ {
		ip, err := GetMyIP()
		if err != nil {
			t.Error(err)
			break
		}
		t.Log(ip)
	}
}
