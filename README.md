# Schoolmeal
[![GoDoc](https://godoc.org/github.com/JedBeom/schoolmeal?status.svg)](https://pkg.go.dev/github.com/JedBeom/schoolmeal)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jedbeom/schoolmeal)](https://goreportcard.com/report/github.com/Jedbeom/schoolmeal)

```bash
go get -u github.com/JedBeom/schoolmeal
```

교육청 페이지를 크롤링해 학교 급식을 얻어올 수 있는 Golang 패키지입니다.

## Example

```go
package main

import (
    "fmt"
    sm "github.com/JedBeom/schoolmeal"
    "time"
)

func main() {

    // 학교 정보를 가져온다
    school, err := sm.Find(sm.Seoul, "서울대학교사범대학부설고등학교")
    if err != nil {
        panic(err)
    }

    // 해당 학교의 일주일치 급식을 가져옴
    meals, err := school.GetWeekMeal(sm.Timestamp(time.Now()), sm.Lunch)
    if err != nil {
        panic(err)
    }

    // 수요일의 급식의 날짜와 급식 메뉴를 출력
    fmt.Println(meals[time.Wednesday].Date, meals[time.Wednesday].Content)
}
```

더 많은 정보는 [godoc](https://pkg.go.dev/github.com/JedBeom/schoolmeal)을 참고하세요.
