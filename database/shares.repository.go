package database

import (
	"errors"
	"gorm.io/gorm"
)

type SharesRepo struct {
	db *gorm.DB
}

func NewSharesRepo(db *gorm.DB) *SharesRepo {
	return &SharesRepo{db: db}
}

func (repo SharesRepo) Create(shares *Share) *Share {
	repo.db.Create(shares)
	return shares
}

func (repo SharesRepo) Save(share *Share) *Share {
	repo.db.Save(share)
	return share
}

func (repo SharesRepo) Update(id int64, share *Share) *Share {
	repo.db.Where("id = ?", id).Updates(share)
	return share
}

func (repo SharesRepo) Delete(share *Share) {
	repo.db.Delete(share)
}

func (repo SharesRepo) FindById(id int) *Share {
	var share *Share
	repo.db.First(share, id)
	return share
}

func (repo SharesRepo) FindBy(param map[string]interface{}) (*Share, bool) {
	var share Share
	result := repo.db.Where(param).First(&share)
	return &share, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo SharesRepo) Filter(param map[string]interface{}) []Share {
	var shares []Share
	repo.db.Where(param).Find(&shares)
	return shares
}
