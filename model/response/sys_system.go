package response

import "go-chains/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
