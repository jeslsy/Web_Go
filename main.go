package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

// ServeHTTP = 함수 하나를 가진 인터페이스임 
// 두개의 인자를 받아 서버를 구동
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func main() {
	// handler 등록
	// func 형태로 handler를 등록
	// url 기능인듯
	// /인덱스 페이지 들어오면 func를 수행
	// w = 요청 작성
	// r = 요청 정보 들어있음
	// 함수를 직접 지정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w에 문자열 출력하라고 Fprint사용
		fmt.Fprint(w, "Hello World")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		// w에 문자열 출력하라고 Fprint사용
		fmt.Fprint(w, "Hello Bar!")
	})

	// 있는 함수 가져옴
	http.Handle("/foo", &fooHandler{})

	//3000포트로 localhost:3000을 호출
	// ListenAndServer로 웹서버를 구동
	// request를 기다리기 시작
	// 이미 핸들러 등록되어 있으면 실행해서 응답 보냄
	// 3000포트에서
	http.ListenAndServe(":3000", nil)
}
