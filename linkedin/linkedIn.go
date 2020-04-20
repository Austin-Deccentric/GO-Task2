package linkedin


//LinkedIn defines the social media platform LinkedIn
type LinkedIn struct {
	Url string
	Name string
	Connections int
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
