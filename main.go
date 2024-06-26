package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

var (
	query      = flag.String("query", "pollberry", "Search term")
	maxResults = flag.Int64("max-results", 5, "Max YouTube results")
)

const developerKey = "API-KEY"

type videoType struct {
	VideoId string
}

func main() {
	flag.Parse()

	//Create Service(서비스 생성)
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id", "snippet"}).
		Q(*query).
		MaxResults(*maxResults).
		Type("video")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error get list: %v", err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	//channels := make(map[string]string)
	//playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
			//case "youtube#channel":
			//	channels[item.Id.ChannelId] = item.Snippet.Title
			//case "youtube#playlist":
			//	playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	//printIDs("Channels", channels)
	//printIDs("Playlists", playlists)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
