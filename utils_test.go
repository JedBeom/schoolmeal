package schoolmeal

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {

	tz, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		t.Error("Unexpected Error", err)
		t.Failed()
	}

	date := time.Date(2019, time.February, 3, 15, 32, 21, 0, tz)
	if stamp := Timestamp(date); stamp != "2019.02.03" {
		t.Error("Expected 2019.02.03 unexpected", stamp)
	}

	date2 := time.Date(2019, time.October, 31, 20, 30, 21, 0, tz)
	if stamp := Timestamp(date2); stamp != "2019.10.31" {
		t.Error("Expected 2019.10.31 unexpected", stamp)
	}

	date3 := time.Date(2019, time.December, 4, 1, 4, 1, 123, tz)
	if stamp := Timestamp(date3); stamp != "2019.12.04" {
		t.Error("Expected 2019.12.04 unexpected", stamp)
	}
}
