package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// TempDir의 존재여부나 접근 권한이 보장되지 않음!
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	// 파일이 이미 존재한다면 파일을 열어 파일의 끝에 새 데이터를 추가(os.O_APPEND)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go!")
}
