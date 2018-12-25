package library

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	ERR_MSG = "ERROR_MESSAGE"
	MSG     = "MESSAGE"
)

type RequestData struct {
	Start time.Time
	w     http.ResponseWriter
	r     *http.Request
}

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case ERR_MSG:
		Obj.Status = "FAILED"

	case MSG:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONResponse(d []byte, code int, rd *RequestData) {
	fmt.Print(rd.r, time.Since(rd.Start).Seconds(), code)
	if code == http.StatusInternalServerError {
		fmt.Print("Status Code:", code, ", Response time:", time.Since(rd.Start), " Response:", string(d))
	} else {
		fmt.Print("Status Code:", code, ", Response time:", time.Since(rd.Start))
	}

	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
}

func writeJSONMessage(msg string, msgType string, httpCode int, rd *RequestData) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponse(d, code, rd)
}

func WriteJSONResponse(i interface{}, code int, rd *RequestData) {
	d, err := json.Marshal(i)
	if err != nil {
		writeJSONMessage("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONResponse(d, code, rd)
}

func LogAndGetContext(w http.ResponseWriter, r *http.Request) *RequestData {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")
	start := time.Now()
	return &RequestData{
		Start: start,
		r:     r,
		w:     w,
	}
}
