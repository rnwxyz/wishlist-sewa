package database

import (
	"fmt"

	"github.com/rnwxyz/wishlist-sewa/config"
	"github.com/rnwxyz/wishlist-sewa/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env.DB_USERNAME,
		config.Env.DB_PASSWORD,
		config.Env.DB_ADDRESS,
		config.Env.DB_NAME,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Owner{},
		model.Product{},
	)
}
