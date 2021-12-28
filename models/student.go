package models

type Student struct {
	Id       int64     `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	Lecture  []Lecture `json:"lectures"`
}
