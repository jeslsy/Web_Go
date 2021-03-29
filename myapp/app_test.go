//이름_test.go하면 테스트 코드로 자동 인식
package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	// 네트워크 사용 않고 그냥 사용할 수 있는 레코더
	res := httptest.NewRecorder()
	// 새 호출
	req := httptest.NewRequest("GET", "/", nil)

	// barHandler(res, req)
	// mux를 이용해서 url 나눠주기.
	// /bar로 안하고 /이걸로 하면 테스트에서 FAIL뜸
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// 매번 검사하지 말고
	//if res.Code != http.StatusBadRequest {
	//	// 프로그램 종료
	//	t.Fatal("Failed!! ", res.Code)
	//}

	// "github.com/stretchr/testify/assert"이 패키지를 이용해 확인
	assert.Equal(http.StatusOK, res.Code)

	//ioutil 패키지로 버퍼값을 다 읽어 오겠음.
	data, _ := ioutil.ReadAll(res.Body)
	// byte array 타입이기 때문에 string으로 바꿔서 비교
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	// 네트워크 사용 않고 그냥 사용할 수 있는 레코더
	res := httptest.NewRecorder()
	// 새 호출
	req := httptest.NewRequest("GET", "/bar?name=tucker", nil)

	// barHandler(res, req)
	// mux를 이용해서 url 나눠주기.
	// /bar로 안하고 /이걸로 하면 테스트에서 FAIL뜸
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// 매번 검사하지 말고
	//if res.Code != http.StatusBadRequest {
	//	// 프로그램 종료
	//	t.Fatal("Failed!! ", res.Code)
	//}

	// "github.com/stretchr/testify/assert"이 패키지를 이용해 확인
	assert.Equal(http.StatusOK, res.Code)

	//ioutil 패키지로 버퍼값을 다 읽어 오겠음.
	data, _ := ioutil.ReadAll(res.Body)
	// byte array 타입이기 때문에 string으로 바꿔서 비교
	// app.go에서 bar는 !를 붙여 줬기 때문에 Hello World!라고 해줘야함.
	assert.Equal("Hello tucker!", string(data))
}

// Foo는 Json 받아서 실행하는 함수였음!
func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	// 네트워크 사용 않고 그냥 사용할 수 있는 레코더
	res := httptest.NewRecorder()
	// 새 호출
	// goconvey에 보내는 건가봐.
	// nil(input 없음)로 보내면 실패코드옴 -> 그래서 아래 assert라인 Badrequest 해야 pass됨.
	req := httptest.NewRequest("POST", "/foo",
		// NewReader로 이 string이 ioreader로 바뀜. (버퍼로 바껴서)
		strings.NewReader(`{"first_name": "tucker", "last_name":"kim", "email":"tucker@naver.com"}`))

	// mux 호출
	// barHandler(res, req)
	// mux를 이용해서 url 나눠주기.
	// /bar로 안하고 /이걸로 하면 테스트에서 FAIL뜸
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	/*여기까지 보내기(요청)*/

	/*여기부터 받기(응답)*/

	//response
	// JSON 있으니까 200코드로 올것 = StatusCreated
	assert.Equal(http.StatusCreated, res.Code)

	//데이터가 잘 왔는지 확인해보자.
	user := new(User)
	// response된 result를 user로 다시 디코더 하겠음.
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err) //에러가 nil(없어야)함.

	assert.Equal("tucker", user.FirstName)
	assert.Equal("kim", user.LastName)

}
