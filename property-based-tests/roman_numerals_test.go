package propertybasedtests

import (
	"fmt"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Num   uint16
		Roman string
	}{
		{Num: 1, Roman: "I"},
		{Num: 2, Roman: "II"},
		{Num: 3, Roman: "III"},
		{Num: 4, Roman: "IV"},
		{Num: 5, Roman: "V"},
		{Num: 6, Roman: "VI"},
		{Num: 7, Roman: "VII"},
		{Num: 8, Roman: "VIII"},
		{Num: 9, Roman: "IX"},
		{Num: 10, Roman: "X"},
		{Num: 14, Roman: "XIV"},
		{Num: 18, Roman: "XVIII"},
		{Num: 20, Roman: "XX"},
		{Num: 39, Roman: "XXXIX"},
		{Num: 40, Roman: "XL"},
		{Num: 47, Roman: "XLVII"},
		{Num: 49, Roman: "XLIX"},
		{Num: 50, Roman: "L"},
		{Num: 100, Roman: "C"},
		{Num: 90, Roman: "XC"},
		{Num: 400, Roman: "CD"},
		{Num: 500, Roman: "D"},
		{Num: 900, Roman: "CM"},
		{Num: 1000, Roman: "M"},
		{Num: 1984, Roman: "MCMLXXXIV"},
		{Num: 3999, Roman: "MMMCMXCIX"},
		{Num: 2014, Roman: "MMXIV"},
		{Num: 1006, Roman: "MVI"},
		{Num: 798, Roman: "DCCXCVIII"},
	}
)

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted into %q", test.Num, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Num)

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Num), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Num {
				t.Errorf("got %d, want %d", got, test.Num)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(num uint16) bool {
		if num > 3999 {
			return true
		}
		t.Log("testing", num)
		roman := ConvertToRoman(num)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == num
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", nil)
	}
}
