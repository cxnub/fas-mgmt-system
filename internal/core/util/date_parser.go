package util

import (
	"fmt"
	"time"
)

func ParseDate(date string) (*time.Time, error) {
	if parsedDate, err := time.Parse("2006-01-02", date); err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return nil, err
	} else {
		return &parsedDate, nil
	}
}
