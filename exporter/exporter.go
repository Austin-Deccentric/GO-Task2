package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"sme"
)

//Exporttxt writes a new .txt file to the current working directory
func Exporttxt(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	
	for _, feed := range platform.Feed() {
		n, err := f.WriteString(feed + "\n")
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes\n",n)
	}
	defer f.Close()
	return nil
}


//Exportjson implements writing a json data to file
func Exportjson(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.WriteString(string(jsondata) + "\n")
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	
	defer f.Close()
	return nil
}

//Exportresponse writes the response data to a json file.
func Exportresponse(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(platform.Response())
	if err != nil {
		return err
	}

	sep := "\n"
	n, err := f.WriteString(string(jsondata) + sep)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes\n", n)

	defer f.Close()
	return nil
}