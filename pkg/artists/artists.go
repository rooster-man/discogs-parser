package artists

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func GetArtists(fn string) []Artist {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully opened %s", fn)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var artists Artists
	xml.Unmarshal(byteValue, &artists)

	return artists.Artists
}

type Artists struct {
	XMLName xml.Name `xml:"artists"`
	Artists []Artist `xml:"artist"`
}

type Artist struct {
	XMLName        xml.Name `xml:"artist"`
	Name           string   `xml:"name"`
	NameVariations []string `xml:"name_variations"`
	DiscogsID      int      `xml:"id"`
	Profile        string   `xml:"profile"`
	Images         []Image  `xml:"images"`
}

type Image struct {
	Height      int    `xml:"height"`
	Width       int    `xml:"width"`
	ResourceURL string `xml:"resource_url"`
	Type        string `xml:"type"`
	URI         string `xml:"uri"`
	URI150      string `xml:"uri150"`
}
