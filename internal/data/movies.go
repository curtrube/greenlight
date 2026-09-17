package data

import "time"

type Movie struct {
	ID        int       `json:"id"`               // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`            // Movie title
	Year      int       `json:"year,omitzero"`    // Move release year
	Runtime   int       `json:"runtime,omitzero"` // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitzero"`  // Slice of genres for the movie (romance, comedy, etc)
	Version   int       `json:"version"`          // The version number start at 1 and will be incremented each time
	//the movie information is updated.
}
