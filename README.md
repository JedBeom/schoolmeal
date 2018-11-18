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
    school := sm.School{
        SchoolCode:     "Q100005451", // 학교 코드
        SchoolKindCode: sm.Middle, // 학교 타입(유치원, 초, 중, 고)
        Zone:           sm.Jeonnam, // 학교를 관할하는 교육청
    }

    meals, err := school.GetWeekMeal("2018.11.15", sm.Lunch)
    if err != nil {
        panic(err)
    }

    fmt.Println(meals[time.Wednesday])
}
```

`SchoolCode`는 [이 사이트](https://www.meatwatch.go.kr/biz/bm/sel/schoolListPopup.do)에서 얻으실 수 있습니다.
