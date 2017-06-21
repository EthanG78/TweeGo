package stream

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/EthanG78/tweego/logger"
	"github.com/sirupsen/logrus"
)


func StreamRetweet (stream *anaconda.Stream, api *anaconda.TwitterApi){
	log := &logger.Newlogger{logrus.New()}

	for v := range stream.C{
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

//TODO: Create test file for stream^^