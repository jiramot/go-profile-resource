package main

import (
	auth "github.com/jiramot/go-profile-resource/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(auth.Auth)
	e.GET("/", func(c echo.Context) error {
		principle := c.Get("UserPrinciple").(*auth.UserPrinciple)
		profile := Profile{
			Firstname: "Foo",
			Lastname:  "Bar",
			CIF:       principle.CIF,
		}
		return c.JSON(http.StatusOK, profile)
	})
	e.Logger.Fatal(e.Start(":10000"))
}

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	CIF       string `json:"cif"`
}
