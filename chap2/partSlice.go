package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	// 첫 5개의 원소
	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	// 마지막 2개의 원소
	fmt.Println(aSlice[l-2 : l])
	fmt.Println(aSlice[l-2:])

	// 첫 5개의 원소
	t := aSlice[0:5:10]
	fmt.Println(len(t), cap(t))

	// 2, 3, 4 인덱스의 원소
	t = aSlice[2:5:10] // 용량은 10-2=8
	fmt.Println(t)
	fmt.Println(len(t), cap(t))
}
