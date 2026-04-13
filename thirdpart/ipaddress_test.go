package thirdpart

import "testing"

func TestGetIP(t *testing.T) {
	t.Log(GetLocalIP())
}

func TestGenDeviceId(t *testing.T) {
	t.Log(GenDeviceId())
	for i := 0; i < 10; i++ {
		t.Log(GenDeviceIdWithSeed("hao123@email.com"))
	}
}
