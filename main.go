// test project main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tarm/serial" //시리얼 통신 라이브러리 추가
)

func main() {
	fmt.Println("HTTP to Serial communication test")
	c := &serial.Config{Name: "COM34", Baud: 9600}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := s.Write([]byte(req.URL.String())) //연결된 GET DATA를 시리얼로 전송
		if err != nil {
			log.Fatal(err)
		}
		res.Write([]byte("OK"))
		fmt.Println(req.URL)
	})

	go http.ListenAndServe(":80", nil) // go 루틴 실행으로 병렬 처리
	fmt.Println("server started")

	for { //시리얼 반환값 출력
		buf := make([]byte, 1)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}
}
