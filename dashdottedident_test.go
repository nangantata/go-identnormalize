package identnormalize_test

import (
	"strings"
	"testing"

	identnormalize "github.com/nangantata/go-identnormalize"
)

func TestDashDottedIdentifierNormalize(t *testing.T) {
	var v string
	if v = identnormalize.DashDottedIdentifier("an_appl3_a-da.y", 16); v != "an_appl3_a-da.y" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.DashDottedIdentifier("an-ap9l.3_a-day", 8); v != "an-ap9l." {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.DashDottedIdentifier("an-appl3~a_day", 16); v != "an-appl3_a_day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.DashDottedIdentifier("An~Appl3.a-Da_y", 16); v != "An_Appl3.a-Da_y" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = strings.ToLower(identnormalize.DashDottedIdentifier("An-Appl3-a-Day", 16)); v != "an-appl3-a-day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.DashDottedIdentifier("*n-Appl3-&-Day", 16); v != "_n-Appl3-_-Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.DashDottedIdentifier("11-App.l3-a-Day", 16); v != "_1-App.l3-a-Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
}

func checkDashDottedIdentifierPathNormalize(t *testing.T,
	inputV string, separatorCh rune, maxIdentifierPathLength int, identTransFunc identnormalize.IdentifierTransformFunc,
	expV string, expF []string) (successVerify bool) {
	var v string
	var f []string
	v, f = identnormalize.DashDottedIdentifierPath(inputV, separatorCh, maxIdentifierPathLength, identTransFunc)
	return verifyIdentifierPathResult(t, inputV, expV, expF, v, f)
}

func TestDashDottedIdentifierPathNormalize(t *testing.T) {
	if !checkDashDottedIdentifierPathNormalize(t, "an/appl_3/a/day", '/', 16, nil,
		"an/appl_3/a/day", []string{"an", "appl_3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 16, nil,
		"an/appl-3/a/day", []string{"an", "appl-3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "an/appl-3/a/day_keeps/doc.tor5/10w", '/', 128, nil,
		"an/appl-3/a/day_keeps/doc.tor5/_0w", []string{"an", "appl-3", "a", "day_keeps", "doc.tor5", "_0w"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "An/Appl-3/a/Day/k339/10w/_.x", '/', 128, nil,
		"An/Appl-3/a/Day/k339/_0w/_.x", []string{"An", "Appl-3", "a", "Day", "k339", "_0w", "_.x"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "*n/Appl-3/&/Day", '/', 16, nil,
		"_n/Appl-3/_/Day", []string{"_n", "Appl-3", "_", "Day"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "1n/Appl-3/&/Day/./keeps", '/', 32, nil,
		"_n/Appl-3/_/Day/_/keeps", []string{"_n", "Appl-3", "_", "Day", "_", "keeps"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "1n/Appl-3/&/Day", '/', 16, strings.ToUpper,
		"_N/APPL-3/_/DAY", []string{"_N", "APPL-3", "_", "DAY"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 8, nil,
		"an/appl-", []string{"an", "appl-"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "an/a99l/", '/', 8, nil,
		"an/a99l", []string{"an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, ".wa/an/a99l/", '/', 16, nil,
		"_wa/an/a99l", []string{"_wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "-wa/an/a99l/", '/', 16, nil,
		"_wa/an/a99l", []string{"_wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "_wa/an/a99l/", '/', 16, nil,
		"_wa/an/a99l", []string{"_wa", "an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkDashDottedIdentifierPathNormalize(t, "an/appl.3//a//day//////", '/', 16, nil,
		"an/appl.3/a/day", []string{"an", "appl.3", "a", "day"}) {
		t.Error("check case failed")
	}
}
