package db

import (
	"cademo/config"
	"fmt"

	"xorm.io/xorm"
)

type DBClient struct {
	engine *xorm.Engine
}

func NewDBClient() (*DBClient, error) {
	dbsrc := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.C.GetString("cadb.host"),
		config.C.GetInt("cadb.port"),
		config.C.GetString("cadb.user"),
		config.C.GetString("cadb.password"),
		config.C.GetString("cadb.dbname"),
		config.C.GetString("cadb.sslmode"),
	)
	eg, err := xorm.NewEngine("postgres", dbsrc)
	if err != nil {
		return nil, err
	}
	client := &DBClient{
		engine: eg,
	}
	return client, nil
}

func (d *DBClient) QueryIdentityStates() ([]Users, error) {
	users := []Users{}
	err := d.engine.Find(&users)
	return users, err
}

func (d *DBClient) DeleteCertificate(cert *Certificates) error {
	_, err := d.engine.Delete(cert)
	return err
}

func (d *DBClient) QueryCredentials() ([]Credentials, error) {
	creds := []Credentials{}
	err := d.engine.Find(&creds)
	return creds, err
}

func (d *DBClient) DeleteCredential(cred *Credentials) error {
	_, err := d.engine.Delete(cred)
	return err
}
