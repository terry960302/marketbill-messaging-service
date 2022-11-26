package datastore

import (
	"fmt"
	"log"
	"marketbill-messaging-service/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresql() (*gorm.DB, error) {
	DSN := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PW") +
		" port=" + os.Getenv("DB_PORT") +
		" database=" + fmt.Sprint(os.Getenv("DB_NAME")) +
		" sslmode=disable" +
		" TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Print(err)
		return nil, err
	}

	db.AutoMigrate(&models.SendSmsLogs{})

	return db, nil
}
