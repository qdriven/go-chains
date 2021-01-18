package request

import "go-chains/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}