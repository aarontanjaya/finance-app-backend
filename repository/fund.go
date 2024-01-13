package repository

import (
	"github.com/aarontanjaya/finance-app-backend/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FundRepositoryConfig struct {
	DB *gorm.DB
}

type fundRepositoryImpl struct {
	db *gorm.DB
}

type FundRepository interface {
	UpsertFunds(fnd []entity.FundSnapshot) ([]entity.FundSnapshot, error)
}

func NewFundRepository(fu FundRepositoryConfig) FundRepository {
	return &fundRepositoryImpl{
		db: fu.DB,
	}
}

func (r *fundRepositoryImpl) UpsertFunds(fnd []entity.FundSnapshot) ([]entity.FundSnapshot, error) {
	err := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
		Columns: []clause.Column{{
			Name: "id",
		}, {
			Name: "last_update",
		}},
	}).Create(&fnd).Error
	return fnd, err
}
