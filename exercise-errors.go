package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x ErrNegativeSqrt) (float64, error) {
	if x < 0 {
		return math.NaN(), x
	}
	
	z := 1.0 // 初期値1.0
	for {
		zTmp := z                       
		z = z - (math.Pow(z, 2)-float64(x))/(2*z)	// ニュートン法

		// 誤差との比較 (Zn - Zn-1)
		if math.Abs(z-zTmp) <= 0.0001 {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
