package controller

import (
	"bootstrap/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunTaskWithRealService(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("delete_logs_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		requestBody := `{"action":"delete_logs","parameter":"group_proposta"}`
		request, _ := http.NewRequest("POST", "/task", strings.NewReader(requestBody))
		context.Request = request

		service := usecase.NewTaskUsecase()
		controller := NewControllerInterface(service)
		controller.RunTask(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		expectedBody := `{"result":"deleted"}`
		assert.JSONEq(t, expectedBody, recorder.Body.String())

	})

	t.Run("update_database_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		requestBody := `{"action":"update_database","parameter":"update proposta set valor=1000 where nome='groselha';"}`
		request, _ := http.NewRequest("POST", "/task", strings.NewReader(requestBody))
		context.Request = request

		service := usecase.NewTaskUsecase()
		controller := NewControllerInterface(service)
		controller.RunTask(context)

		assert.Equal(t, http.StatusOK, recorder.Code)
		expectedBody := `{"result":"Database has been updated with: update proposta set valor=1000 where nome='groselha';"}`
		assert.JSONEq(t, expectedBody, recorder.Body.String())

	})

	t.Run("catch exception", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		requestBody := `{"action":"xx","parameter":"xxxx"}`
		request, _ := http.NewRequest("POST", "/task", strings.NewReader(requestBody))
		context.Request = request

		service := usecase.NewTaskUsecase()
		controller := NewControllerInterface(service)
		controller.RunTask(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})

}
