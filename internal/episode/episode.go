package episode

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/DFwJZ/goscraper/internal/logging"
)

var logger *slog.Logger

func init() {
	handler := logging.NewColorHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger = slog.New(handler)
}

type Episode struct {
	Number     int
	EpisodeStr string
	client     *http.Client
	BaseURL    string
}

func New(number int) *Episode {
	episodeStr := fmt.Sprintf("sn%04d", number)
	return &Episode{
		Number:     number,
		EpisodeStr: episodeStr,
		BaseURL:    "https://twit.cachefly.net/",
		client:     &http.Client{Timeout: 10 * time.Second},
	}
}

func (e *Episode) checkFileExists(ctx context.Context, url string) bool {
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "Error creating request",
			slog.String("url", url),
			slog.String("error", err.Error()),
		)
		return false
	}

	resp, err := e.client.Do(req)
	if err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "Error checking file",
			slog.String("url", url),
			slog.String("error", err.Error()),
		)
		return false
	}
	defer resp.Body.Close()

	exists := resp.StatusCode == http.StatusOK
	logger.LogAttrs(ctx, slog.LevelDebug, "File existence check",
		slog.String("url", url),
		slog.Bool("exists", exists),
		slog.Int("statusCode", resp.StatusCode),
	)
	return exists
}

func (e *Episode) findMP3File() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger.LogAttrs(ctx, slog.LevelInfo, "Starting MP3 file search...",
		slog.Int("episodeNumber", e.Number),
	)

	baseURL := fmt.Sprintf("%saudio/sn/%s/", e.BaseURL, e.EpisodeStr)
	url := fmt.Sprintf("%s%s.mp3", baseURL, e.EpisodeStr)

	logger.LogAttrs(ctx, slog.LevelDebug, "Checking MP3 file",
		slog.String("url", url),
	)

	if e.checkFileExists(ctx, url) {
		logger.LogAttrs(ctx, slog.LevelInfo, "MP3 file found",
			slog.String("url", url),
		)
		return url, nil
	}

	logger.LogAttrs(ctx, slog.LevelWarn, "MP3 file not found",
		slog.Int("episodeNumber", e.Number),
	)
	return "", fmt.Errorf("MP3 file not found for episode %d", e.Number)
}

func (e *Episode) MP3URL() (string, error) {
	return e.findMP3File()
}