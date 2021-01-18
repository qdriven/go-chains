package request

import "go-chains/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
