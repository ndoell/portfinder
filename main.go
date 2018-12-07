package main

import (
	"math/rand"
	"fmt"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}



func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(random(5000, 9000))
}
