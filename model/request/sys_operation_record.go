package request

import "go-chains/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
