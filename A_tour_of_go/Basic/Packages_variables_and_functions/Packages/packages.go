package main

import (
	"fmt"
	"math/rand"
	"time"
)

const interval int = 100

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Mi favorite number is ", rand.Intn(interval))
}
