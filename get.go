package schoolmeal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/buger/jsonparser"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

func (s School) GetDayMeal(date string, mealType int) (m Meal, err error) {
	weekMeals, err := s.GetWeekMeal(date, mealType)
	if err != nil {
		return
	}

	parse := parseTime(date)
	m = weekMeals[parse.Weekday()]
	return
}

// GetWeekMeal 함수는 인자로 받는 날짜가 포함된 주의 급식이 담긴 []Meal{}을 리턴합니다.
func (s School) GetWeekMeal(date string, mealType int) (meals []Meal, err error) {
	reqFormat := `{"schulCode": "%s", "schulCrseScCode": %d, "schulMmealScCode": %d, "schYmd": "%s"}`
	reqJSON := []byte(fmt.Sprintf(reqFormat, s.Code, s.Kind, mealType, date))

	// POST
	req, err := http.NewRequest("POST", makeURL(s.Zone, linkMealWeekly), bytes.NewBuffer(reqJSON))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// if cookie expires, reload session
	if s.sess == nil || s.sess.Expires.Sub(time.Now()) <= 0 {
		err = s.reloadSession()
		if err != nil {
			return
		}
	}
	req.AddCookie(s.sess)

	// do request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Decode body -> byte
	doc, err := ioutil.ReadAll(res.Body)
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
	meals[0] = rdToMeal(rds[0].Sun, rds[1].Sun, rds[2].Sun, mealType)
	meals[1] = rdToMeal(rds[0].Mon, rds[1].Mon, rds[2].Mon, mealType)
	meals[2] = rdToMeal(rds[0].Tue, rds[1].Tue, rds[2].Tue, mealType)
	meals[3] = rdToMeal(rds[0].Wed, rds[1].Wed, rds[2].Wed, mealType)
	meals[4] = rdToMeal(rds[0].Thu, rds[1].Thu, rds[2].Thu, mealType)
	meals[5] = rdToMeal(rds[0].Fri, rds[1].Fri, rds[2].Fri, mealType)
	meals[6] = rdToMeal(rds[0].Sat, rds[1].Sat, rds[2].Sat, mealType)

	return
}
