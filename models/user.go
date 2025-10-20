package models

import "time"

type User struct {
	ID                   int64     `json:"id"`
	Email                string    `json:"email"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Password             string    `json:"password"`
	CreatedAt            time.Time `json:"created_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	ModifiedAt           time.Time `json:"modified_at"`
	LastConnectedAt      time.Time `json:"last_connected_at"`
	ResetPasswordExpires string    `json:"reset_password_expires"`
	ResetPasswordToken   string    `json:"reset_password_token"`
}
