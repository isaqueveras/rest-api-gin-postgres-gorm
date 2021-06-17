package migrations

import (
	"github.com/isaqueveras/rest-api-gin-postgres-gorm/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
