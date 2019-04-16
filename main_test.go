package main

import "testing"

func TestSpongeText(t *testing.T) {

	tables := []struct {
		text string
		want string
	}{
		{"Why are people talking like this",
			"WhY aRe PeOpLe TaLkInG lIkE tHiS"},
		{"HELLO THERE", "HeLlO tHeRe"},
	}

	for _, table := range tables {
		got := spongeMockText([]rune(table.text))
		if got != table.want {
			t.Errorf("Alternate case was incorrect, got: %s, want: %s.", table.text, table.want)
		}
	}
}
