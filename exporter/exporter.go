package exporter
import (
	"sme"
	"os"
	"fmt"
	"encoding/json"
)

//Exporttxt writes a new .txt file to the current working directory
func Exporttxt(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	
	for _, feed := range platform.Feed() {
		n, err := f.Write([]byte(feed + "\n"))
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes\n",n)
	}
	
	return nil
}


//Exportjson implements writing a json data to file
func Exportjson(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(platform)
	if err != nil {
		return err
	}

	n, err := f.Write([]byte(jsondata))
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}

//Exportresponse writes the response data to a json file.
func Exportresponse(platform sme.SocialMedia, filename string) error{
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	jsondata, err := json.Marshal(platform.Response())
	if err != nil {
		return err
	}

	n, err := f.Write([]byte(jsondata))
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes\n", n)
	return nil
}