package initializers

import "github.com/hardytee1/FP_PBKK_Go/Backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Blog{})
}