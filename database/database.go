package database

import (
	"fmt"
	"goal-tracking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectDb() {
	var err error
	dsn := "host=localhost user=aydin password=1 dbname=goaltracking port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("veritabanına bağlanılamadı.")
		panic("failed : veritabanına bağlanılamadı.")
	}

	fmt.Println("Veritabanına başarıyla bağlanıldı.")

	err = db.AutoMigrate(
		&models.User{},
		&models.Goal{},
		&models.Classroom{},
		&models.GoalDoneStatus{},
	)
	if err != nil {
		panic("failed migrate")
	}
	fmt.Println("Veritabanı migrate edildi.")
}
