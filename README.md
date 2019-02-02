# Schoolmeal

```bash
go get -u github.com/JedBeom/schoolmeal
```

교육청 페이지를 크롤링해 학교 급식을 얻어올 수 있는 Golang 패키지입니다.

## Example

```go
package main

import (
    "fmt"
    "time"

    sm "github.com/JedBeom/schoolmeal"
)

func main() {

	school, err := sm.Find(sm.Seoul, "서울대학교사범대학부설고등학교")
	if err != nil {
		panic(err)
	}

    meals, err := school.GetWeekMeal(sm.Timestamp(time.Now()), sm.Lunch)
    if err != nil {
        panic(err)
    }

    fmt.Println(meals[time.Wednesday].Date, meals[time.Wednesday].Content)
}
```
