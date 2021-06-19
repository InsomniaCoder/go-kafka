package main

import (
	"github.com/insomniacoder/go-kafka/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
