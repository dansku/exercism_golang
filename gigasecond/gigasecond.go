// Package gigasecond provides all the tools for the gigasecond exercism task.
package gigasecond

import "time"

// AddGigasecond calculates when someone has lived for one gigasecond 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9)
	return t
}
