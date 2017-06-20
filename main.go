package main

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"

	"net/url"
)



func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	/*res,_ := api.GetSearch("golang", nil)
	for _, tweet := range res.Statuses{
		fmt.Print(tweet.Text)
	}*/

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
