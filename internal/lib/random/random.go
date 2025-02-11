package random

import (
	"math/rand"
	"time"
)

func RandomURL(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijkmnopqrstuvwxyz" +
		"0123456789")

	res := make([]rune, size)
	for i := range res {
		res[i] = chars[rnd.Intn(len(chars))]
	}
	return string(res)
}
