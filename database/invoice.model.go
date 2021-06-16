package database

import (
	"gorm.io/gorm"
	"time"
)

type Invoice struct {
	ID          int64 `gorm:"primaryKey"`
	UserID      int64
	User        User `gorm:"constraint:OnDelete:CASCADE;"`
	Status      string
	Price       int64
	Months      int
	MessageId   int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	PaidAt      time.Time
}

func (i *Invoice) GetId() int64 {
	return i.ID
}

func (i *Invoice) GetUserId() int64 {
	return i.UserID
}

func (i *Invoice) GetUser() User {
	return i.User
}

func (i *Invoice) GetStatus() string {
	return i.Status
}

func (i *Invoice) GetPrice() int64 {
	return i.Price
}

func (i *Invoice) GetMonths() int {
	return i.Months
}

func (i *Invoice) GetMessageId() int {
	return i.MessageId
}

func (i *Invoice) GetDescription() string {
	return i.Description
}

func (i *Invoice) GetCreatedAt() time.Time {
	return i.CreatedAt
}

func (i *Invoice) GetUpdatedAt() time.Time {
	return i.UpdatedAt
}

func (i *Invoice) GetPaidAt() time.Time {
	return i.PaidAt
}
