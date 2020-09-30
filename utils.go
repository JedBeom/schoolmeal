package schoolmeal

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if rd == "" || rd == " " {
		return nil
	}

	li := strings.Split(rd, "<br />")
	day, _ := strconv.Atoi(li[0])

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

func eventToSchedule(event string) (sches []Schedule, err error) {
	if len(event) == 0 || event == " " {
		err = errors.New("event empty")
		return
	}
	events := strings.Split(event, "|")
	for _, e := range events {
		fields := strings.Split(e, ":")
		if len(fields) != 4 {
			err = errors.New("event field not enough")
			return
		}

		sche := Schedule{}
		sche.DateString = fields[1]
		sche.Name = fields[2]
		sche.Type, err = strconv.Atoi(fields[3])
		if err != nil {
			return
		}

		if strings.Contains(sche.Name, "1학년") || strings.Contains(sche.Name, "/1") {
			sche.Grade1 = true
		} else if strings.Contains(sche.Name, "2학년") || strings.Contains(sche.Name, "/2") {
			sche.Grade2 = true
		} else if strings.Contains(sche.Name, "3학년") || strings.Contains(sche.Name, "/3") {
			sche.Grade3 = true
		} else if strings.Contains(sche.Name, "1, 2학년") || strings.Contains(sche.Name, "1,2학년") {
			sche.Grade1 = true
			sche.Grade2 = true
		} else if strings.Contains(sche.Name, "2, 3학년") || strings.Contains(sche.Name, "2,3학년") {
			sche.Grade2 = true
			sche.Grade3 = true
		} else if strings.Contains(sche.Name, "1, 3학년") || strings.Contains(sche.Name, "1,3학년") {
			sche.Grade1 = true
			sche.Grade3 = true
		}

		i := strings.Index(sche.Name, "(")
		if i > 0 {
			sche.Name = sche.Name[:i]
		}

		y, _ := strconv.Atoi(fields[1][0:4])
		m, _ := strconv.Atoi(fields[1][4:6])
		d, _ := strconv.Atoi(fields[1][6:8])
		sche.Date = time.Date(y, time.Month(m), d, 0, 0, 0, 0, seoul)

		sches = append(sches, sche)
	}

	return
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

func post(s School, url string, reqJSON []byte) (doc []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJSON))
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
	doc, err = ioutil.ReadAll(res.Body)
	return
}
