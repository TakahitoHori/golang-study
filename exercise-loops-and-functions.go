// Exercise: Loops and Functions
// 2017/04/13 堀貴仁
package main

import (
	"fmt"
	"math"
)

const (
	E = 0.0001 // 誤差
)

// ニュートン法による、平方根の計算
func Sqrt(x float64) float64 {
	// 虚数の場合はNaNを返す
	if x < 0 {
		return math.NaN()
	}

	z := 1.0 // 初期値1.0
	for {
		zTmp := z                        // Zn-1
		z = z - (math.Pow(z, 2)-x)/(2*z) // Zn

		// 誤差との比較 (Zn - Zn-1)
		if math.Abs(z-zTmp) <= E {
			return z
		}
	}
}

func main() {
	fmt.Println("Sqrt(2)\t\t:", Sqrt(2))
	fmt.Println("math.Sqrt(2)\t:", math.Sqrt(2))
}
