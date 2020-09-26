package schoolmeal

import (
	"testing"
	"time"
)

func TestSchool_GetMonthSchedule(t *testing.T) {
	s, err := Find(Jeonnam, "광양제철고등학교")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	sches, err := s.GetMonthSchedule(2020, 9)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	sche := sches[12]

	name := "추석연휴"
	if sche.Name != name {
		t.Error("Expected", name, "Unexpected", sche.Name)
	}

	seoul, _ := time.LoadLocation("Asia/Seoul")
	date := time.Date(2020, time.Month(9), 30, 0, 0, 0, 0, seoul)
	if !sche.Date.Equal(date) {
		t.Error("Expected", date, "Unexpected", sche.Date)
	}

	ty := 1 // 공휴일
	if sche.Type != ty {
		t.Error("Expected", ty, "Unexpected", sche.Type)
	}
}
