package util

import (
	"math/rand"
)

func GenerateOTP() int {
	num := rand.Intn(900000) + 100000
	return num
}
