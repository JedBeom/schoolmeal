package schoolmeal

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anaskhan96/soup"
)

var (
	client *http.Client
)

func init() {
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client = &http.Client{Transport: trans}
}

// GetWeekMeal 함수는 인자로 받는 날짜가 포함된 주의 급식이 담긴 string 슬라이스를 리턴합니다.
func (school School) GetWeekMeal(date string, mealtype int) (meals []Meal, err error) {
	originLink := "https://stu.%s.go.kr/sts_sci_md01_001.do?schulCode=%s&schulCrseScCode=%d&schulKndScCode=0%d&schMmealScCode=%d&schYmd=%s"

	link := fmt.Sprintf(originLink, school.Zone, school.SchoolCode, school.SchoolKindCode, school.SchoolKindCode, mealtype, date)

	resp, err := client.Get(link)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	doc := soup.HTMLParse(string(bodyBytes))
	td := doc.Find("tbody").FindAll("tr")[1].FindAll("td")

	for _, day := range td {
		var menu []string

		for i, food := range day.Children() {

			if i%2 == 0 {
				menu = append(menu, food.Pointer.Data)
			}

		}

		menus := strings.Join(menu, "\n")

		meal := Meal{Content: menus}

		meals = append(meals, meal)

	}

	th := doc.Find("thead").Find("tr").FindAll("th")
	for i, day := range th {
		meals[i].Date = day.Text()
	}

	return

}
