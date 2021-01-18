package response

import "go-chains/model"

type SysAPIResponse struct {
	Api model.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []model.SysApi `json:"apis"`
}
