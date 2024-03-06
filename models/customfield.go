package models

type CustomField struct {
	ID   int32  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type CustomFields []CustomField

func (c *CustomField) TableName() string {
	return "CustomFields"
}
