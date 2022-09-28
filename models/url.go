package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortID string
	LongURL string
}

func init() {
	db.AutoMigrate(&URL{})
}

func (url *URL) Create() *gorm.DB {
	return db.Create(&url)
}

func (url *URL) ReadByLongURL() *gorm.DB {
	return db.Where("long_url = ?", url.LongURL).First(&url)
}

func (url *URL) ReadByShortID() *gorm.DB {
	return db.Where("short_id = ?", url.ShortID).First(&url)
}
