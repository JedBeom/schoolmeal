package schoolmeal

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
