package main

import (
	"fmt"

	"github.com/aarontanjaya/finance-app-backend/db"
	"github.com/aarontanjaya/finance-app-backend/jobs"
	"github.com/aarontanjaya/finance-app-backend/repository"
	fdc "github.com/aarontanjaya/finance-data-collector/collector"
	"github.com/robfig/cron"
)

func main() {
	db.Connect()
	Coll, _ := fdc.NewPasardanaSnaphotCollector()
	db.Connect()

	repo := repository.NewFundRepository(repository.FundRepositoryConfig{
		DB: db.Get(),
	})
	cr := cron.New()
	cr.AddFunc("0 0 * * * *", func() {
		jobs.GetAll(Coll, repo)
	})
	cr.Start()
	fmt.Scanln()

}
