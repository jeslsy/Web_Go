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

	indexHandler(res, req)

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
