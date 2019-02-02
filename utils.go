package schoolmeal

import (
	"fmt"
	"strconv"
	"time"
)

// Timestamp 함수는 GetWeekMeal 함수의 첫번째 인자로 쓰기 좋습니다.
// 인자로 받은 시간을 2018.11.18의 형태로 리턴합니다.
func Timestamp(date time.Time) (stamp string) {
	y, m, d := date.Date()

	mStr := m.String()
	if m < 10 {
		mStr = "0" + mStr
	}

	format := "%d.%s"
	if d < 10 {
		stamp = fmt.Sprintf(format+".%s", y, mStr, "0"+strconv.Itoa(d))
	} else {
		stamp = fmt.Sprintf(format+".%d", y, mStr, d)
	}
	return
}
