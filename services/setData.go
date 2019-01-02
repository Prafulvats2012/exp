package services

import (
	"exp/daos"
	"exp/dtos"
)

type setData struct {
	r *daos.RedisDao
}

func New() *setData {
	return &setData{
		r: daos.New(),
	}
}

func (sd *setData) SetData(req *dtos.SetDataReq) bool {
	flag, err := sd.r.SetData(req)
	if err != nil {
		return false
	}
	return flag
}
