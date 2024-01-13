package entity

import (
	"time"

	"gorm.io/gorm"
)

type FundSnapshot struct {
	Source       int     `json:"source"`
	ID           int     `json:"id" gorm:"column:id"`
	ISINCode     string  `json:"ISINCode"`
	Name         string  `json:"name"`
	IMName       string  `json:"IMName"`
	IMFee        string  `json:"IMFee"`
	Type         string  `json:"type"`
	TypeId       int     `json:"typeId"`
	IndexFlag    bool    `json:"indexFlag"`
	ETFFlag      bool    `json:"ETFFlag"`
	ShariaFlag   bool    `json:"shariaFlag"`
	NAV          float64 `json:"NAV"`
	DailyRR      float64 `json:"dailyRR" gorm:"column:daily_rr"`
	MonthlyRR    float64 `json:"monthlyRR" gorm:"column:monthly_rr"`
	YTDRR        float64 `json:"YTDRR" gorm:"column:ytd_rr"`
	YearlyRR     float64 `json:"yearlyRR" gorm:"column:yearly_rr"`
	AUM          float64 `json:"AUM"`
	AUMUpdatedAt uint64  `json:"AUMUpdatedAt"`
	Units        float64 `json:"units"`
	LastUpdate   uint64  `json:"lastUpdate"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
