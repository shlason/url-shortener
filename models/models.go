package models

import (
	"fmt"

	"github.com/shlason/url-shortener/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?%s",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Protocol,
		configs.Database.Host,
		configs.Database.Name,
		configs.Database.Options,
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&URL{})
	db = d
}
