package main

// import (
// 	"echoApp/handler"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	println("Start!🐮")
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		println("🌟")
// 		return c.String(http.StatusOK, "Hello world..aa")

// 	})

// 	e.GET("/users/:id", handler.GetUser)
// 	e.POST("/users", handler.SaveUser)
// 	// e.PUT("/users/:id", updateUser)
// 	// e.DELETE("/users/:id", deleteUser)

// 	e.Logger.Fatal(e.Start(":1323"))
// }
