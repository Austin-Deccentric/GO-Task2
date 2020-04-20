package linkedin


//LinkedIn defines the social media platform LinkedIn
type LinkedIn struct {
	Url string
	Name string	`xml:"name,omitempty"`
	Connections int
}


//Linked is an instance of LinkedIn struct
var Linked *LinkedIn = &LinkedIn{
	Url:  "linkedin/Austin-Deccentric",
	Name: "Chukwu Austin",
	Connections: 78,
}


//Data defines the social media plaform LinkedIn
type Data struct {
	//XMLNAME xml.Name `xml:"info"`
	User string		`xml:"Name,attr"`
	Url string		`xml:"Url"`
	Connections int	`xml:"Connections"`
	LinkedIn			
}

//LinkedData is  mock LinkedIn data 
var LinkedData *Data = &Data{
	Url: "linkedin.com/kal_drogo",
	User: "Emeka Ola",
	Connections: 3629,
}


// Feed returns the latest LinkedIn posts
func (l *LinkedIn) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, here's my cool new selfie",
		"What is code?",
	}
}

// Fame tells how famous a user is. It is proprtional to the number of friends
func (l *LinkedIn) Fame() int {
	return l.Connections
}

//Response implements a mock json data response
func (l *LinkedIn) Response() map[string]interface{} {
	response := make(map[string]interface{})
	response["Name"] = "Sean Micheals"
	response["Number of Friends"] = 5000
	response["Number of followers"] = 2715
	return response
}
