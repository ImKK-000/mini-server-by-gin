package api_test

import (
	"io/ioutil"
	. "mini-serve/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_Handler_Call_Method_GET_Should_Be_JSON(t *testing.T) {
	expectedResponseBody := `{"method":"GET","ping":"pong"}`
	request := httptest.NewRequest(http.MethodGet, "/say", nil)
	writer := httptest.NewRecorder()
	routeTest := gin.Default()
	routeTest.GET("/say", Handler)
	routeTest.ServeHTTP(writer, request)

	response := writer.Result()
	actualResponseBody, _ := ioutil.ReadAll(response.Body)

	if expectedResponseBody != string(actualResponseBody) {
		t.Errorf("expect\n%s\nbut it got\n%s", expectedResponseBody, actualResponseBody)
	}
}

func Test_Handler_Call_Method_POST_Should_Be_JSON(t *testing.T) {
	expectedResponseBody := `{"method":"POST","ping":"pong"}`
	request := httptest.NewRequest(http.MethodPost, "/say", nil)
	writer := httptest.NewRecorder()
	routeTest := gin.Default()
	routeTest.POST("/say", Handler)
	routeTest.ServeHTTP(writer, request)

	response := writer.Result()
	actualResponseBody, _ := ioutil.ReadAll(response.Body)

	if expectedResponseBody != string(actualResponseBody) {
		t.Errorf("expect\n%s\nbut it got\n%s", expectedResponseBody, actualResponseBody)
	}
}

func Test_Handler_Call_Method_PUT_Should_Be_JSON(t *testing.T) {
	expectedResponseBody := `{"method":"PUT","ping":"pong"}`
	request := httptest.NewRequest(http.MethodPut, "/say", nil)
	writer := httptest.NewRecorder()
	routeTest := gin.Default()
	routeTest.PUT("/say", Handler)
	routeTest.ServeHTTP(writer, request)

	response := writer.Result()
	actualResponseBody, _ := ioutil.ReadAll(response.Body)

	if expectedResponseBody != string(actualResponseBody) {
		t.Errorf("expect\n%s\nbut it got\n%s", expectedResponseBody, actualResponseBody)
	}
}

func Test_Handler_Call_Method_DELETE_Should_Be_JSON(t *testing.T) {
	expectedResponseBody := `{"method":"DELETE","ping":"pong"}`
	request := httptest.NewRequest(http.MethodDelete, "/say", nil)
	writer := httptest.NewRecorder()
	routeTest := gin.Default()
	routeTest.DELETE("/say", Handler)
	routeTest.ServeHTTP(writer, request)

	response := writer.Result()
	actualResponseBody, _ := ioutil.ReadAll(response.Body)

	if expectedResponseBody != string(actualResponseBody) {
		t.Errorf("expect\n%s\nbut it got\n%s", expectedResponseBody, actualResponseBody)
	}
}
