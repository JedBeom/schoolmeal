package schoolmeal

import (
	"testing"
)

func TestFind(t *testing.T) {
	school, err := Find(Seoul, "서울대학교사범대학부설고등학교")
	if err != nil {
		t.Error("Unexpected error:", err)
		t.Failed()
	}

	if school.Code != "A000003561" {
		t.Error("Expected A000003561 unexpected", school.Code)
	}

	if school.Kind != High {
		t.Error("Expected", High, "unexpected", school.Kind)
	}

	if school.Zone != Seoul {
		t.Error("Expected", Seoul, "unexpected", school.Zone)
	}

	if school.Name != "서울대학교사범대학부설고등학교" {
		t.Error("Expected 서울대학교사범대학부설고등학교, unexpected", school.Name)
	}
}
