package schoolmeal

// School 구조체는 급식 정보를 얻어오기 위해 필요한 학교 정보들을 필드로 가집니다.
type School struct {
	Name string // 학교 이름; GetWeekMeals()를 사용하기 위한 필수 필드는 아닙니다. Find() 사용 시 자동으로 채워집니다.
	Zone string // 학교를 관할하는 교육청
	Code string // 학교 코드
	Kind int    // 학교의 타입(유, 초, 중, 고)
}

// Meal 구조체는 급식 정보를 저장합니다.
type Meal struct {
	Date    string // 2018.11.30 형태의 타임스탬프
	Content string // 메뉴
}
