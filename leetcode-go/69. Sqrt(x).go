package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func main() {
	fmt.Println(mySqrt(8))

}
