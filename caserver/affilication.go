package caserver

import "github.com/hyperledger/fabric-ca/api"

func AllAffiliations() (*api.AffiliationInfo, error) {
	admin, err := getAdminIdentity()
	if err != nil {
		return nil, err
	}
	resp, err := admin.GetAllAffiliations("")
	if err != nil {
		return nil, err
	}
	return &resp.AffiliationInfo, nil
}

// add affiliation
func AddAffiliation(req *api.AddAffiliationRequest) error {
	admin, err := getAdminIdentity()
	if err != nil {
		return err
	}
	_, err = admin.AddAffiliation(req)
	return err
}
