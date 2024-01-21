package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	e := echo.New()
	e.GET("/read", func(c echo.Context) error {

		key := c.QueryParam("key")

		value := client.Get(c.Request().Context(), fmt.Sprintf("public:%s", key))

    sVal, err:= value.Result()
    if err != nil {
			return c.String(http.StatusBadRequest, "ERROR")
		}

		return c.String(http.StatusOK, sVal)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
