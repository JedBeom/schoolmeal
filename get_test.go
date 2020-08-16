package schoolmeal

import (
	"testing"
	"time"
)

func TestSchool_GetWeekMeal(t *testing.T) {
	school, err := Find(Jeonnam, "광양제철고등학교")
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	meals, err := school.GetWeekMeal("2020.08.08", Lunch)
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	meal := meals[time.Saturday]

	menu := `쌀밥
미니냉면1.3.5.6.13.16.
모듬떡볶이1.2.5.6.10.12.13.15.16.18.
잡채김말이튀김1.5.6.16.18.
찐순대6.10.
숙주나물무침1.5.6.13.
배추김치9.13.`
	if meal.Content != menu {
		t.Error("Expected", menu, "Unexpected", meal.Content)
	}

	dateString := "2020.08.08(토)"
	if meal.DateString != dateString {
		t.Error("Expected", dateString, "Unexpected", meal.DateString)
	}

	people := 409
	if meal.People != people {
		t.Error("Expected", people, "Unexpected", meal.People)
	}

}

func TestSchool_GetDayMeal(t *testing.T) {
	school, err := Find(Jeonnam, "광양제철고등학교")
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	m, err := school.GetDayMeal("2020.08.20", Lunch)
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	dateString := "2020.08.20(목)"
	if m.DateString != dateString {
		t.Error("Expected", dateString, "Unexpected", m.DateString)
	}
}
