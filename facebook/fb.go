package facebook


//Facebook is a mock type containing basic Info to fufill this learning stage
type Facebook struct {
	URL string
	Name string
	Friends int
}

// Feed returns the latest Facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

// Fame tee how famous a user is. It is proprtional to the number of friends
func (f *Facebook) Fame() int {
	return f.Friends
}

//Response is a mock server response
func (f *Facebook) Response() map[string]interface{} {
	response := make(map[string]interface{})
	response["Handle"] = "@kal_Drogo"
	response["Number of followers"] = 3289
	response["Number of following"] = 2715
	return response
}

