package time

import (
	"strings"
	"time"
)

const layout = "2006-01-02T15:04:05"

type BitFlyerTime struct {
	time.Time
}

func (t *BitFlyerTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "Z\"")
	t.Time, err = time.Parse(layout, s)
	return err
}
