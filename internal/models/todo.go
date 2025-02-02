package models

type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Text  string `gorm:"type:text;not null;index" json:"text"`
	Done  bool   `json:"done"`
	Done2 bool   `json:"done"`
}
