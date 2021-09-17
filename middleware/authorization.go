package middleware

import (
    echo "github.com/labstack/echo/v4"
    "net/http"
    "strings"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cif := c.Request().Header.Get("X-CIF")
        scope := c.Request().Header.Get("X-API-SCOPE")
        s := strings.Split(scope, " ")
        if cif == "" {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }
        if !contains(s, "profile") {
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }
        c.Set("UserPrinciple", &UserPrinciple{
            CIF:   cif,
            Scope: scope,
        })
        return next(c)
    }
}
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }

    return false
}

type UserPrinciple struct {
    CIF   string
    Scope string
}
