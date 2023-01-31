package main

import (
	"demo-ws/common/config"
	"demo-ws/common/log"
	"demo-ws/example-1/delivery"
	"demo-ws/example-1/global"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := config.GetConfig(nil, "./config.yml", &global.Cfg); err != nil {
		log.Fatalf("GetConfig error: %v", err)
	}
	log.Infof("config: %+v", global.Cfg)

	handler := delivery.NewHandler()

	e := echo.New()
	e.GET("/api/demo", handler.Handler1())
	log.Fatalf("%v", e.Start(global.Cfg.Address))
}
