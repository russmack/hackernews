// This is an example usage of the HackerNews api cliient library.
package main

import (
	"fmt"
	"github.com/russmack/hackernews"
)

func main() {
	// Create a client.
	client := hackernews.NewClient()
	// Get ids of top stories.
	stories, err := client.GetTopStories()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Got stories: %v\n", stories)
	// Get the top story.
	topStoryId := stories[0]
	fmt.Println("Top story id:", topStoryId)
	story, err := client.GetItem(topStoryId)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Display the top story.
	fmt.Printf("Top story: %v\n", story)
	fmt.Println("Title:", story.Title)
	fmt.Println("Url:", story.Url)
	fmt.Println("Type:", story.Type)
}
