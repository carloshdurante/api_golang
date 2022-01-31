package migrations

import (
	"github.com/carloshdurante/api_golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
