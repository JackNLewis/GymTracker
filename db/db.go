package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB

	dsn = "jack:Password1@tcp(127.0.0.1:3306)/gymtracker?charset=utf8mb4&parseTime=True&loc=Local"
)

type Exercise struct {
	Exercise_ID int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Muscle      string `json:"muscle"`
}

func InitializeDB() {
	var err error
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetExercises(muscle string) []Exercise {
	var ex []Exercise

	//if not muscle is passed returned exercises for all muscle groups
	if muscle == "" {
		Database.Find(&ex)
		return ex
	}

	Database.Where("muscle = ?", muscle).Find(&ex)
	return ex
}
