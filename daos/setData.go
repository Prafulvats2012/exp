package daos

import (
	"encoding/json"
	"exp/daos/db"
	"exp/dtos"
	"fmt"
)

type RedisDao struct {
	r *db.RedisConn
}

func New() *RedisDao {
	return &RedisDao{
		r: db.New(),
	}
}

func (rd *RedisDao) SetData(req *dtos.SetDataReq) (bool, error) {
	fmt.Print("Func called")
	byteObj, err := json.Marshal(req)
	if err != nil {
		fmt.Errorf("Error while marshalling Initial constraints object to string", err)
	}
	fmt.Print("Func called1")
	err = rd.r.GetQueryer().HSet("testData", fmt.Sprintf("testkey"), string(byteObj)).Err()

	if err != nil {
		fmt.Errorf("Unable to set value in redis", err)
		return false, err
	}

	return true, nil
}
