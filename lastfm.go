package main

import (
	"net/http"
	"encoding/json"
	"math/rand"
)

type TagSearchResults struct {
	TopArtists TopArtists `json:"topartists"`
}

type TopArtists struct {
	Artists []Artist `json:"artist"`
}

func (topArtists *TopArtists) getRandomArtist() Artist {
	artistsCount := len(topArtists.Artists)
	randomArtistIndex := rand.Intn(artistsCount)

	return topArtists.Artists[randomArtistIndex]
}

type Artist struct {
	Name   string `json:"name"`
	URL    string  `json:"url"`
	Images []Image `json:"image"`
}

func (artist *Artist) GetPreview() string {
	for _, image := range artist.Images {
		if image.Size == "large" {
			return image.URL
		}
	}
	return ""
}

type Image struct {
	URL  string `json:"#text"`
	Size string `json:"size"`
}

func GetData(tag string) Artist {
	apiKey := "e7726640ce219c4dd79c609ea0c4ab73"
	searchResults := TagSearchResults{}
	url := "http://ws.audioscrobbler.com//2.0/?method=tag.gettopartists&tag=" + tag + "&api_key=" + apiKey + "&format=json"
	GetJson(url, &searchResults)
	artist := searchResults.TopArtists.getRandomArtist()

	return artist
}

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}