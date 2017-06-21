package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"
	"github.com/EthanG78/tweego/stream"

	"net/url"
	"os"
)

var(
	consumerKey = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = getenv("TWITTER_CONSUMER_SECRET")
	accessToken = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv (name string) string{
	env := os.Getenv(name)
	if env == ""{
		panic("Missing environment variable " + name)
	}
	return env
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger.Newlogger{logrus.New()}
	api.SetLogger(log)

	s := api.PublicStreamFilter(url.Values{
		"track":[]string{"#trump"}, //PLACEHOLDER
	})

	defer s.Stop()

	stream.StreamRetweet(s , api)


}
