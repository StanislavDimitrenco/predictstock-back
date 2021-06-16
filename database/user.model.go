package database

import (
	"time"
)

type User struct {
	ID         int64 `gorm:"primaryKey"`
	TelegramId int
	Username   string
	Name       string
	Lastname   string
	IsPaid     bool
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PaidUntil  time.Time `gorm:"default:null"`
}

func (u *User) GetId() int64 {
	return u.ID
}

func (u *User) GetTelegramId() int {
	return u.TelegramId
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetLastname() string {
	return u.Lastname
}

func (u *User) GetIsPaid() bool {
	return u.IsPaid
}

func (u *User) GetIsActive() bool {
	return u.IsActive
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdateAt() time.Time {
	return u.UpdatedAt
}

func (u *User) GetPaidUntil() time.Time {
	return u.PaidUntil
}
