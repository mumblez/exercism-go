// Package gigasecond adds a gigasecond to the date
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// comvert 10^9 seconds into usable constant
const gigasec = time.Second * 1000000000

// AddGigasecond adds 1,000,000,000 (10^9) seconds to the time
func AddGigasecond(tt time.Time) time.Time {
	return tt.Add(gigasec)
}
