package bot

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/imishinist/ayb/pkg/bot/twitter"
	"github.com/imishinist/ayb/pkg/witticism"
)

type Server struct {
	e *echo.Echo
}

func CreateServer() *Server {
	conf := loadConfig()

	e := echo.New()
	s := &Server{
		e: e,
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	api.GET("/tweets", s.tweetList)

	g := e.Group("/bot")
	if conf.Env == "prod" {
		g.Use(echo.WrapMiddleware(fromCron))
	}
	g.GET("/tweet", s.tweet)
	return s
}

func (s *Server) tweetList(c echo.Context) error {
	ws := witticism.Get()
	type response struct {
		Tweets witticism.Witticisms `json:"tweets"`
	}

	return c.JSON(http.StatusOK, &response{
		Tweets: ws,
	})
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
	w := witticism.Get()
	if err := client.Tweet(w.Random().Text); err != nil {
		return err
	}

	type response struct {
		Message string `json:"message"`
	}
	r := &response{
		Message: "OK",
	}
	return c.JSON(http.StatusOK, r)
}

func (s *Server) ListenAndServe(address string) error {
	return s.e.Start(address)
}
