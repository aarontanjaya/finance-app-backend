package jobs

import (
	"fmt"

	"github.com/aarontanjaya/finance-app-backend/repository"
	"github.com/aarontanjaya/finance-app-backend/utils"
	fdc "github.com/aarontanjaya/finance-data-collector/collector"
	fe "github.com/aarontanjaya/finance-data-collector/entity"
)

func GetAll(c fdc.FundSnapshotCollector, repo repository.FundRepository) {
	page := 1
	var result []fe.FundSnapshot
	for true {
		res, err := c.GetAll(fe.Pagination{
			Start:     page,
			Length:    1000,
			SortField: "",
			SortOrder: "ASC",
			KeyWord:   "",
		})

		if err != nil {
			fmt.Println("Get All Error", err)
			break
		}

		result = append(result, res...)

		if len(res) == 0 {
			break
		}
		page++
	}

	_, err := repo.UpsertFunds(utils.MapCollectorFundSnapshot(result))
	fmt.Println(err)
}
