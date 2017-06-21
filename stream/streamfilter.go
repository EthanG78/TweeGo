package stream

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

func StreamFilter (api *anaconda.TwitterApi, track string) (stream *anaconda.Stream){
	return api.PublicStreamFilter(url.Values{
		"track": []string{track},
	})
}

//TODO: Create test file for stream^^