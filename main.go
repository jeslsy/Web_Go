package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handler 로직을 작성
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	//3000포트로 localhost:3000을 호출
	http.ListenAndServe(":3000", nil)
}
