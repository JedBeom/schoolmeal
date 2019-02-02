package schoolmeal

// School 구조체는 급식 정보를 얻어오기 위해 필요한 학교 정보들을 필드로 가집니다.
type School struct {
	Code string // 학교 코드
	Name string // 학교 이름
	Kind int    // 학교의 타입(유, 초, 중, 고)
	Zone string // 학교를 관할하는 교육청
}

type Meal struct {
	Date    string
	Content string
}
