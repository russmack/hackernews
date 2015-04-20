package main

import (
	"fmt"
	"github.com/russmack/hackernews"
)

func main() {
	client := hackernews.NewClient()
	stories, err := client.GetTopStories()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got stories: %v\n", stories)

	topStoryId := stories[0]
	fmt.Println("Top story id:", topStoryId)

	story, err := client.GetItem(topStoryId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Top story: %v\n", story)
	fmt.Println("Title:", story.Title)
	fmt.Println("Url:", story.Url)
	fmt.Println("Type:", story.Type)
}
