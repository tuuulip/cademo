package db

type Users struct {
	ID    string `json:"id" xorm:"id"`
	State string `json:"state" xorm:"state"`
}
