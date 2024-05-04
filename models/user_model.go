package models

type TUser struct {
	TBaseModel

	FirstName string `gorm:"type:string;size:15;null"`

	LastName string `gorm:"type:string;size:25;null"`

	Email string `gorm:"type:string;size:64;not null;unique"`

	Password string `gorm:"type:string;size:64;not null"`
}
