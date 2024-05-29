package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	dsn := "host=localhost user=admin password=root dbname=test_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

var DB *gorm.DB

func InitDatabase() {
	DB = connectDb()
	DB.AutoMigrate(&Employee{}) // SYNKAR databas med kod KOD = sanningen
	// var antal int64
	// DB.Model(&Employee{}).Count(&antal) // Seed

	// DB.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"})
	// DB.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
	// DB.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})

}
