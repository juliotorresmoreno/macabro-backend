package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-redis/redis"
	"github.com/juliotorresmoreno/macabro/db"
	"github.com/juliotorresmoreno/macabro/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
}

// Credentials s
type Credentials struct {
	Email    string
	Password string
}

func (el authController) POSTLogin(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer conn.Close()
	p := &Credentials{}
	err = c.Bind(p)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	u := &models.User{}
	u.Email = p.Email
	_, err = conn.Get(u)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(p.Password),
	)

	token := bson.NewObjectId().Hex()
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisCli.Set(token, u.ID, 24*time.Hour)
	go redisCli.Close()

	c.SetCookie(&http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   60 * 60 * 10,
		HttpOnly: true,
	})
	if err == nil {
		return c.JSON(200, u)
	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"message": "password: password is not valid",
	})
}

// CredentialsRecovery s
type CredentialsRecovery struct {
	Email string
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (el authController) POSTRecovery(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer conn.Close()
	p := &CredentialsRecovery{}
	err = c.Bind(p)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	token := StringWithCharset(40, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz")
	if p.Email == "" {
		return &echo.HTTPError{
			Code:    http.StatusNotAcceptable,
			Message: "email is required.",
		}
	}
	u := &models.User{RecoveryToken: token}
	q := &models.User{Email: p.Email}

	_, err = conn.Omit("acl").Update(u, q)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return c.String(204, "")
}

// CredentialsReset s
type CredentialsReset struct {
	Password string
	Token    string
}

func (el authController) POSTReset(c echo.Context) error {
	conn, err := db.NewEngigne()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer conn.Close()
	p := &CredentialsReset{}
	err = c.Bind(p)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if p.Token == "-" || p.Token == "" {
		return &echo.HTTPError{
			Code:    http.StatusNotAcceptable,
			Message: "Token is required.",
		}
	}
	q := &models.User{RecoveryToken: p.Token}
	u := &models.User{}
	_, err = conn.Get(u)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		}
	}

	u.SetPassword(p.Password)
	u.RecoveryToken = "-"
	if err := u.Check(); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		}
	}
	_, err = conn.Omit("acl").Update(u, q)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return c.String(204, "")
}

// AuthController s
func AuthController(g *echo.Group) {
	c := authController{}
	g.POST("/login", c.POSTLogin)
	g.POST("/recovery", c.POSTRecovery)
	g.POST("/reset", c.POSTReset)
}
