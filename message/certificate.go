package message

import "time"

type Certificate struct {
	Id             string    `json:"id"`
	AuthorityKeyId string    `json:"aki"`
	SerialNumber   string    `json:"serialNumber"`
	Pem            string    `json:"pem"`
	NotBefore      time.Time `json:"notBefore"`
	NotAfter       time.Time `json:"notAfter"`
}
