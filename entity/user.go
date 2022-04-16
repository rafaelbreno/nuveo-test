package entity

import "time"

type (
	// User stores all values
	// related to a client.
	User struct {
		UUID      string    `gorm:"type:uuid;default:gen_random_uuid()"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func (u *User) UpdateFields(newUser User) {
	u.Name = newUser.Name
	u.Address = newUser.Address
}
