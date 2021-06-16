package database

import (
	"errors"
	"gorm.io/gorm"
)

type RatingRepo struct {
	db *gorm.DB
}

func NewRatingRepo(db *gorm.DB) *RatingRepo {
	return &RatingRepo{db: db}
}

func (repo RatingRepo) Create(rating *Rating) *Rating {
	repo.db.Create(rating)
	return rating
}

func (repo RatingRepo) Save(rating *Rating) *Rating {
	repo.db.Save(rating)
	return rating
}

func (repo RatingRepo) Update(id int64, rating *Rating) *Rating {
	repo.db.Where("id=", id).Updates(rating)
	return rating
}

func (repo RatingRepo) Delete(rating *Rating) {
	repo.db.Delete(rating)
}

func (repo RatingRepo) FindById(id int) *Rating {
	var rating *Rating
	repo.db.First(rating, id)
	return rating
}

func (repo RatingRepo) FindBy(param map[string]interface{}) (*Rating, bool) {
	var rating Rating
	result := repo.db.Where(param).First(&rating)
	return &rating, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo RatingRepo) Filter(param map[string]interface{}) []Rating {
	var rating []Rating
	repo.db.Where(param).Find(&rating)
	return rating
}
