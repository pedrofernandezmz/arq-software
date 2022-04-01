package div

import "testing"

func TestDivisionOK(t *testing.T) {
	_, err := Division(8, 4)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDivisionError(t *testing.T) {
	_, err := Division(8, 0)
	if err == nil {
		t.Error(err)
		return
	}
}
