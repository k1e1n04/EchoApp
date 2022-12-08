package routes

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func Route() {
    e := echo.New()

    e.GET("/", getHandler)

    e.Logger.Fatal(e.Start(":1323"))
}

// サイトを表示させる
func getHandler(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}