package lib

import (
	"math/rand"
	"time"
)

func Random(args []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	return args[index]
}
