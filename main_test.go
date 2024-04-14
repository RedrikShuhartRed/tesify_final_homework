package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerWhenOk тестирует корректность запроса, код ответа 200 тело ответа не пустое.
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	require.NotEmpty(t, responseRecorder.Body)

}

// TestMainHandlerWhenWrongCityValue тестирует запрос неподдерживаемого города.
func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=Imperial-City", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Result().StatusCode)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}

// TestMainHandlerWhenCountMoreThanTotal тестирует запрос с числом кафе больше существующего.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := len(cafeList["moscow"])
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	list := strings.Split(responseRecorder.Body.String(), ",")

	require.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	assert.Equal(t, totalCount, len(list))

}
