package controller

import (
	"bootstrap/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(c *gin.Context, param gin.Params, u url.Values, method string, body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.URL.RawQuery = u.Encode()
	c.Request.Body = body
}

func TestPost(t *testing.T) {

	service := usecase.NewTaskUsecase()
	taskController := NewControllerInterface(service)

	t.Run("strategy_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		MakeRequest(context, nil, url.Values{}, "POST", nil)
		taskController.RunTask(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

}
