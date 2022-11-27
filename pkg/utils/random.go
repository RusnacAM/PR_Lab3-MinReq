package utils

import (
	"math/rand"
	"time"
)

func Random() int {
	var min = 0
	var max = 10000
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
