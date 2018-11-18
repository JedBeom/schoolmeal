package schoolmeal

// School 구조체는 급식 정보를 얻어오기 위해 필요한 학교 정보들을 필드로 가집니다.
type School struct {
	SchoolCode     string // 학교 코드
	SchoolKindCode int    // 학교의 타입(유, 초, 중, 고)
	Zone           string // 학교를 관할하는 교육청
}

// GetWeekMeal 함수의 두번째 인자로 들어가는 mealtype에 사용하는 상수들입니다.
const (
	Breakfast = iota + 1 // 아침
	Lunch                // 점심
	Dinner               // 저녁
)

// School 구조체의 SchoolKindCode에 사용할 수 있습니다.
const (
	Kindergarden = iota + 1 // 유치원
	Elementary              // 초등학교
	Middle                  // 중학교
	High                    // 고등학교
)

// School 구조체의 Zone 필드에 사용 가능합니다.
const (
	Seoul     = "sen"
	Busan     = "pen"
	Daegu     = "dge"
	Incheon   = "ice"
	Gwangju   = "gen"
	Daejeon   = "dje"
	Ulsan     = "use"
	Sejong    = "sje"
	Gyeonggi  = "goe"
	Gangwon   = "gwe"
	Chungbuk  = "cbe"
	Chungnam  = "cne"
	Jeonbuk   = "jbe"
	Jeonnam   = "jne"
	Gyeongbuk = "kbe"
	Geyongnam = "gne"
	Jeju      = "jje"
)
