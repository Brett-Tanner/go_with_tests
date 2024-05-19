package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"converts 1 to I", 1, "I"},
		{"converts 2 to II", 2, "II"},
		{"converts 3 to III", 3, "III"},
		{"converts 4 to IV", 4, "IV"},
		{"converts 5 to V", 5, "V"},
		{"converts 6 to VI", 6, "VI"},
		{"converts 7 to VII", 7, "VII"},
		{"converts 8 to VIII", 8, "VIII"},
		{"converts 9 to IX", 9, "IX"},
		{"converts 10 to X", 10, "X"},
		{"converts 15 to XV", 15, "XV"},
		{"converts 23 to XXIII", 23, "XXIII"},
		{"converts 39 to IXXXX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
