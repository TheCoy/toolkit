package thirdpart

import (
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var deviceId string
var once sync.Once

func GenDeviceId() string {
	once.Do(func() {
		deviceId = uuid.NewString()
	})
	return deviceId
}

func GenTraceId() string {
	traceId, _ := uuid.NewUUID()
	return traceId.String()
}

// GenerateRandomString returns a securely generated random string like "1uC1hy6ebiVnWaUXm7SIp".
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
