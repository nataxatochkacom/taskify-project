package models

type Task struct {
    ID           uint   `gorm:"column:id_Task;primaryKey" json:"id_Task"`
    Name         string `gorm:"column:Name" json:"Name"`
    Description  string `gorm:"column:Description" json:"Description"`
    Start_date   string `gorm:"column:Start_date" json:"Start_date"`
    End_date     string `gorm:"column:End_date" json:"End_date"`
    IsOverdue    bool   `gorm:"column:Is_overdue" json:"Is_overdue"`
    ProjectID    uint   `gorm:"column:id_Project" json:"id_Project"`
    Responsible  string `gorm:"-" json:"Responsible,omitempty"` // опционально
	Priority string `gorm:"column:Priority" json:"Priority"`
}
