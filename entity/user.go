package entity

import "time"

type (
	// User stores all values
	// related to a client.
	User struct {
		UUID      string    `json:"uuid"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
