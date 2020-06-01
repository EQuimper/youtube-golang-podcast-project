package main

import (
	"log"

	"ytpodcast/feeds"
	"ytpodcast/itunes"
)

func main() {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search("Full Stack Radio")
	if err != nil {
		log.Fatalf("error while searching: %v", err)
	}

	for _, item := range res.Results {
		log.Println("-------------------")
		log.Printf("Artist: %s", item.ArtistName)
		log.Printf("Podcast Name: %s", item.TrackName)
		log.Printf("Feed url: %s", item.FeedURL)

		feed, err := feeds.GetFeed(item.FeedURL)
		if err != nil {
			log.Fatalf("error while get feed: %v", err)
		}

		for _, pod := range feed.Channel.Item {
			log.Println("--------------------")
			log.Printf("Title: %s", pod.Title)
			log.Printf("Duration: %s", pod.Duration)
			log.Printf("Description: %s", pod.Description)
			log.Printf("URL: %s", pod.Enclosure.URL)
			log.Println("--------------------")
		}

		log.Println("-------------------")
	}
}
