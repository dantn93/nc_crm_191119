package main

import (
	"fmt"
	"log"

	"github.com/golang191119/nc_crm/config"
	mw "github.com/golang191119/nc_crm/middleware"
	"github.com/golang191119/nc_crm/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.SimpleLogger())
	route.All(e)

	log.Println(e.Start(":9094"))
}
