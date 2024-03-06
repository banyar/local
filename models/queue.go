package models

type Queue struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:Name"`
}

type Queues []Queue

func (e *Queue) TableName() string {
	return "Queues"
}
