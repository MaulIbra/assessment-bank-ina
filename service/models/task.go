package models

type Task struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	BaseModel
}

func (Task) TableName() string {
	return "tasks"
}
