package web

import (
	l "log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

func Clean(in string) string {
	out := strings.TrimSpace(in)
	return out
}

func CleanInt(in string) int {
	out, _ := strconv.Atoi(in)
	return out
}

func CleanDate(in string) time.Time {

	loc, err := time.LoadLocation("GMT")
	time.Local = loc

	out, err := dateparse.ParseLocal(in)
	if err != nil {
		l.With("error", err, "date", in).Error("Error parsing date")
	}

	return out
}

func CleanDateWithDefault(in string, defaultTime time.Time) time.Time {

	loc, err := time.LoadLocation("GMT")
	time.Local = loc

	out, err := dateparse.ParseLocal(in)
	if err != nil {

		// l.With("error", err, "date", in).Error("Error parsing date")
		// metrics.LogError("web", "CleanDate", err)
		return defaultTime
	}

	return out
}
