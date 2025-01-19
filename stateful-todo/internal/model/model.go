package model

type Task struct {
	Id          int    `json:"id,omitempty" db:"id"`
	Title       string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
}
