package facebook


//Facebook is a mock type containing basic Info to fufill this learning stage
type Facebook struct {
	Url string
	Name string	`xml:"Name,omitempty"`
	Friends int
}

//Fb is an instance of the Facebook struct
var Fb  *Facebook  = &Facebook{
	Url: "fb/Abgero",
	Name: "Segun Elohor", 
	Friends: 567,
}

//Data defines the social media plaform Facebook
type Data struct {
	//XMLNAME xml.Name `xml:"info"`
	User string		`xml:"Name,attr"`
	Url string		`xml:"Url"`
	Friends int	`xml:"Friends"`
	Facebook			
}

//FbData is  mock Facebook data 
var FbData *Data = &Data{
	Url: "twitter.com/kal_drogo",
	User: "Emeka Ola",
	Friends: 3629,
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

