package models

type Responsible struct {
	ID   uint   `gorm:"column:id_Responsible;primaryKey" json:"id_Responsible"`
	Name string `gorm:"column:Name" json:"Name"`
}
