package database

import (
	"fmt"
	"goal-tracking/config"
	"goal-tracking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectDb() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

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
		&models.School{},
		&models.ClassroomTeacher{},
		&models.ClassroomStudent{},
		&models.ClassroomDirectory{},
	)
	if err != nil {
		panic("failed migrate")
	}
	fmt.Println("Veritabanı migrate edildi.")
}
