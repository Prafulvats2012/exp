package handler

import (
	"exp/dtos"
	"exp/library"
	"fmt"
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
