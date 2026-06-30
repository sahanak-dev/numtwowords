package numtwowords_test

import (
	"testing"

	"github.com/sahanak-dev/numtwowords"
)

func TestIvalidNumber(t *testing.T) {
	_, err := numtwowords.Convert(numtwowords.MaxNum + 1)
	if err == nil {
		t.Log("Expected error for number out of range")
		t.Fail()
	}
	_, err = numtwowords.Convert(numtwowords.MinNum - 1)
	if err == nil {
		t.Log("Expected error for number out of range")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtwowords.Convert(0)
	if err != nil {
		t.Log("Unexpected error for zero")
		t.FailNow()
	}
	if result != "Not a valid number" {
		t.Logf("Expected 'Not a valid number', got '%v'", result)
		t.FailNow()
	}
}

func TestValidNumbers(t *testing.T) {
	units := [19]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	for i, v := range units {
		result, err := numtwowords.Convert(i + 1)
		if err != nil {
			t.Logf("Unexpected error for number %v got error %v", i+1, err)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%v', got '%v'", v, result)
			t.Fail()
		}
	}
}

func TestValidTens(t *testing.T) {
	testcases := map[int]string{
		34: "thirty four",
		45: "forty five",
		56: "fifty six",
		67: "sixty seven",
		78: "seventy eight",
		89: "eighty nine",
	}
	for k, v := range testcases {
		t.Logf("Testing number %v", k)
		result, err := numtwowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for number %v got error %v", k, err)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%v', got '%v'", v, result)
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		109: "one hundred and nine",
		333: "three hundred and thirty three",
		700: "seven hundred",
		340: "three hundred and forty",
	}
	for k, v := range testcases {
		t.Logf("Testing number %v", k)
		result, err := numtwowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for number %v got error %v", k, err)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%v', got '%v'", v, result)
			t.Fail()
		}
	}
}

func TestNegatives(t *testing.T) {
	testcases := map[int]string{-700: "minus seven hundred", -999: "minus nine hundred and ninety nine", -0: "Not a valid number"}
	for k, v := range testcases {
		t.Logf("Testing number %v", k)
		result, err := numtwowords.Convert(k)
		if err != nil {
			t.Logf("Unexpected error for number %v got error %v", k, err)
			t.Fail()
		}
		if result != v {
			t.Logf("Expected '%v', got '%v'", v, result)
			t.Fail()
		}
	}
}
