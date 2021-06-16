package database

import (
	"errors"
	"gorm.io/gorm"
)

type StatsRepo struct {
	db *gorm.DB
}

func NewStatsRepo(db *gorm.DB) *StatsRepo {
	return &StatsRepo{db: db}
}

func (repo StatsRepo) Create(stats *Stats) *Stats {
	repo.db.Create(stats)
	return stats
}

func (repo StatsRepo) Save(stats *Stats) *Stats {
	repo.db.Save(stats)
	return stats
}

func (repo StatsRepo) Update(id int64, stats *Stats) *Stats {
	repo.db.Where("id=", id).Updates(stats)
	return stats
}

func (repo StatsRepo) Delete(stats *Stats) {
	repo.db.Delete(stats)
}

func (repo StatsRepo) FindById(id int) *Stats {
	var stats *Stats
	repo.db.First(stats, id)
	return stats
}

func (repo StatsRepo) FindBy(param map[string]interface{}) (*Stats, bool) {
	var stats Stats
	result := repo.db.Where(param).First(&stats)
	return &stats, errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (repo StatsRepo) Filter(param map[string]interface{}) []Stats {
	var stats []Stats
	repo.db.Where(param).Find(&stats)
	return stats
}
