package models

type ImageModel struct {
	Model
	Filename string `json:"filename"`
	Content  string `gorm:"type:text" json:"content"` //base64
}
