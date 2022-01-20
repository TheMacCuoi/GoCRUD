package model
type User struct {
	ID   int    `gorm:"primarykey" json:"id"`
	Name string `json:"name"`
}

func (User) TableName() string {
	return "Users"
}
