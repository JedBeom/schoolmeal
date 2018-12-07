package schoolmeal

import (
	"fmt"
	"time"
)

// Timestamp 함수는 GetWeekMeal 함수의 첫번째 인자로 쓰기 좋습니다.
// 인자로 받은 시간을 2018.11.18의 형태로 리턴합니다.
func Timestamp(date time.Time) (stamp string) {
	format := "%d.%d.%d"
	y, m, d := date.Date()

	stamp = fmt.Sprintf(format, y, m, d)
	return
}
