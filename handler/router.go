package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() http.Handler {
	router := httprouter.New()
	router.GET("/ping", ping)
	return router
}
