package db

import (
	"fmt"
	"io"
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

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize file for logger")
	}

	logs := io.MultiWriter(os.Stdout, file)

	return logger.New(
		log.New(logs, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
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

func Get() *gorm.DB {
	return db
}
