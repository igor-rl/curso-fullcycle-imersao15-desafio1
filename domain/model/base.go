package model

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID int32 `json:"id" gorm:"primaryKey;autoIncrement"`
}
