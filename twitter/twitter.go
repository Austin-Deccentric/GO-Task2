package twitter


//Twitter defines the social media platform Twitter
type Twitter struct {
	Url string
	Name string
	Followers int
}

//Tweet is an instance of the Twitter struct
var Tweet  *Twitter = &Twitter{
	Url: "twitter.com/kal_drogo",
	Name: "Emeka Olusegun",
	Followers: 3629,
}

// Feed returns the latest Twitter posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Hey, here's my cool new selfie",
		"What! is you a bitch?",
	}
}

// Fame tells how famous a user is. It is proprtional to the number of friends
func (t *Twitter) Fame() int {
	return t.Followers
}


//Response is a mock server response
func (t *Twitter) Response() map[string]interface{} {
	response := make(map[string]interface{})
	response["Handle"] = "@kal_Drogo"
	response["Number of followers"] = 3289
	response["Number of following"] = 2715
	
	return response
}