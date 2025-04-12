package entity

type User struct {
	ID       int `xorm:"pk autoincr"`
	Name     string
	Email    string
	Password string
}
