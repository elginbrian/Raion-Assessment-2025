package domain

import "time"

type User struct {
    ID           string    `json:"id"`          
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`             
    CreatedAt    time.Time `json:"created"`
    UpdatedAt    time.Time `json:"updated"`
}