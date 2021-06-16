package database

import (
	"time"
)

type Stats struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    int64
	Symbol    string
	CreatedAt time.Time
}

func (s *Stats) GetId() int64 {
	return s.Id
}

func (s *Stats) GetUserId() int64 {
	return s.UserId
}

func (s *Stats) GetSymbol() string {
	return s.Symbol
}

func (s *Stats) GetCreatedAt() time.Time {
	return s.CreatedAt
}
