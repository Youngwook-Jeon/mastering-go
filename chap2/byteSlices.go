package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b) // [0 0 0 0 0 0 0 0 0 0 0 0]
	b = []byte("Byte slice ₩")
	fmt.Println("Byte slice:", b) // [66 121 116 101 32 115 108 105 99 101 32 226 130 169]
	fmt.Println(len(b))           // 14

	// 바이트 슬라이스를 문자열로
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))
}
