package validator

import (
	utils "go-chains/utils"
)

var (
	ChainMetaValidator = utils.Rules{"ChainId": {utils.NotEmpty()}, "ChainName": {utils.NotEmpty()},
		"ChainType":{utils.NotEmpty()}}
)

