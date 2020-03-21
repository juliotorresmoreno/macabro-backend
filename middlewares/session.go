package middlewares

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
)

func Session(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, _ := c.Cookie("token")
		if token == nil {
			return handler(c)
		}
		redisCli := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		go redisCli.Close()
		r := redisCli.Get(token.Value).Val()
		if r != "" {
			conn, err := db.NewEngigne()
			if err == nil {
				id, _ := strconv.Atoi(r)
				u := &models.User{ID: uint(id)}
				if ok, _ := conn.Get(u); ok {
					u.Password = ""
					u.ValidPassword = ""
					c.Set("session", u)
				}
			}
			go conn.Close()
		}
		return handler(c)
	}
}
