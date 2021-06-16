package database

import (
	"time"
)

type Share struct {
	//gorm.Model
	Id        int64 `gorm:"primaryKey"`
	Name      string
	Symbol    string
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Share) GetId() int64 {
	return s.Id
}

func (s *Share) GetName() string {
	return s.Name
}

func (s *Share) GetSymbol() string {
	return s.Symbol
}

func (s *Share) GetRating() int {
	return s.Rating
}

func (s *Share) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *Share) GetUpdateAt() time.Time {
	return s.UpdatedAt
}
