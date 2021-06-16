package database

import (
	"time"
)

type Rating struct {
	Rate     int64
	UpdateAt time.Time
}

func (r *Rating) GetRate() int64 {
	return r.Rate
}

func (r *Rating) GetUpdateAt() time.Time {
	return r.UpdateAt
}
