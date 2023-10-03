package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aarontanjaya/finance-app-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func getLogger() logger.Interface {
	err := recover()
	if err != nil {
		log.Println(err)
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() error {
	c := config.Config.DBConfig
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})

	if err != nil {
		return err
	}

	return err
}
