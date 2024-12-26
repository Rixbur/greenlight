package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// It’s crucial to point out here that all the fields in our Movie struct are exported (i.e. start with a capital letter),
// which is necessary for them to be visible to Go’s encoding/json package.
// Any fields which aren’t exported won’t be included when encoding a struct to JSON
type Movies struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime,omitempty"` // if you want just to add omitempty ",omitempty", comma is still required
	Genres    []string  `json:"genresm,omitempty"`
	Version   int32     `json:"version"`
}

// - directive is used to never output the field when encoding to JSON.
// omitempty - directive is used to only output the field when it has a non-zero value.
// string - represented as a JSON string. Note that the string directive will only work on struct fields which have int*, uint*, float* or bool types

func (m Movies) MarshalJSON() ([]byte, error) {
	var runtime string
	if m.Runtime > 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	// You can also use initialize empty anonymous struct and set the values later.
	// aux := struct {
	// 	ID        int64    `json:"id"`
	// 	Title     string   `json:"title"`
	// 	Year      int32    `json:"year,omitempty"`
	// 	Runtime   string   `json:"runtime,omitempty"`
	// 	Genres    []string `json:"genres,omitempty"`
	// 	Version   int32    `json:"version"`
	// }{}

	aux := struct {
		ID      int64    `json:"id"`
		Title   string   `json:"title"`
		Year    int32    `json:"year,omitempty"`
		Runtime string   `json:"runtime,omitempty"`
		Genres  []string `json:"genres,omitempty"`
		Version int32    `json:"version"`
	}{
		// Set the values for the anonymous struct.
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime, // Note that we’re using the runtime variable here, rather than the m.Runtime field directly.
		Genres:  m.Genres,
		Version: m.Version,
	}

	return json.Marshal(aux)
}
