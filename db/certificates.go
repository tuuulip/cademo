package db

type Certificates struct {
	Id             string `json:"id" xorm:"id"`
	SerialNumber   string `json:"serialNumber" xorm:"serial_number"`
	AuthorityKeyId string `json:"aki" xorm:"authority_key_identifier"`
}
