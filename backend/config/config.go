package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"latihan-assessment/migration"
)

func ConnectToDatabase() *gorm.DB {
	//err := godotenv.Load()
	//
	//dbUser : os.Getenv("DB_USERNAME")
	//dbPass : os.Getenv("DB_PASSWORD")
	//dbHost : os.Getenv("DB_HOST")
	//dbName : os.Getenv("DB_NAME")

	dsn := "root:root@tcp(localhost)/assessment_excellent-echo?charset=utf8mb4&parseTime=True&loc=Local"


	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&migration.User{})
	db.AutoMigrate(&migration.Books{})

	return db

}