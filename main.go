package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"
	"github.com/EthanG78/tweego/stream"
	"github.com/EthanG78/tweego/actions"
	env "github.com/EthanG78/tweego/env_variables"
)

var(
	consumerKey = env.GetEnv("TWITTER_CONSUMER_KEY")
	consumerSecret = env.GetEnv("TWITTER_CONSUMER_SECRET")
	accessToken = env.GetEnv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = env.GetEnv("TWITTER_ACCESS_TOKEN_SECRET")

	//TEST

)


func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	var new_log = logrus.New()
	log := &logger.Newlogger{Logger: new_log}
	api.SetLogger(log)

	s := stream.StreamFilter(api, "#trump")

	defer s.Stop()

	actions.StreamRetweet(s , api)


}
