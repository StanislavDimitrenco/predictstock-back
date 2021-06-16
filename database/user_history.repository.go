package database

import (
	"gorm.io/gorm"
	"strconv"
)

type UserHistoryRepo struct {
	db *gorm.DB
}

func NewUserHistoryRepo(db *gorm.DB) *UserHistoryRepo {
	return &UserHistoryRepo{db: db}
}

func (repo UserHistoryRepo) Create(userHistory *UserHistory) *UserHistory {
	repo.db.Create(&userHistory)
	return userHistory
}

func (repo UserHistoryRepo) FindAllRecordsByUsersID(userId string) (usersHistory []UserHistory) {
	userIdInt, _ := strconv.Atoi(userId)
	repo.db.Where("user_id = ?", userIdInt).Find(&usersHistory)
	return usersHistory
}
