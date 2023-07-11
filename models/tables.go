package models

import "gorm.io/gorm"

type Anime struct {
	gorm.Model
	URL                 string       `gorm:"type:varchar(255);not null"`
	Name                string       `gorm:"type:varchar(255);not null"`
	Image               string       `gorm:"type:varchar(255)"`
	LastReleasedEpisode string       `gorm:"type:varchar(255)"`
	IsSeasonOver        bool         `gorm:"default:false"`
	Subscribers         []Subscriber `gorm:"foreignKey:AnimeID"`
}

type Subscriber struct {
	gorm.Model
	TelegramID int64 `gorm:"not null"`
	AnimeID    uint
	Anime      Anime `gorm:"foreignKey:AnimeID"`
}
