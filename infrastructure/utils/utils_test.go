package utils

import (
	"encoding/json"
	"fmt"
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345", "54321!"}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

type ContainsTestCase struct {
	ArrayData      []string
	TextToFind     string
	ExpectedResult bool
}

func FuzzContains(f *testing.F) {
	testCases := []ContainsTestCase{
		{ArrayData: []string{"22", "20", "21"}, TextToFind: "22", ExpectedResult: true},
		{ArrayData: []string{"al", "ert0", "5fdx45"}, TextToFind: "al", ExpectedResult: true},
		{ArrayData: []string{"hds", "dfg0", "qwe+"}, TextToFind: "qwe+", ExpectedResult: true},
		{ArrayData: []string{"hds", "dfg0", "qwe+"}, TextToFind: "qwe", ExpectedResult: false},
		{ArrayData: []string{"hds", "dfg0", "qwe+"}, TextToFind: "ds", ExpectedResult: false},
	}

	for _, tc := range testCases {
		jsonByte, _ := json.Marshal(tc)
		fmt.Println(string(jsonByte))
		f.Add(jsonByte)
	}
	fmt.Println("Begining test")

	f.Fuzz(func(t *testing.T, jsonByte []byte) {
		var testCase ContainsTestCase
		fmt.Println("json:%S", string(jsonByte))
		json.Unmarshal(jsonByte, &testCase)

		if Contains(testCase.ArrayData, testCase.TextToFind) != testCase.ExpectedResult {
			t.Errorf("Fail searching the string %s", testCase.TextToFind)
		}
	})
}
