package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Printf("Now you have %g problems! \n", math.Sqrt(rand.Float64()))
}
