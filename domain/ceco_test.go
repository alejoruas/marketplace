package domain_test

import (
	"marketplace/domain"
	"strings"
	"testing"
)

func FuzzCreateCeco(f *testing.F) {
	testCases := []string{"222003248", "288145003", "218145005", "208145005", "268145005", "008145005", "008145"}

	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, source string) {
		ceco, err := domain.NewCeco(source)

		if (err != nil || ceco.Ceco() != source) && !strings.HasPrefix(source, "00") {
			t.Errorf("Invalid behavior with %s", source)
		}
	})
}
