package main

import "golang.org/x/tour/pic"

// グレースケール画像の作成
func Pic(dx, dy int) [][]uint8 {
	picSlices := make([][]uint8, dy) // picSlicesにuint8の二次元スライスをdy個割り当て

	for i := 0; i < dy; i++ {
		picSlices[i] = make([]uint8, dx) // [][]picSliceに[]uint8をdx個割り当て
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			picSlices[i][j] = uint8(i ^ j) // グレースケール値の計算
		}
	}
	return picSlices
}

func main() {
	pic.Show(Pic)
}
