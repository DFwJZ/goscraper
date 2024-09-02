package episode 

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
	"net/http"
	"github/dfwjz/goscraper/internal/logging"
)

const baseURL = "https://twit.cachefly.net/video/sn/"

var logger *slog.Logger

func init() {
	handler := logging.NewColorHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger = slog.New(handler)
}

type Episode struct {
	Number int
	FolderURL string
}


func New(number int) *Episode {
	episodeStr := fmt.Sprintf("sn%04d", number)
	FolderURL := baseURL + episodeStr + "/"
	return &Episode{
		Number: number,
		FolderURL: FolderURL,
	}
}

func (e *Episode) findFile(extension string) (string, error) {
	resp, err = http.Get(e.FolderURL)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var files []string
	lines:= strings.Split(string(body), "\n")
	for _, line := range lines{
		if strings.
	}


	return fmt.Sprintf("%s/%s.%s", e.FolderURL, e.Number, extension)
}


func (e *Episode) MP4URL() string {
	return fmt.
}