package main

import (
	"sme/exporter"
	"sme/facebook"
	"sme/twitter"
)


func main() {
	fb := new(facebook.Facebook)
	fb = &facebook.Facebook{
		Name: "Segun Elohor",
		Friends: 567,
	}
	tweet := new(twitter.Twitter)
	// // linked := new(linkedin.LinkedIn)
	err := exporter.Exporttxt(fb,"fb_data.txt")
	err = exporter.Exporttxt(tweet, "tweet_data.txt")
	err = exporter.Exportjson(fb, "fbdata.json")
	if err != nil {
		panic(err)
	}
	err = exporter.Exportresponse(tweet, "twitter_response.json")
	if err != nil {
		panic((err))
	}

}

