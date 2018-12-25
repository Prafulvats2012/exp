package main

import (
	"exp/config"
	"exp/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Application started at port:", config.PORT_NUMBER)
	http.ListenAndServe(":"+config.PORT_NUMBER, handler.GetRouter())
}
