package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// フィボナッチ数を返す関数を返す
// 返された関数は呼び出すたびに次項のフィボナッチ数を返す
func fibonacci() func() int {
	previous, next := 0, 1 // フィボナッチ数列のAnとA(n+1)
	count := 0             // フィボナッチ数計算関数の呼び出し回数(初項と第2項とそれ以外の選別に必要)

	return func() int {
		count++
		switch count {
		case 1: // フィボナッチ数列の初項(0)の場合
			return 0
		case 2: // フィボナッチ数列の第2項(1)の場合
			return 1
		default:
			previous, next = next, previous+next // フィボナッチ数計算
			return next
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
