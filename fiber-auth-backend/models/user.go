package models

import "time"

// User merepresentasikan tabel users di database.
// Password disimpan dalam bentuk hash (bcrypt), tidak pernah plain text.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" agar tidak pernah ikut ter-serialize ke JSON response
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RegisterRequest adalah bentuk payload saat user melakukan registrasi.
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// LoginRequest adalah bentuk payload saat user melakukan login.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse adalah bentuk response setelah register/login berhasil.
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
