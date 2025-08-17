package main

import (
	"backend/src/app"
)

func main() {
	router, cfg := app.SetupRouter()

	httpPort := ":" + cfg.HTTP.Port
	router.Run(httpPort)
}
