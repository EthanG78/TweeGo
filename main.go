package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"
	"github.com/EthanG78/tweego/stream"
	"github.com/EthanG78/tweego/actions"
	env "github.com/EthanG78/tweego/env_variables"
	"net/url"
)

var(
	//This will be the final way to fetch tokens
	/*consumerKey = env.GetEnv("TWITTER_CONSUMER_KEY")
	consumerSecret = env.GetEnv("TWITTER_CONSUMER_SECRET")
	accessToken = env.GetEnv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = env.GetEnv("TWITTER_ACCESS_TOKEN_SECRET")*/


	//Using this method of fetching tokens for quick testing without the need of creating new env variables
	consumerKey = env.ConsumerKey
	consumerSecret = env.ConsumerSecret
	accessToken = env.AccessToken
	accessTokenSecret = env.AccessTokenSecret

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


	//Testing different actions ->>

	//actions.StreamRetweet(s , api)

	v := url.Values{}
	v.Set("url", "https://api.twitter.com/1.1/statuses/home_timeline.json")
	actions.GetTimeline(api, v)


}
