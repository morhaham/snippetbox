package models

import (
	"errors"
	"time"
)

var (
	ErrorNoRecord        = errors.New("models: no matching record found")
	ErInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail    = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
