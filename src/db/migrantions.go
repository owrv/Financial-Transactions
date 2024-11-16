package db

import (
	"Financial-Transactions/src/models"
	"fmt"
)

func CreateDatabase(){
	conn := Connection()
	fmt.Println("conn.HasTable(&models.Users{}): ", conn.HasTable(&models.Users{}))
	if !conn.HasTable(&models.Users{}) {
		_ = conn.DropTableIfExists(&models.Users{})

		_ = conn.AutoMigrate(&models.Users{})
	}

}
