package entity

type Role struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Name         string     `json:"name"`
	DepartmentId string     `json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentId;references:ID"`
}