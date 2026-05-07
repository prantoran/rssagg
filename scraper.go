package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/prantoran/rssagg/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting scraper with concurrency=%d (goroutines) and timeBetweenRequest=%s\n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		log.Println("Scraper tick: fetching feeds...")
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("Error fetching feeds: %v\n", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching RSS feed: %v\n", err)
		return
	}
	for _, item := range rssFeed.Channel.Items {
		log.Println("Foound post:", item.Title, " on feed:", feed.Name)
	}
	log.Printf("Feed '%s' fetched successfully with %d items\n", feed.Name, len(rssFeed.Channel.Items))
}
