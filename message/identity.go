package message

import "github.com/hyperledger/fabric-ca/api"

type IdentityInfoExt struct {
	api.IdentityInfo
	State int `json:"state"`
}
