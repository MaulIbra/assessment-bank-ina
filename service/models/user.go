package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	BaseModel
}

func (User) TableName() string {
	return "users"
}
