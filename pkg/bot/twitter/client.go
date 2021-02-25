package twitter

import (
	twi "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type Client struct {
	inner *twi.Client
}

func GetClient(creds *Credentials) (*Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twi.NewClient(httpClient)

	verifyParams := &twi.AccountVerifyParams{
		SkipStatus:   twi.Bool(true),
		IncludeEmail: twi.Bool(true),
	}
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return &Client{inner: client}, nil
}
