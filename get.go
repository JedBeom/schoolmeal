package schoolmeal

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/anaskhan96/soup"
)

var (
	client    *http.Client
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) GoParser/0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
)

func init() {
	trans := &http.Transport{
		// Korean government sites don't have secured certifications
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{Transport: trans}
}

// GetWeekMeal 함수는 인자로 받는 날짜가 포함된 주의 급식이 담긴 string 슬라이스를 리턴합니다.
func (school School) GetWeekMeal(date string, mealType int) (meals []Meal, err error) {
	originLink := "https://stu.%s.go.kr/sts_sci_md01_001.do?schulCode=%s&schulCrseScCode=%d&schulKndScCode=0%d&schMmealScCode=%d&schYmd=%s"
	// https://stu.sen.go.kr/sts_sci_md01_001.do?schulCode=A000003561&schulCrseScCode=4&schulKndScCode=04&schMmealScCode=2&schYmd=2018.11.30

	link := fmt.Sprintf(originLink, school.Zone, school.Code, school.Kind, school.Kind, mealType, date)

	// Make new request
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		err = errors.Wrap(err, "schoolmeal")
		return
	}
	// Set user-agent for Chrome
	req.Header.Set("User-Agent", userAgent)

	// Get body
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "schoolmeal")
		return
	}

	// Convert body to bytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "schoolmeal")
		return
	}

	// Parse html
	doc := soup.HTMLParse(string(bodyBytes))
	// <thead>
	thead := doc.Find("thead")
	// Get <td>
	mealExistTD := thead.Find("td")
	// Prevent runtime pointer error
	if mealExistTD.Pointer != nil {
		if mealExistTD.Text() == "자료가 없습니다." {
			err = errors.New("schoolmeal: Can't get meals; Might be wrong arguments?")
			return
		}
	}

	// <td>
	tds := doc.FindAll("td")
	// length should be greater than 14
	if len(tds) < 14 {
		err = errors.New("schoolmeal: length of meals is too short")
		return
	}

	// tds[0:7] are for the number of people who eat
	for _, day := range tds[7:] {
		var menu string

		// soup thinks menu strings as 'tags', so we should get children
		for i, food := range day.Children() {
			// i%2 != 0 -> <br>
			if i%2 == 0 {
				menu += food.Pointer.Data + "\n"
			}
		}

		// Remove last '\n'
		if len(menu) > 1 {
			menu = menu[:len(menu)-1]
		}

		meal := Meal{Content: menu}
		meals = append(meals, meal)
	}

	th := thead.FindAll("th")
	if len(th) < 2 {
		err = errors.New("schoolmeal: Index out of range")
		return
	}

	for i, day := range th[1:] {
		meals[i].Date = day.Text()
	}

	return

}
