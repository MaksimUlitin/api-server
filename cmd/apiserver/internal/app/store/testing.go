package store

import (
	"fmt"
	"strings"
	"testing"
)

func Tester(t *testing.T, DataBaseURL string) (*Store, func(...string)) {
	t.Helper()

	Config := NewConfig()
	Config.DataBaseURL = DataBaseURL
	s := New(Config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}

}
