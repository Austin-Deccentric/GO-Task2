package sme

//SocialMedia defines methods to implement to be a social media
type SocialMedia interface {
	Feed() []string
	Fame() int
	Response() map[string]interface{}
}


