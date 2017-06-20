package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"github.com/EthanG78/tweego/logger"

	"net/url"
	"os"
)

var(
	consumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

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
			logrus.Warningf("Received unexpected value of %T", v)
			continue
		}

		if tweet.RetweetedStatus != nil {
			continue
		}
		_, err := api.Retweet(tweet.Id, false)

		if err != nil{
			logrus.Errorf("Could not retweet %d: %v",tweet.Id, err)
			continue
		}

		logrus.Infof("Retweeted %d", tweet.Id)
	}

}
