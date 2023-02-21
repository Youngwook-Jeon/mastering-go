package main

import "fmt"

type aStruct struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)

	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f) // Value of f: 146.97

	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	var k *aStruct // 포인터 타입의 제로 값인 nil을 가리킴
	fmt.Println(k)
	if k == nil {
		k = new(aStruct)
	}

	fmt.Printf("%+v\n", k) // &{field1:(0+0i) field2:0}
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
