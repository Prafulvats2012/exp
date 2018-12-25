package config

import "os"

var (
	PORT_NUMBER string
)

func init() {
	PORT_NUMBER = os.Getenv("PORT_NUMBER")
}
