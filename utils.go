package schoolmeal

import (
	"fmt"
	"time"
)

func Timestamp() (date string) {
	format := "%d.%d.%d"
	y, m, d := time.Now().Date()

	date = fmt.Sprintf(format, y, m, d)
	return
}
