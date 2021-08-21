// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
// Get the moment after 10 ^ 9 seconds has past
func AddGigasecond(t time.Time) time.Time {
	elapsedSeconds := 1_000_000_000
	return time.Unix(t.Unix()+int64(elapsedSeconds), int64(t.Nanosecond()))
}
