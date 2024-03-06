package models

type ObjectCustomFieldValue struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Content     string `json:"content"`
	CustomField uint   `json:"customfield"`
}

// CustomFields ...
type ObjectCustomFieldValues []ObjectCustomFieldValue

func (c *ObjectCustomFieldValue) TableName() string {
	// ORM mapping table name, this is default
	return "ObjectCustomFieldValues"
}
