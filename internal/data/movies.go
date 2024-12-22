package data

import "time"

// It’s crucial to point out here that all the fields in our Movie struct are exported (i.e. start with a capital letter),
// which is necessary for them to be visible to Go’s encoding/json package.
// Any fields which aren’t exported won’t be included when encoding a struct to JSON
type Movies struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}
