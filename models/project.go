package models

type Project struct {
	ID          uint   `json:"id_Project" gorm:"column:id_Project;primaryKey"`
	Name        string `json:"Name" gorm:"column:Name"`
	Description string `json:"Description,omitempty" gorm:"column:Description"`
	StartDate   string `json:"Start_date,omitempty" gorm:"column:Start_date"`
	EndDate     string `json:"End_date,omitempty" gorm:"column:End_date"`
	IsOverdue   bool   `json:"Is_overdue" gorm:"column:Is_overdue"`
}

func (Project) TableName() string {
	return "Projects"
}
