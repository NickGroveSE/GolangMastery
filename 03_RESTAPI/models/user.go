package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Bio       sql.NullString
	Location  sql.NullString
	CreatedAt time.Time
}
