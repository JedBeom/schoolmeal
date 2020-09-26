package schoolmeal

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/buger/jsonparser"
)

// GetMonthSchedule 함수는 월별 일정을 가져옵니다.
func (s School) GetMonthSchedule(year, month int) (schedules []Schedule, err error) {
	reqFormat := `{"schulCode": "%s", "schulCrseScCode": %d, "schulKndScCode": "0%d", "ay": "%d", "mm": "%02d"}`
	reqJSON := []byte(fmt.Sprintf(reqFormat, s.Code, s.Kind, s.Kind, year, month))

	doc, err := post(s, makeURL(s.Zone, linkScheduleMonthly), reqJSON)
	if err != nil {
		return
	}

	docSche, _, _, err := jsonparser.Get(doc, "resultSVO", "selectMonth")
	if err != nil {
		return
	}

	docSche = docSche[1 : len(docSche)-1] // remove [ ]
	docs := bytes.Split(docSche, []byte("},"))

	schedules = make([]Schedule, 0, 31)
	for _, elem := range docs {
		elem = append(elem, 125) // add } behind
		for i := 1; i <= 7; i++ {
			// add } behind
			event, err := jsonparser.GetString(elem, fmt.Sprintf("event%d", i))

			if err != nil {
				if errors.Is(err, jsonparser.KeyPathNotFoundError) {
					break
				}

				continue
			}

			sches, err := eventToSchedule(event)
			if err != nil {
				return nil, err
			}

			schedules = append(schedules, sches...)
		}
	}

	return
}
