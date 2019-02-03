package schoolmeal

import (
	"io/ioutil"
	"strconv"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

// Find 함수는 학교를 찾아 School을 리턴합니다. 여러 개의 학교가 찾아질 경우 첫번째 학교를 사용합니다.
func Find(zone, schoolName string) (school School, err error) {
	link := "https://par." + zone + ".go.kr/spr_ccm_cm01_100.do?kraOrgNm=" + schoolName

	resp, err := client.Get(link)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	code, err := jsonparser.GetString(data, "resultSVO", "orgDVOList", "[0]", "orgCode")
	if err != nil {
		err = errors.New("schoolmeal: Error while finding school code")
		return
	}

	kindString, err := jsonparser.GetString(data, "resultSVO", "orgDVOList", "[0]", "schulCrseScCode")
	if err != nil {
		err = errors.New("schoolmeal: Error while finding school kind")
		return
	}

	kind, err := strconv.Atoi(kindString)
	if err != nil {
		err = errors.New("schoolmeal: Error while converting kind to int from string")
		return
	}

	name, err := jsonparser.GetString(data, "resultSVO", "orgDVOList", "[0]", "kraOrgNm")
	if err != nil {
		err = errors.New("schoolmeal: Error while finding school name")
		return
	}

	school.Code = code
	school.Kind = kind
	school.Name = name
	school.Zone = zone

	return

}
