package facebook


//Facebook is a mock type containing basic Info to fufill this learning stage
type Facebook struct {
	URL string
	Name string
	Friends int
}

//Fb is an instance of the Facebook struct
var Fb  *Facebook  = &Facebook{
	Name: "Segun Elohor",
	Friends: 567,
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
	response["Name"] = "Sean Micheals"
	response["Number of Friends"] = 5000
	response["Number of followers"] = 2715
	return response
}

