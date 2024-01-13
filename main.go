package main

import (
	"fmt"

	"github.com/aarontanjaya/finance-app-backend/db"
	"github.com/aarontanjaya/finance-app-backend/repository"
	"github.com/aarontanjaya/finance-app-backend/utils"
	fdc "github.com/aarontanjaya/finance-data-collector/collector"
	fe "github.com/aarontanjaya/finance-data-collector/entity"
)

func main() {
	db.Connect()
	Coll, _ := fdc.NewPasardanaSnaphotCollector()
	res, err := Coll.GetAll(fe.Pagination{
		Start:     1,
		Length:    100,
		SortField: "",
		SortOrder: "ASC",
		KeyWord:   "",
	})
	db.Connect()
	fmt.Println(utils.PrettifyObject(res))
	repo := repository.NewFundRepository(repository.FundRepositoryConfig{
		DB: db.Get(),
	})

	_, err = repo.UpsertFunds(utils.MapCollectorFundSnapshot(res))
	fmt.Println(err)
}
