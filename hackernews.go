package hackernews

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"private/russmack/hackernews/types"
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

type Client struct {
}

var ()

func init() {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) unmarshalItem(jsonItem string) (*types.Item, error) {
	item := new(types.Item)
	err := json.Unmarshal([]byte(jsonItem), &item)
	if err != nil {
		fmt.Println("Error unmarshalling item.")
		return nil, err
	}
	return item, nil
}

func (c *Client) GetTopStories() ([]int, error) {
	jsonRes, err := c.getUrl(EndpointTopStories)
	if err != nil {
		fmt.Println("Err:", err)
		return nil, err
	}
	stories := make([]int, 0)
	err = json.Unmarshal([]byte(jsonRes), &stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (c *Client) GetItem(id int) (*types.Item, error) {
	itemUrl := fmt.Sprintf(EndpointItem, id)
	jsonRes, err := c.getUrl(itemUrl)
	if err != nil {
		fmt.Println("Error getting item.")
		return nil, err
	}
	return c.unmarshalItem(jsonRes)
}

func (c *Client) getUrl(url string) (string, error) {
	fmt.Println("Getting url:", url)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
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
		return "", err
	}
	return string(body), nil
}
