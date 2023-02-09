package main

import "fmt"

func main() {
	aSlice := []float64{} // 빈 슬라이스. 길이와 용량은 모두 0
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -23.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// 길이가 4인 슬라이스
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// 슬라이스에 남은 공간이 없을 경우 append()를 사용
	t = append(t, -5)
	fmt.Println(t)

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
