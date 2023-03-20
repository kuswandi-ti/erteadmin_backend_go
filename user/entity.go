package user

import (
	"erteadmin_backend/lingkungan"
	"time"
)

type User struct {
	ID           int
	LingkunganID int
	Email        string
	Password     string
	Avatar       string
	Role         string
	Lingkungan   lingkungan.Lingkungan
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
