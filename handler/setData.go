package handler

import (
	"encoding/json"
	"exp/dtos"
	"exp/library"
	"exp/services"
	"fmt"
	"lynk/dbLibrary/config/globals"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print("Ping request received!")
	rd := library.LogAndGetContext(w, r)
	res := dtos.SteDataRes{}
	res.Success = "true"
	library.WriteJSONResponse(res, http.StatusOK, rd)
}

func setData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	rd := library.LogAndGetContext(w, r)
	req := &dtos.SetDataReq{}
	err := decoder.Decode(req)
	if err != nil {
		globals.Logger.Error("Error while Decoding request params", err)
		library.WriteJSONMessage(err.Error(), library.ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	fmt.Print("SetData request params:", req)
	resp := services.SetData(req)

	res := dtos.SteDataRes{}
	res.Success = "false"
	if resp == true {
		res.Success = "true"
	}

	library.WriteJSONResponse(res, http.StatusOK, rd)
}
