package response

import "go-chains/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
