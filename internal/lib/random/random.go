package random

import (
	"math/rand"
)

func RandomURL(length int) string {
	res := make([]byte, length)

	for i := range res {
		res[i] = byte(rand.Intn(255))
	}
	return string(res)
}
