package main

import (
	"database/sql"
	"demo-ws/common/config"
	"demo-ws/common/log"
	"demo-ws/example-1/delivery"
	"demo-ws/example-1/global"
	"demo-ws/example-1/repo"
	"demo-ws/example-1/service"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := config.GetConfig(nil, "./config.yml", &global.Cfg); err != nil {
		log.Fatalf("GetConfig error: %v", err)
	}
	log.Infof("config: %+v", global.Cfg)

	// repo
	cfg := mysql.Config{
		User:                 global.Cfg.MySql.Username,
		Passwd:               global.Cfg.MySql.Passsword,
		Net:                  "tcp",
		Addr:                 global.Cfg.MySql.Address,
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Open error: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping error: %v", err)
	}
	repo1 := repo.NewRepo1(db)

	// service
	service1 := service.NewService1(repo1)

	// echo
	handler := delivery.NewHandler1(service1)

	e := echo.New()
	e.GET("/api/albums/count", handler.Handler1())
	log.Fatalf("%v", e.Start(global.Cfg.Address))

}
