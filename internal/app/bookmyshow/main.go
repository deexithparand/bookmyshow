package main

import (
	"fmt"

	"github.com/deexithparand/bookmyshow/internal/app/bookmyshow/api"
)

func main() {
	fmt.Println("Start Bookmyshow")

	var api api.API

	api.Run()
}
