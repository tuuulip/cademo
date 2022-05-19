package caserver

import "github.com/hyperledger/fabric-ca/api"

func AllAffiliations() (*api.AffiliationInfo, error) {
	id, err := getAdminIdentity()
	if err != nil {
		return nil, err
	}
	resp, err := id.GetAllAffiliations("")
	if err != nil {
		return nil, err
	}
	return &resp.AffiliationInfo, nil
}
