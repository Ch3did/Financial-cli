package utils

import (
	"strings"
	"time"
)

func ParseOFXDate(s string) time.Time {
	if idx := strings.Index(s, "["); idx != -1 {
		s = s[:idx]
	}

	layout := "20060102150405"
	tempo, err := time.Parse(layout, s)
	if err != nil {
		if len(s) >= 8 {
			if t2, err2 := time.Parse("20060102", s[:8]); err2 == nil {
				return t2
			}
		}
		return time.Time{}
	}
	return tempo
}
