package bot

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/imishinist/ayb/pkg/bot/twitter"
)

type Server struct {
	e *echo.Echo
}

func CreateServer() *Server {
	e := echo.New()
	s := &Server{
		e: e,
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/bot")
	// g.Use(echo.WrapMiddleware(fromCron))
	g.GET("/tweet", s.tweet)
	return s
}

func (s *Server) tweet(c echo.Context) error {
	client, err := twitter.GetClient(&twitter.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	})
	if err != nil {
		return err
	}
	client.Tweet("test")

	type response struct {
		Message string `json:"message"`
	}
	r := new(response)
	r.Message = "OK"
	return c.JSON(http.StatusOK, r)
}

func (s *Server) ListenAndServe(address string) error {
	return s.e.Start(address)
}
