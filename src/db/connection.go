package db
var db *gorm.DB

func Connection() *gorm.DB {
	return db
}