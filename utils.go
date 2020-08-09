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
	seoul, _ = time.LoadLocation("Asia/Seoul")
}

// Timestamp 함수는 GetWeekMeal 함수의 첫번째 인자로 쓰기 좋습니다.
// 인자로 받은 시간을 2018.11.18의 형태로 리턴합니다.
func Timestamp(date time.Time) (stamp string) {
	y, m, d := date.Date()

	mStr := ""
	if m < 10 {
		mStr = fmt.Sprintf("0%d", m)
	} else {
		mStr = strconv.Itoa(int(m))
	}

	format := "%d.%s"
	if d < 10 {
		stamp = fmt.Sprintf(format+".%s", y, mStr, "0"+strconv.Itoa(d))
	} else {
		stamp = fmt.Sprintf(format+".%d", y, mStr, d)
	}
	return
}

func makeURL(zone, link string) string {
	return "https://stu." + zone + ".go.kr/" + link
}

func rdToMeal(date, peopleStr, content string, t int) (m Meal) {
	m.DateString = date
	m.Date = parseTime(date)
	if content != "" {
		m.Content = strings.Replace(content, "<br />", "\n", -1)
		m.Content = m.Content[:len(m.Content)-len("\n")]
	}
	if peopleStr != "" && peopleStr != " " {
		m.People, _ = strconv.Atoi(peopleStr[:len(peopleStr)-len("명")])
	}
	m.Type = t
	return
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
