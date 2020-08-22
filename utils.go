package schoolmeal

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	seoul *time.Location
)

func init() {
	var err error
	seoul, err = time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
}

// Timestamp 함수는 GetWeekMeal 함수의 첫번째 인자로 쓰기 좋습니다.
// 인자로 받은 시간을 2018.11.18의 형태로 리턴합니다.
func Timestamp(date time.Time) (stamp string) {
	y, m, d := date.Date()

	format := "%d.%02d.%02d"
	stamp = fmt.Sprintf(format, y, m, d)
	return
}

func makeURL(zone, link string) string {
	return "https://stu." + zone + ".go.kr/" + link
}

func rdToMealWeek(date, peopleStr, content string, t int) (m Meal) {
	m.DateString = date
	m.Date = parseTime(date)
	if content != "" && content != " " {
		m.Content = strings.Replace(content, "<br />", "\n", -1)
		m.Content = m.Content[:len(m.Content)-len("\n")]
	}
	if peopleStr != "" && peopleStr != " " {
		m.People, _ = strconv.Atoi(peopleStr[:len(peopleStr)-len("명")])
	}
	m.Type = t
	return
}

func rdToMealMonth(year, month int, rd string) []Meal {
	if rd == "" {
		return nil
	}

	li := strings.Split(rd, "<br />")
	day, _ := strconv.Atoi(li[0])

	if len(li) == 1 {
		return nil
	}

	start := 1
	end := 0
	ms := make([]Meal, 3, 3)

	mealType := []string{"[조식]", "[중식]", "[석식]"}
	for ti, t := range mealType {
		ms[ti].Type = ti + 1
		ms[ti].DateString = fmt.Sprintf("%d.%02d.%02d", year, month, day)
		ms[ti].Date = parseTime(ms[ti].DateString)
		hangul := weekdayHangul[ms[ti].Date.Weekday()]
		ms[ti].DateString += "(" + hangul + ")"

		if len(li) == start || li[start] != t {
			continue
		}
		start++
		end = start
		for end < len(li) && !in(li[end], mealType) {
			end++
		}
		ms[ti].Content = strings.Join(li[start:end], "\n")
		start = end
	}

	return ms
}

func in(s string, l []string) bool {
	for _, r := range l {
		if s == r {
			return true
		}
	}

	return false
}

func parseTime(t string) time.Time {
	if len(t) < 10 {
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	}
	y, _ := strconv.Atoi(t[0:4])
	m, _ := strconv.Atoi(t[5:7])
	d, _ := strconv.Atoi(t[8:10])
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, seoul)
}

func appendIfNotNil(mm *[][]Meal, m []Meal) {
	if m == nil {
		return
	}

	*mm = append(*mm, m)
}
