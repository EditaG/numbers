package server

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"session_manager/model"
	"testing"
)

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

func TestCreateSession(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/session/",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusCreated, response.Code)
	session := model.SessionOutput{}
	json.Unmarshal(response.Body.Bytes(), &session)
	_, err := uuid.Parse(session.Id)
	assert.Nil(t, err)
}

func TestGetSession(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/session/",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusCreated, response.Code)
	session := model.SessionOutput{}
	json.Unmarshal(response.Body.Bytes(), &session)
	_, err := uuid.Parse(session.Id)
	assert.Nil(t, err)

	getSessionResponse := performRequest(
		router,
		"GET",
		"/session/"+session.Id,
		"",
		map[string]string{})

	assert.Equal(t, http.StatusOK, getSessionResponse.Code)
	getSession := model.SessionOutput{}
	json.Unmarshal(getSessionResponse.Body.Bytes(), &getSession)
	assert.Equal(t, getSession.Id, session.Id)
}

func TestGetNonExistingSession(t *testing.T) {
	initTestEnv()
	router := setRouter()

	response := performRequest(
		router,
		"GET",
		"/session/fake",
		"",
		map[string]string{})

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestDeleteSession(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"POST",
		"/session/",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusCreated, response.Code)
	session := model.SessionOutput{}
	json.Unmarshal(response.Body.Bytes(), &session)
	_, err := uuid.Parse(session.Id)
	assert.Nil(t, err)

	deleteSessionResponse := performRequest(
		router,
		"DELETE",
		"/session/"+session.Id,
		"",
		map[string]string{})

	assert.Equal(t, http.StatusOK, deleteSessionResponse.Code)

	getSessionResponse := performRequest(
		router,
		"GET",
		"/session/"+session.Id,
		"",
		map[string]string{})

	assert.Equal(t, http.StatusNotFound, getSessionResponse.Code)
}

func Test404(t *testing.T) {
	initTestEnv()
	router := setRouter()
	response := performRequest(
		router,
		"GET",
		"/fake",
		"",
		map[string]string{})
	assert.Equal(t, http.StatusNotFound, response.Code)
}
