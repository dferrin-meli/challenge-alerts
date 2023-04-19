package common

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateGinTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	responseRecorder := httptest.NewRecorder()
	ginTestContext, _ := gin.CreateTestContext(responseRecorder)
	return ginTestContext, responseRecorder
}
