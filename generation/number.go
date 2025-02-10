package generation

import (
	"math/rand"
	"time"
)

func RandBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
