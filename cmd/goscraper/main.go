package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/DFwJZ/goscraper/internal/episode"
)

func main() {
	numberOfEpisodesToScrape := 20
	files := make(map[string]string)
	errors := 0

	fmt.Printf("Starting to scrape %d episodes...\n", numberOfEpisodesToScrape)

	for num := 1; num <= numberOfEpisodesToScrape; num++ {
		fmt.Printf("\rTrying episode number: %d", num)
		ep := episode.New(num)

		mp3URL, err := ep.MP3URL()
		if err != nil {
			if !strings.Contains(err.Error(), "MP3 file not found") {
				log.Printf("\nError getting MP3 URL for episode %d: %v", num, err)
			}
			errors++
			continue
		}

		files[ep.EpisodeStr] = mp3URL
	}

	fmt.Printf("\n\nScraping completed. Found %d/%d episodes.\n", len(files), numberOfEpisodesToScrape)
	if errors > 0 {
		fmt.Printf("Encountered errors for %d episodes.\n", errors)
	}

	fmt.Println("\nAll found MP3 URLs:")

	// Sort the episode numbers for ordered output
	var sortedEpisodes []string
	for episodeStr := range files {
		sortedEpisodes = append(sortedEpisodes, episodeStr)
	}
	sort.Strings(sortedEpisodes)

	for _, episodeStr := range sortedEpisodes {
		fmt.Printf("%s: %s\n", episodeStr, files[episodeStr])
	}
}