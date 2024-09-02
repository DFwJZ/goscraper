package logging

import (
    "context"
    "encoding/json"
    "io"
    "log/slog"
)

type ColorHandler struct {
    out io.Writer
    opts *slog.HandlerOptions
}

func NewColorHandler(out io.Writer, opts *slog.HandlerOptions) *ColorHandler {
    if opts == nil {
        opts = &slog.HandlerOptions{}
    }
    return &ColorHandler{out: out, opts: opts}
}

func (h *ColorHandler) Enabled(ctx context.Context, level slog.Level) bool {
    return level >= h.opts.Level.Level()
}

func (h *ColorHandler) Handle(ctx context.Context, r slog.Record) error {
    level := r.Level.String()
    message := r.Message

    // Add color based on log level
    var colorCode string
    switch r.Level {
    case slog.LevelDebug:
        colorCode = "\033[36m" // Cyan
    case slog.LevelInfo:
        colorCode = "\033[32m" // Green
    case slog.LevelWarn:
        colorCode = "\033[33m" // Yellow
    case slog.LevelError:
        colorCode = "\033[31m" // Red
    default:
        colorCode = "\033[0m" // Default
    }

    // Create a map for JSON output
    logMap := map[string]interface{}{
        "level":   level,
        "message": message,
        "time":    r.Time.Format("2006-01-02 15:04:05"),
    }

    // Add attributes
    r.Attrs(func(a slog.Attr) bool {
        logMap[a.Key] = a.Value.Any()
        return true
    })

    // Convert to JSON
    jsonData, err := json.Marshal(logMap)
    if err != nil {
        return err
    }

    // Write colored output
    _, err = h.out.Write([]byte(colorCode))
    if err != nil {
        return err
    }
    _, err = h.out.Write(jsonData)
    if err != nil {
        return err
    }
    _, err = h.out.Write([]byte("\033[0m\n")) // Reset color and add newline
    return err
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    return h
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
    return h
}