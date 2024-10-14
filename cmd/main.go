package main

import (
	"mods/internal/interfaces/config"
	"mods/internal/interfaces/http"
)

func main() {

	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	http.StartServer(c)

	alertService := NewAlertService(DicscordChat{}, chatBandwidth{})
	alertService.Start()

}
