package controller

import (
	"bootstrap/controller/task/model"
	mock "bootstrap/tests/mocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockTaskUsecase(ctrl)
	controller := NewControllerInterface(service)

	t.Run("strategy_is_invalid", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		MakeRequest(context, nil, url.Values{}, "POST", nil)
		controller.RunTask(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("request_validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		request := model.TaskRequest{
			Action:    "1",
			Parameter: "1",
		}

		b, _ := json.Marshal(request)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, nil, url.Values{}, "POST", stringReader)
		controller.RunTask(context)
		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("delete_logs_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		expectedBody := `{"result":"deleted"}`

		request := model.TaskRequest{
			Action:    "delete_logs",
			Parameter: "group_proposta",
		}

		b, _ := json.Marshal(request)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().Task(gomock.Any()).Return("deleted")

		MakeRequest(context, nil, url.Values{}, "POST", stringReader)
		controller.RunTask(context)
		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, expectedBody, recorder.Body.String())
	})

}
