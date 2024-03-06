package models

type CustomFieldValue struct {
	ID       int32  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type CustomFieldValues []CustomFieldValue

func (cfv *CustomFieldValue) GetCustomFieldName() string {
	return cfv.Name
}
func (cfv *CustomFieldValue) GetCustomFieldCategory() string {
	return cfv.Category
}

func (cfv *CustomFieldValue) TableName() string {
	return "CustomFieldValues"
}
