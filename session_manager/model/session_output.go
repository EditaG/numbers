package model

import "time"

type SessionOutput struct {
	Id        string    `json:"id"`
	ExpiresAt time.Time `json:"expiresAt"`
}
