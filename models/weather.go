package models

type Weathers struct{
	GormModel
	Water int `gorm:"not null;" json:"water" form:"water" valid:"required"` 
	Wind int `gorm:"not null;" json:"wind" form:"wind" valid:"required"`
}