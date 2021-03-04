package bot

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/imishinist/ayb/pkg/bot/twitter"
	"github.com/imishinist/ayb/pkg/witticism"
)

type Server struct {
	e    *echo.Echo
	conf *Config
}

func CreateServer() *Server {
	e := echo.New()
	s := &Server{
		e:    e,
		conf: loadConfig(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	api.GET("/tweets", s.tweetList)

	g := e.Group("/bot")
	if s.conf.Env == "prod" {
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
		AccessToken:       s.conf.AccessToken,
		AccessTokenSecret: s.conf.AccessTokenSecret,
		ConsumerKey:       s.conf.ConsumerKey,
		ConsumerSecret:    s.conf.ConsumerSecret,
	})
	if err != nil {
		return err
	}
	w := witticism.Get()
	tweetText := w.Random().Text
	if err := client.Tweet(tweetText); err != nil {
		return err
	}
	log.Println(tweetText)

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
