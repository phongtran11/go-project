package models

type User struct {
	BaseModel

	FirstName string `gorm:"type:string;size:15;null;default:null"`

	LastName string `gorm:"type:string;size:25;null;default:null"`

	Email string `gorm:"type:string;size:64;not null;unique"`

	Password string `gorm:"type:string;size:64;not null"`
}
