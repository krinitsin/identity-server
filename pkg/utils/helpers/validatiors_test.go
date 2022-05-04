package helpers

import "testing"

func TestIsValidCountry(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		want        bool
	}{
		{
			"Country test valid",
			"DEU",
			true,
		},
		{
			"Country test valid",
			"Germany",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidCountry(tt.countryCode); got != tt.want {
				t.Errorf("IsValidCountry() = %v, want %v", got, tt.want)
			}
		})
	}
}
