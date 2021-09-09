package entity

import "time"

type LogEntity struct {
	ID   int64  `gorm:"primary_key:auto_increment"`
	Name string `gorm:"type:varchar(100)"`
	Date time.Time
}
