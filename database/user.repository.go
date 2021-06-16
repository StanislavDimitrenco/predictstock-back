package database

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo UserRepo) Create(user *User) *User {
	repo.db.Create(&user)
	return user
}

func (repo UserRepo) Save(user *User) *User {
	repo.db.Save(user)
	return user
}

func (repo UserRepo) Update(id int64, user *User) *User {
	repo.db.Where("id=", id).Updates(user)
	return user
}

func (repo UserRepo) Delete(user *User) {
	repo.db.Delete(user)
}

func (repo UserRepo) FindById(id int64) *User {
	var user User
	repo.db.First(&user, id)
	return &user
}

func (repo UserRepo) GetAll() (users []User) {
	repo.db.Order("id desc").Find(&users)
	return users
}

func (repo UserRepo) FindBy(param map[string]interface{}) (*User, bool) {
	var user User
	result := repo.db.Where(param).First(&user)
	return &user, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo UserRepo) Filter(param map[string]interface{}) []User {
	var users []User
	repo.db.Where(param).Find(&users)
	return users
}

func (repo UserRepo) MarkAsPaid(user *User, months int) *User {
	paidUntil := time.Now().AddDate(0, months, 0)
	if user.IsPaid && user.GetPaidUntil().After(time.Now()) {
		paidUntil = user.PaidUntil.AddDate(0, months, 0)
	}
	user.IsPaid = true
	user.PaidUntil = paidUntil
	repo.db.Save(user)

	return user
}

func (repo UserRepo) CheckPaidUntil(user *User) *User {
	if time.Now().After(user.GetPaidUntil()) && user.IsPaid == true {
		user.IsPaid = false
		repo.db.Save(&user)
	}

	return user
}
