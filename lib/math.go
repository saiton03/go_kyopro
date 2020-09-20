package kyopro

import (
	"math"
)

func Gcd(nums ...int64) int64 {
	if len(nums) == 0 {
		panic("Gcd func requires one or more integers")
	}
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret = gcd(ret, nums[i])
	}
	return ret
}

func gcd(x, y int64) int64 {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func IsPrime(x int64) bool {
	if x < 2 {
		return false
	}
	bound := int64(math.Sqrt(float64(x)))

	for i := int64(2); i <= bound; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
