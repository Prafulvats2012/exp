package services

import (
	"exp/daos"
	"exp/dtos"
)

func SetData(req *dtos.SetDataReq) bool {
	return daos.SetData(req)
}
