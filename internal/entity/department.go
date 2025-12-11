package entity

type Department struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}