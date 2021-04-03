package identnormalize_test

import (
	"strings"
	"testing"

	identnormalize "github.com/nangantata/go-identnormalize"
)

func TestAlphabetNumberOnlyIdentifierNormalize(t *testing.T) {
	var v string
	if v = identnormalize.AlphabetNumberOnlyIdentifier("an_appl3_a-da.y", 16); v != "anappl3aday" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberOnlyIdentifier("an-ap9l.3_a-day", 8); v != "anap9l3a" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberOnlyIdentifier("an-appl3~a_day", 16); v != "anappl3aday" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberOnlyIdentifier("An~Appl3.a-Da_y", 16); v != "AnAppl3aDay" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = strings.ToLower(identnormalize.AlphabetNumberOnlyIdentifier("An-Appl3-a-Day", 16)); v != "anappl3aday" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberOnlyIdentifier("*n-Appl3-&-Day", 16); v != "nAppl3Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberOnlyIdentifier("11-App.l3-a-Day", 16); v != "Appl3aDay" {
		t.Errorf("unexpected result: [%s]", v)
	}
}

func checkAlphabetNumberOnlyIdentifierPathNormalize(t *testing.T,
	inputV string, separatorCh rune, maxIdentifierPathLength int, identTransFunc identnormalize.IdentifierTransformFunc,
	expV string, expF []string) (successVerify bool) {
	var v string
	var f []string
	v, f = identnormalize.AlphabetNumberOnlyIdentifierPath(inputV, separatorCh, maxIdentifierPathLength, identTransFunc)
	return verifyIdentifierPathResult(t, inputV, expV, expF, v, f)
}

func TestAlphabetNumberOnlyIdentifierPathNormalize(t *testing.T) {
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/appl_3/a/day", '/', 16, nil,
		"an/appl3/a/day", []string{"an", "appl3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 16, nil,
		"an/appl3/a/day", []string{"an", "appl3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/appl-3/a/day_keeps/doc.tor5/10w", '/', 128, nil,
		"an/appl3/a/daykeeps/doctor5/w", []string{"an", "appl3", "a", "daykeeps", "doctor5", "w"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "An/Appl-3/a/Day/k339/10w/_.x", '/', 128, nil,
		"An/Appl3/a/Day/k339/w/x", []string{"An", "Appl3", "a", "Day", "k339", "w", "x"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "*n/Appl-3/&/Day", '/', 16, nil,
		"n/Appl3/Day", []string{"n", "Appl3", "Day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "1n/Appl-3/&/Day/./keeps", '/', 32, nil,
		"n/Appl3/Day/keeps", []string{"n", "Appl3", "Day", "keeps"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "1n/Appl-3/&/Day", '/', 16, strings.ToUpper,
		"N/APPL3/DAY", []string{"N", "APPL3", "DAY"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 8, nil,
		"an/appl3", []string{"an", "appl3"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/a99l/", '/', 8, nil,
		"an/a99l", []string{"an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, ".wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "-wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "_wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberOnlyIdentifierPathNormalize(t, "an/appl.3//a//day//////", '/', 16, nil,
		"an/appl3/a/day", []string{"an", "appl3", "a", "day"}) {
		t.Error("check case failed")
	}
}
