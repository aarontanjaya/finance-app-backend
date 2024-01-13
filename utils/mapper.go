package utils

import (
	"github.com/aarontanjaya/finance-app-backend/entity"
	fe "github.com/aarontanjaya/finance-data-collector/entity"
)

func MapCollectorFundSnapshot(inp []fe.FundSnapshot) []entity.FundSnapshot {
	filteredInput := []fe.FundSnapshot{}
	for i := range inp {
		if inp[i].Id != 0 {
			filteredInput = append(filteredInput, inp[i])
		}
	}

	result := make([]entity.FundSnapshot, len(filteredInput))
	for idx, item := range filteredInput {
		result[idx] = entity.FundSnapshot{
			Source:       item.Source,
			ID:           item.Id,
			ISINCode:     "",
			Name:         item.Name,
			IMName:       item.IMName,
			IMFee:        item.IMFee,
			Type:         item.Type,
			TypeId:       item.TypeID,
			IndexFlag:    item.IsIndex,
			ETFFlag:      item.IsETF,
			ShariaFlag:   item.IsSharia,
			NAV:          item.NAV,
			DailyRR:      item.DailyRR,
			MonthlyRR:    item.MonthlyRR,
			YTDRR:        item.YTDRR,
			YearlyRR:     item.YearlyRR,
			AUM:          item.AUM,
			AUMUpdatedAt: item.AUMLastUpdate,
			LastUpdate:   item.LastUpdate,
		}
	}

	return result
}
