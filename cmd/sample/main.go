package main

import (
	"fmt"
	k "go_kyopro/lib"
	"math"
)

func main() {
	var a, b int64 = 9, int64(math.Max(11, 18))
	d := k.Gcd(a, b)
	fmt.Println(d)
}
