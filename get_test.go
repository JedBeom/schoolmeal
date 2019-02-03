package schoolmeal

import (
	"testing"
	"time"
)

func TestSchool_GetWeekMeal(t *testing.T) {
	school, err := Find(Seoul, "서울대학교사범대학부설고등학교")
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	meals, err := school.GetWeekMeal("2018.11.30", Lunch)
	if err != nil {
		t.Error("Unexpected", err)
		t.Failed()
	}

	meal := meals[time.Friday]
	menu := `쌀밥
바지락된장찌개5.6.13.18.
소버섯불고기5.6.13.16.
치커리유자무침5.6.13.
깍두기9.13.
피칸파이1.2.5.6.13.
김치전5.6.9.13.`
	if meal.Content != menu {
		t.Error("Expected", menu, "Unexpected", meal.Content)
	}

	date := "2018.11.30(금)"
	if meal.Date != date {
		t.Error("Expected", date, "Unexpected", meal.Date)
	}
}
