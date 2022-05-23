package db

import "time"

type Credentials struct {
	Id     string    `json:"id" xorm:"id"`
	Cred   string    `json:"cred" xorm:"cred"`
	Status string    `json:"status" xorm:"status"`
	Expiry time.Time `json:"expiry" xorm:"expiry"`
}
