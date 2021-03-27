package main

import (
	"net/http"

	"github.com/Tucker_Programming/Web_with_Go/myapp"
)

func main() {
	//3000포트로 localhost:3000을 호출
	// ListenAndServer로 웹서버를 구동
	// request를 기다리기 시작
	// 이미 핸들러 등록되어 있으면 실행해서 응답 보냄
	// 3000포트에서
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
