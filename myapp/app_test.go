//이름_test.go하면 테스트 코드로 자동 인식
package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	// 네트워크 사용 않고 그냥 사용할 수 있는 레코더
	res := httptest.NewRecorder()
	// 새 호출
	req := httptest.NewRequest("GET", "/bar", nil)

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
	assert.Equal("Hello World!", string(data))
}
