package bankocr

import (
	"testing"
)

//run test
func TestArabicToRoman(t *testing.T) {

	s := ArabicToRoman(0)
	if s != "" {
		t.Errorf("0 must return empty string, got %s", s)
	}

	s2 := ArabicToRoman(2)
	if s2 != "II" {
		t.Errorf("2 must return II , got %s", s2)
	}

	s3 := ArabicToRoman(11)
	if s3 != "XXX" {
		t.Errorf("Not cero and bigger than 3 must return XXX , got %s", s2)
	}
}
