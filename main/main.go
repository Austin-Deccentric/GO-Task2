package main

import (
	"sme/exporter"
	"sme/facebook"
	"sme/linkedin"
	"sme/twitter"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	err := exporter.Exporttxt(facebook.Fb,"fb_data.txt")
	check(err)

	err = exporter.Exporttxt(twitter.Tweet, "tweet_data.txt")
	check(err)

	err = exporter.Exportjson(facebook.Fb, "fbdata.json")
	check(err)

	err = exporter.Exportjson(twitter.Tweet,"twitter_data.json")
	check(err)

	err = exporter.Exportresponse(twitter.Tweet, "twitter_response.json")
	check(err)

	err = exporter.Exporttxt(linkedin.Linked, "LinkedIn_data.txt")
	check(err)

	err = exporter.Exportjson(linkedin.Linked,"LinkedIn_data.json")
	check(err)

}

