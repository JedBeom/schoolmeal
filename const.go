package schoolmeal

// GetWeekMeal 함수의 mealType 인자에 사용하는 상수들입니다.
const (
	Breakfast = iota + 1 // 아침
	Lunch                // 점심
	Dinner               // 저녁
)

// School 구조체의 Code 필드에 사용할 수 있습니다.
const (
	Kindergarten = iota + 1 // 유치원
	Elementary              // 초등학교
	Middle                  // 중학교
	High                    // 고등학교
)

// School 구조체의 Zone 필드에 사용 가능합니다.
// 학교를 관할하는 교육청에 맞게 사용하면 됩니다.
const (
	Seoul     = "sen" // 서울특별시교육청
	Busan     = "pen" // 부산광역시교육청
	Daegu     = "dge" // 대구광역시교육청
	Incheon   = "ice" // 인천광역시교육청
	Gwangju   = "gen" // 광주광역시교육청
	Daejeon   = "dje" // 대전광역시교육청
	Ulsan     = "use" // 울산광역시교육청
	Sejong    = "sje" // 세종특별자치시교육청
	Gyeonggi  = "goe" // 경기도교육청
	Gangwon   = "gwe" // 강원도교육청
	Chungbuk  = "cbe" // 충청북도교육청
	Chungnam  = "cne" // 충남북도교육청
	Jeonbuk   = "jbe" // 전라북도교육청
	Jeonnam   = "jne" // 전라남도교육청
	Gyeongbuk = "kbe" // 경상북도교육청
	Geyongnam = "gne" // 경상남도교육청
	Jeju      = "jje" // 제주특별자치도교육청
)
