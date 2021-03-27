package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//json 데이터 담을 sturct
type User struct {
	FirstName string    `json:"first_name"` // 설명붙여줌 = 어노테이션이라함
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

// 람다에서 따로 빼줘서 함수로 /인덱스 페이지
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// ServeHTTP = 함수 하나를 가진 인터페이스임
// 두개의 인자를 받아 서버를 구동
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// User struct의 인스턴스 생성
	user := new(User)
	// json형태로 파싱을 해올것.
	// Reader 인터페이스에 body넣어줌
	// 읽어서 decode 할것.
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// bad상태를 알려주고
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	//err가 nil이면 = 문제 없으면
	// struct 상태
	user.CreatedAt = time.Now()
	// json형태로 다시 바꿔야함.
	// json encoding
	data, _ := json.Marshal(user)
	// text형식 아니고 json형태로 전달 받기위해 json type 설정
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json은 string이기 때문에 바이너리 형태를 string으로 변환해서 넘겨주자
	fmt.Fprint(w, string(data))

}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// 요청에 필요한 input값 넣을 수 있음
	// URL에서 정보를 뽑아냄
	// name이라는 인자를 뽑아내겠음
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	// 여기는 Fprintf임!
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	// http.Handler는 serveHttp 함수를 가지고 있는 인터페이스임.
	// mux 인스턴스를 하나 만들어서 mux를 넘겨주는 방식
	mux := http.NewServeMux()

	// router = URL에 따라 경로를 지정해주는것
	// handler 등록
	// func 형태로 handler를 등록
	// url 기능인듯
	// /인덱스 페이지 들어오면 func를 수행
	// w = 요청 작성
	// r = 요청 정보 들어있음
	// 함수를 직접 지정

	// 있는 함수 가져옴
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	return mux
}
