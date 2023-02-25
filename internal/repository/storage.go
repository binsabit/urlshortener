package repository

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	srtLen        = 10
)

//interface
type UserStorage interface {
	CreateUser(*User) (int64, error)
	DeleteUser(int64) error
	GetUserByID(int64) (*User, error)
	UpdateUser(int64, User) error
}

type LinkStorage interface {
	SaveURL(int64, *Link) error
	DeleteURL(int64) error
	UpdateURL(Link) error
}

//models user , url
type Link struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	SrtURL    string    `json:"srtURL"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires"`
}

func NewLink(URL string, expires time.Time) *Link {
	link := Link{
		URL:       URL,
		ExpiresAt: expires,
	}
	link.shorten(srtLen)

	return &link

}

func (l *Link) shorten(len int) {
	//function to make URL short
	l.SrtURL = randString(10)
}

//generates random strng of length n
func randString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

type User struct {
	ID        int64     `json:"userID"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
