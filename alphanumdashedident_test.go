package identnormalize_test

import (
	"strings"
	"testing"

	identnormalize "github.com/nangantata/go-identnormalize"
)

func TestAlphabetNumberDashOnlyIdentifierNormalize(t *testing.T) {
	var v string
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("an_appl3_a-da.y", 16); v != "anappl3a-day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("an-ap9l.3_a-day", 8); v != "an-ap9l3" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("an-appl3~a_day", 16); v != "an-appl3aday" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("An~Appl3.a-Da_y", 16); v != "AnAppl3a-Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = strings.ToLower(identnormalize.AlphabetNumberDashOnlyIdentifier("An_Appl3_a_Day", 16)); v != "anappl3aday" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("*n-Appl3-&-Day", 16); v != "n-Appl3--Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.AlphabetNumberDashOnlyIdentifier("11-App.l3-a-Day", 16); v != "Appl3-a-Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
}

func checkAlphabetNumberDashOnlyIdentifierPathNormalize(t *testing.T,
	inputV string, separatorCh rune, maxIdentifierPathLength int, identTransFunc identnormalize.IdentifierTransformFunc,
	expV string, expF []string) (successVerify bool) {
	var v string
	var f []string
	v, f = identnormalize.AlphabetNumberDashOnlyIdentifierPath(inputV, separatorCh, maxIdentifierPathLength, identTransFunc)
	return verifyIdentifierPathResult(t, inputV, expV, expF, v, f)
}

func TestAlphabetNumberDashOnlyIdentifierPathNormalize(t *testing.T) {
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/appl_3/a/day", '/', 16, nil,
		"an/appl3/a/day", []string{"an", "appl3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 16, nil,
		"an/appl-3/a/day", []string{"an", "appl-3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/appl-3/a/day_keeps/doc.tor5/10w", '/', 128, nil,
		"an/appl-3/a/daykeeps/doctor5/w", []string{"an", "appl-3", "a", "daykeeps", "doctor5", "w"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "An/Appl-3/a/Day/k339/10w/_.x", '/', 128, nil,
		"An/Appl-3/a/Day/k339/w/x", []string{"An", "Appl-3", "a", "Day", "k339", "w", "x"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "*n/Appl-3/&/Day", '/', 16, nil,
		"n/Appl-3/Day", []string{"n", "Appl-3", "Day"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "1n/Appl-3/&/Day/./keeps", '/', 32, nil,
		"n/Appl-3/Day/keeps", []string{"n", "Appl-3", "Day", "keeps"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "1n/Appl-3/&/Day", '/', 16, strings.ToUpper,
		"N/APPL-3/DAY", []string{"N", "APPL-3", "DAY"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 8, nil,
		"an/appl-", []string{"an", "appl-"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/a99l/", '/', 8, nil,
		"an/a99l", []string{"an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, ".wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "-wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "_wa/an/a99l/", '/', 16, nil,
		"wa/an/a99l", []string{"wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkAlphabetNumberDashOnlyIdentifierPathNormalize(t, "an/appl.3//a//day//////", '/', 16, nil,
		"an/appl3/a/day", []string{"an", "appl3", "a", "day"}) {
		t.Error("check case failed")
	}
}
