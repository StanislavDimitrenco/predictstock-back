package database

import "time"

type UserHistory struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	Message   string
	CreatedAt time.Time
}

func (uh *UserHistory) GetId() int64 {
	return uh.ID
}

func (uh *UserHistory) GetUserID() int64 {
	return uh.UserID
}

func (uh *UserHistory) GetMessage() string {
	return uh.Message
}

func (uh *UserHistory) GetCreatedAt() time.Time {
	return uh.CreatedAt
}
