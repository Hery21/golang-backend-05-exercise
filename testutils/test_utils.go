package testutils

import (
	"encoding/json"
	"fmt"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ServeReq(opts *server.RouterConfig, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.NewRouter(opts)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func ErrorToString(httpErr httperror.AppError) string {
	return fmt.Sprintf(`{"statusCode":%d,"code":"%s","message":"%s"}`,
		httpErr.StatusCode, httpErr.Code, httpErr.Message,
	)
}
