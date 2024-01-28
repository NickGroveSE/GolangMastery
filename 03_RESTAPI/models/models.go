package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Username  string
	Bio       sql.NullString
	Location  sql.NullString
	CreatedAt time.Time
	CheckIns  []CheckIn
}

type CheckIn struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	BeerID      uint
	Rating      sql.NullFloat64
	Description sql.NullString
	CreatedAt   time.Time
	User        User
	Beer        Beer
}

type Brewery struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	Name        string
	BreweryType string
	Location    string
	CreatedAt   time.Time
	Beers       []Beer
}

type Beer struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	BreweryID uint
	Style     string
	ABV       float64
	IBU       sql.NullFloat64
	CreatedAt time.Time
	Brewery   Brewery
	CheckIns  []CheckIn
}
