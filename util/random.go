package util

import (
	"fmt"
	"strings"

	"github.com/mazen160/go-random"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(number int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < number; i++ {
		n, _ := random.GetInt(k)
		c := alphabet[n]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}
