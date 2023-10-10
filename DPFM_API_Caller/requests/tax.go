package requests

type Tax struct {
	Product                  string `json:"Product"`
	BusinessPartner          *int   `json:"BusinessPartner"`
	Country                  string `json:"Country"`
	ProductTaxCategory              string `json:"ProductTaxCategory"`
	ProductTaxClassification string `json:"ProductTaxClassification"`
}
