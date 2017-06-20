package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"

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
		panic("Missing environment variable " + env)
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
		"track":[]string{"#trump"},
	})

	defer s.Stop()

	for v := range s.C{
		tweet, ok := v.(anaconda.Tweet)
		if !ok{
			log.Warningf("Received unexpected value of %T", v)
			continue
		}

		if tweet.RetweetedStatus != nil {
			continue
		}
		_, err := api.Retweet(tweet.Id, false)

		if err != nil{
			log.Errorf("Could not retweet %d: %v",tweet.Id, err)
			continue
		}

		log.Infof("Retweeted %d", tweet.Id)
	}

}
