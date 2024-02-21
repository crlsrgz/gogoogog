package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"23:11:00", true},
		{"23:11:60", false},
		{"", false},
		{"13:11:00", true},
		{"3:25:00", true},
		{"3:25:00:", false},
	}
	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v. error should be nil", data.time, err)
		}
	}
}
