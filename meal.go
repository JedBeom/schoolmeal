package schoolmeal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/buger/jsonparser"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

// GetDayMeal 함수는 하루의 급식을 가져옵니다.
func (s School) GetDayMeal(date string, mealType int) (m Meal, err error) {
	weekMeals, err := s.GetWeekMeal(date, mealType)
	if err != nil {
		return
	}

	parse := parseTime(date)
	m = weekMeals[parse.Weekday()]
	return
}

// GetMonthMeal 함수는 인자로 받은 연도와 달의 전체 급식을 리턴합니다.
// Meal[날짜-1][타입-1]
func (s School) GetMonthMeal(year, month int) (monthMeals [][]Meal, err error) {
	reqFormat := `{"schulCode": "%s", "schulCrseScCode": %d, "schYm": "%d%02d"}`
	reqJSON := []byte(fmt.Sprintf(reqFormat, s.Code, s.Kind, year, month))

	doc, err := post(s, makeURL(s.Zone, linkMealMonthly), reqJSON)
	if err != nil {
		return
	}

	// get weekDietList json
	docDiet, _, _, err := jsonparser.Get(doc, "resultSVO", "mthDietList")
	if err != nil {
		return
	}

	rds := make([]resultDiet, 0, 6)
	if err = json.Unmarshal(docDiet, &rds); err != nil {
		return
	}

	if len(rds) == 0 {
		err = errors.New("unexpected: month meal is none")
		return
	}

	monthMeals = make([][]Meal, 0, 32)

	for _, rd := range rds {
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Sun))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Mon))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Tue))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Wed))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.The))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Fri))
		appendIfNotNil(&monthMeals, rdToMealMonth(year, month, rd.Sat))
	}

	return
}

// GetWeekMeal 함수는 인자로 받는 날짜가 포함된 주의 급식이 담긴 []Meal{}을 리턴합니다.
func (s School) GetWeekMeal(date string, mealType int) (meals []Meal, err error) {
	reqFormat := `{"schulCode": "%s", "schulCrseScCode": %d, "schMmealScCode": %d, "schYmd": "%s"}`
	reqJSON := []byte(fmt.Sprintf(reqFormat, s.Code, s.Kind, mealType, date))

	doc, err := post(s, makeURL(s.Zone, linkMealWeekly), reqJSON)
	if err != nil {
		return
	}

	// get weekDietList json
	docDiet, _, _, err := jsonparser.Get(doc, "resultSVO", "weekDietList")
	if err != nil {
		return
	}

	rds := make([]resultDiet, 0, 3)
	if err = json.Unmarshal(docDiet, &rds); err != nil {
		return
	}

	if len(rds) < 3 {
		err = errors.New("schoolmeal: no diet in this week")
		return
	}

	// rds[0] == DateString
	// rds[1] == People
	// rds[2] == Content
	meals = make([]Meal, 7, 7)
	meals[0] = rdToMealWeek(rds[0].Sun, rds[1].Sun, rds[2].Sun, mealType)
	meals[1] = rdToMealWeek(rds[0].Mon, rds[1].Mon, rds[2].Mon, mealType)
	meals[2] = rdToMealWeek(rds[0].Tue, rds[1].Tue, rds[2].Tue, mealType)
	meals[3] = rdToMealWeek(rds[0].Wed, rds[1].Wed, rds[2].Wed, mealType)
	meals[4] = rdToMealWeek(rds[0].The, rds[1].The, rds[2].The, mealType)
	meals[5] = rdToMealWeek(rds[0].Fri, rds[1].Fri, rds[2].Fri, mealType)
	meals[6] = rdToMealWeek(rds[0].Sat, rds[1].Sat, rds[2].Sat, mealType)

	return
}
