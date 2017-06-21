package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"
	"github.com/EthanG78/tweego/stream"
)

var(
	consumerKey = logger.GetEnv("TWITTER_CONSUMER_KEY")
	consumerSecret = logger.GetEnv("TWITTER_CONSUMER_SECRET")
	accessToken = logger.GetEnv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = logger.GetEnv("TWITTER_ACCESS_TOKEN_SECRET")
)


func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger.Newlogger{logrus.New()}
	api.SetLogger(log)

	s := stream.StreamFilter(api, "#trump")

	defer s.Stop()

	stream.StreamRetweet(s , api)


}
