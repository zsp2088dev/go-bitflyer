package time

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func check(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got : %v\n", got)
		t.Errorf("want: %v\n", want)
	}
}

func TestToTime(t *testing.T) {
	var err error
	bft := BitFlyerTime{}
	t1 := []byte("\"2006-01-02T15:04:05\"")
	t2 := []byte("\"2006-01-02T15:04:05.678\"")
	t3 := []byte("\"2006-01-02T15:04:05.678Z\"")

	err = bft.UnmarshalJSON(t1)
	if err != nil {
		log.Fatal(err.Error())
	}
	check(t, bft.Time, time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC))

	err = bft.UnmarshalJSON(t2)
	if err != nil {
		log.Fatal(err.Error())
	}
	check(t, bft.Time, time.Date(2006, 1, 2, 15, 4, 5, 678000000, time.UTC))

	err = bft.UnmarshalJSON(t3)
	if err != nil {
		log.Fatal(err.Error())
	}
	check(t, bft.Time, time.Date(2006, 1, 2, 15, 4, 5, 678000000, time.UTC))

}
