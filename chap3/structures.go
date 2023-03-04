package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

// Go 컴파일러에 의해 자동으로 해당 타입의 제로값으로 초기화
func zeroS() Entry {
	return Entry{}
}

func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// Go 컴파일러에 의해 초기화 후 포인터를 반환
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := initS("Lucas", "Jeon", 2023)
	fmt.Println("s2:", s2)
}
