package entity

type Category struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"category_name" db:"category_name"`
}
