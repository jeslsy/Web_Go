//이름_test.go하면 테스트 코드로 자동 인식
package myapp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPathHandler(t *testing.T) {
	// 네트워크 사용 않고 그냥 사용할 수 있는 레코더
	res := httptest.NewRecorder()
	// 새 호출
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req)

	// 실패하면
	if res.Code != http.StatusBadRequest {
		// 프로그램 종료
		t.Fatal("Failed!! ", res.Code)
	}

}
