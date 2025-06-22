# Timeless
Converting user inputted time lengths and dates into time objects.
## Examples
```go
package main

import "github.com/golittie/timeless"

func main() {
	timeless.SYSTEM_TIMEZONE = timeless.UTC
	timeless.DEFAULT_DATE_FORMAT = timeless.DDMMYY
	
	timeless.Parse("5m") // Time.now() + 5 minutes
	timeless.Parse("1h 5m") // Time.now() + 1 hour and 5 minutes
	timeless.Parse("14:12") // Time.time of today at 14:12
	timeless.Parse("04/06/2025") // Time.time of the start of the day 04/06/2025
	timeless.Parse("04/06/2025 14:12") // Time.time of 04/06/2025 at 14:12
	timeless.Parse("04/06/2025 14:12 5m") // Time.time of 04/06/2025 at 14:17
	timeless.Parse("04/06/2025 14:12 5m", timeless.WithTimezone(timeless.BST)) // Time.time of 04/06/2025 at 15:17
	timeless.Parse("06/04/2025 14:12 5m", timeless.WithDateFormat(timeless.MMDDYY)) // Time.time of 04/06/2025 at 14:17

	timeless.ParseTimeLength("10m") // time.Minute * 10
	timeless.ParseTimeLength("2h 30m") // time.Hour * 2 + time.Minute * 30
	timeless.ParseTimeLength("1d 7m") // timeless.Day + time.Minute * 7
	timeless.ParseTimeLength("-3w 1m") // timeless.Week * -3 + time.Minute
	timeless.ParseTimeLength("-3w 1m", timeless.WithoutNegatives()) // time.Week * 3 + time.Minute

	timeless.ParseRelativeTimeLength("10m") // Time.now() + 10 minutes
	timeless.ParseRelativeTimeLength("2h 30m") // Time.now() + 2 hours and 30 minutes
	timeless.ParseRelativeTimeLength("1d 7m") // Time.now() + 1 day and 7 minutes
	timeless.ParseRelativeTimeLength("3h 10m", timeless.WithTimezone(timeless.BST)) // Time.now() + 4 hours and 10 minutes
}
```
## Example Console Application

```go
package main

import (
	"bufio"
	"fmt"
	"github.com/golittie/timeless"
	"os"
	"strings"
)

func main() {
	timeless.SYSTEM_TIMEZONE = timeless.UTC
	timeless.DEFAULT_DATE_FORMAT = timeless.DDMMYY
	
	reader := bufio.NewReader(os.Stdin)

	for {
		s := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		
		t := timeless.Parse(s)
		fmt.Println(t.String())
	}
}
```