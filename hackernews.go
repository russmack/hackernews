// Package hackernews is a HackerNews api client.
package hackernews

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	EndpointTopStories string = "https://hacker-news.firebaseio.com/v0/topstories.json"
	EndpointNewStories string = "https://hacker-news.firebaseio.com/v0/newstories.json"
	EndpointMaxItem    string = "https://hacker-news.firebaseio.com/v0/maxitem.json?print=pretty"
	EndpointItem       string = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
	EndpointAskStories string = "https://hacker-news.firebaseio.com/v0/askstories.json?print=pretty"
	EndpointUpdates    string = "https://hacker-news.firebaseio.com/v0/updates.json?print=pretty"
	EndpointUser       string = "https://hacker-news.firebaseio.com/v0/user/%s.json?print=pretty"
)

// Client is the api client.
type Client struct{}

// NewClient returns a new api Client client.
func NewClient() *Client {
	return &Client{}
}

// unmarshalItem retuns a HackerNews Item, given the json.
func (c *Client) unmarshalItem(jsonItem string) (*Item, error) {
	item := new(Item)
	err := json.Unmarshal([]byte(jsonItem), &item)
	if err != nil {
		log.Println("Error calling json.Unmarshal from unmarshalItem.")
		return nil, err
	}
	return item, nil
}

// GetTopStories returns the ids of the top stories.
func (c *Client) GetTopStories() ([]int, error) {
	jsonRes, err := c.getUrl(EndpointTopStories)
	if err != nil {
		log.Println("Err calling getUrl from GetTopStories.")
		return nil, err
	}
	stories := make([]int, 0)
	err = json.Unmarshal([]byte(jsonRes), &stories)
	if err != nil {
		log.Println("Err calling json.Unmarshal from GetTopStories.")
		return nil, err
	}
	return stories, nil
}

// GetItem returns the Item specified by the id.
func (c *Client) GetItem(id int) (*Item, error) {
	itemUrl := fmt.Sprintf(EndpointItem, id)
	jsonRes, err := c.getUrl(itemUrl)
	if err != nil {
		log.Println("Err calling getUrl form GetItem.")
		return nil, err
	}
	return c.unmarshalItem(jsonRes)
}

// getUrl is the http client which requests a given endpoint and returns the response.
func (c *Client) getUrl(url string) (string, error) {
	tr := &http.Transport{
		TLSHandshakeTimeout: 20 * time.Second,
	}
	client := &http.Client{Transport: tr}
	//client := http.Client{}
	log.Println("Getting:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Err in http.NewRequest.")
		return "", err
	}
	req.Header.Add("Accept-Encoding", "identity")
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Err in client.Do.")
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = errors.New("Http server returned status code: " + string(resp.StatusCode))
		return "", err
	}
	body := []byte{}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Err in ioutil.ReadAll.")
		return "", err
	}
	return string(body), nil
}
