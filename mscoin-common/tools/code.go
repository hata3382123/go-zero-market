package tools

import (
	"fmt"
	"math/rand/v2"
)

func GetCode() string {
	code := rand.IntN(10000)
	return fmt.Sprint("%04d", code)
}
