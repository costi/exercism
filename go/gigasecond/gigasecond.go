// Package gigasecond adds gigaseconds to time
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = 1000000000

// AddGigasecond adds a Gigasecond to a time.Time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond) * time.Second)
}
