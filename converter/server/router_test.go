package server

import (
	"bytes"
	"converter/auth"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var client = http.Client{Timeout: 10 * time.Second}

func performRequest(r http.Handler, method, path, body string, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	for header, headerValue := range headers {
		req.Header.Add(header, headerValue)
	}

	writer := httptest.NewRecorder()
	r.ServeHTTP(writer, req)
	return writer
}

func initTestEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		panic(err)
	}
}

func TestHealthCheck(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "GET", "/healthcheck", "", map[string]string{})
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "\"OK\"", response.Body.String())
}

func TestLoginNonexistentUser(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "POST", "/api/login",
		"{\n    \"username\": \"fake_user\",\n    \"password\" : \"Pa33m0rD*&!\"\n}", map[string]string{})
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestLoginBadPassword(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "POST", "/api/login",
		"{\n    \"username\": \"fake_user\",\n    \"password\" : \"nopassword\"\n}", map[string]string{})
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestLoginSuccess(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "POST", "/api/login",
		"{\n    \"username\": \"demo_user\",\n    \"password\" : \"Pa33m0rD*&!\"\n}", map[string]string{})
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestLoginMissingUsername(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "POST", "/api/login",
		"{\n    \"username\": \"demo_user\"\n}", map[string]string{})
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestLoginMissingPassword(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(router, "POST", "/api/login",
		"{\n    \"password\" : \"Pa33m0rD*&!\"\n}", map[string]string{})
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestLogoutMissingAuthHeader(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/logout",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestLogoutWrongAuthHeader(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/logout",
		"",
		map[string]string{
			"Authorization": "wrong",
		})
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestLogoutSuccessful(t *testing.T) {
	initTestEnv()

	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/logout",
		"",
		map[string]string{
			"Authorization": getSessionToken(),
		})
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestConvertMissingAuthHeader(t *testing.T) {
	initTestEnv()

	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/convert/roman",
		"{\"input\": \"IV\"}",
		map[string]string{})
	assert.Equal(t, http.StatusUnauthorized, response.Code)
}

func TestConvertMissingInput(t *testing.T) {
	initTestEnv()

	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/convert/roman",
		"",
		map[string]string{
			"Authorization": getSessionToken(),
		})
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestConvertInvalidInput(t *testing.T) {
	initTestEnv()

	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/convert/roman",
		"{\"input\": \"a\"}",
		map[string]string{
			"Authorization": getSessionToken(),
		})
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestConvertSuccess(t *testing.T) {
	initTestEnv()

	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/api/convert/roman",
		"{\"input\": \"IV\"}",
		map[string]string{
			"Authorization": getSessionToken(),
		})
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"output\":4}", response.Body.String())
}

func Test404(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"GET",
		"/api/login",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusNotFound, response.Code)
}

// todo mock this
func getSessionToken() string {
	resp, _ := client.Post(os.Getenv("SESSION_API_URL"), "application/json", bytes.NewBuffer([]byte(``)))
	bodyBytes, _ := io.ReadAll(resp.Body)

	var session *auth.Session
	json.Unmarshal(bodyBytes, &session)
	return session.ID
}
