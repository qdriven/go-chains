package request

import "go-chains/model"

type SysDictionarySearch struct{
    model.SysDictionary
    PageInfo
}