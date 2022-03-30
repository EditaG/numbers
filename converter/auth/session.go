package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Session struct {
	ID        string `json:"id"`
	ExpiresAt string `json:"expiresAt"`
}

func CreateSession() (*Session, error) {
	client := getClient()
	postData := bytes.NewBuffer([]byte(``))
	resp, err := client.Post(os.Getenv("SESSION_API_URL"), "application/json", postData)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusCreated {
		return getSessionFromResponse(resp)
	}

	return nil, errors.New("failed to create session")
}

func GetSession(id string) (*Session, error) {
	client := getClient()
	resp, err := client.Get(os.Getenv("SESSION_API_URL") + id)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return getSessionFromResponse(resp)
	}

	return nil, errors.New("failed to retrieve session")
}

func DeleteSession(id string) (bool, error) {
	client := getClient()
	req, err := http.NewRequest("DELETE", os.Getenv("SESSION_API_URL")+id, nil)
	if err != nil {
		return false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, errors.New("failed to delete session")
}

func getClient() http.Client {
	return http.Client{Timeout: 10 * time.Second}
}

func getSessionFromResponse(response *http.Response) (*Session, error) {
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var session *Session
	err = json.Unmarshal(bodyBytes, &session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
