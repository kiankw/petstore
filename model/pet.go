package model

type Pet struct {
	Id   int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column" json:"name"`
}
