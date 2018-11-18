package schoolmeal

type School struct {
	SchoolCode     string
	SchoolKindCode int
	Zone           string
}

const (
	Breakfast = iota + 1
	Lunch
	Dinner
)

const (
	Kindergarden = iota + 1
	Elementary
	Middle
	High
)

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
