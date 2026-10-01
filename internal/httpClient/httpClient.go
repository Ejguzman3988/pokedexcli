package httpClient

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	. "github.com/ejguzman3988/pokedexcli/internal/pokecache"
)

type mapData struct {
	Name string `json:"name"`
}

type LocUrls struct {
	Next string
	Prev string
}

type Client struct {
	Cache      Cache
	Locations  LocUrls
	httpClient http.Client
}

type resData struct {
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []mapData `json:"results"`
}

func (c *Client) GetLocations(pageUrl string) (resData, error) {

	body, exist := c.Cache.Get(pageUrl)

	if !exist {

		req, err := http.NewRequest("GET", pageUrl, nil)

		if err != nil {
			log.Fatal(err)
			return resData{}, err
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			log.Fatal(err)
			return resData{}, err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if err != nil {
			log.Fatal(err)
			return resData{}, err
		}

		c.Cache.Add(pageUrl, body)
	}

	locations := resData{}

	if err := json.Unmarshal(body, &locations); err != nil {
		return resData{}, err
	}

	c.Locations.Next = locations.Next
	c.Locations.Prev = locations.Previous

	return locations, nil
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: *NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		Locations: LocUrls{
			Next: "",
			Prev: "",
		},
	}
}
