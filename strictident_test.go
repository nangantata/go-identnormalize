package identnormalize_test

import (
	"strings"
	"testing"

	identnormalize "github.com/nangantata/go-identnormalize"
)

func TestStrictIdentifierNormalize(t *testing.T) {
	var v string
	if v = identnormalize.StrictIdentifier("an_appl3_a_day", 16); v != "an_appl3_a_day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.StrictIdentifier("an_appl3_a_day", 8); v != "an_appl3" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.StrictIdentifier("an-appl3-a-day", 16); v != "an_appl3_a_day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.StrictIdentifier("An-Appl3-a-Day", 16); v != "An_Appl3_a_Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = strings.ToLower(identnormalize.StrictIdentifier("An-Appl3-a-Day", 16)); v != "an_appl3_a_day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.StrictIdentifier("*n-Appl3-&-Day", 16); v != "_n_Appl3___Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
	if v = identnormalize.StrictIdentifier("11-Appl3-a-Day", 16); v != "_1_Appl3_a_Day" {
		t.Errorf("unexpected result: [%s]", v)
	}
}

func verifyIdentifierPathResult(t *testing.T, inputV, expV string, expF []string, resultV string, resultF []string) (successVerify bool) {
	successVerify = true
	if expV != resultV {
		t.Errorf("unexpect identifier path result [input=%s]: [%s] vs. [%s]", inputV, expV, resultV)
		successVerify = false
	}
	if len(expF) != len(resultF) {
		t.Errorf("unexpect identifier fragements size [input=%s]: %d vs. %d", inputV, len(expF), len(resultF))
		successVerify = false
	}
	for idx, vF := range expF {
		if idx == len(resultF) {
			break
		}
		if vF != resultF[idx] {
			t.Errorf("unexpect identifier fragement [input=%s, index=%d]: [%s] vs. [%s]", inputV, idx, vF, resultF[idx])
			successVerify = false
		}
	}
	return
}

func checkStrictIdentifierPathNormalize(t *testing.T,
	inputV string, separatorCh rune, maxIdentifierPathLength int, identTransFunc identnormalize.IdentifierTransformFunc,
	expV string, expF []string) (successVerify bool) {
	var v string
	var f []string
	v, f = identnormalize.StrictIdentifierPath(inputV, separatorCh, maxIdentifierPathLength, identTransFunc)
	return verifyIdentifierPathResult(t, inputV, expV, expF, v, f)
}

func TestStrictIdentifierPathNormalize(t *testing.T) {
	if !checkStrictIdentifierPathNormalize(t, "an/appl_3/a/day", '/', 16, nil,
		"an/appl_3/a/day", []string{"an", "appl_3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 16, nil,
		"an/appl_3/a/day", []string{"an", "appl_3", "a", "day"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "An/Appl-3/a/Day", '/', 16, nil,
		"An/Appl_3/a/Day", []string{"An", "Appl_3", "a", "Day"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "*n/Appl-3/&/Day", '/', 16, nil,
		"_n/Appl_3/_/Day", []string{"_n", "Appl_3", "_", "Day"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "1n/Appl-3/&/Day", '/', 16, nil,
		"_n/Appl_3/_/Day", []string{"_n", "Appl_3", "_", "Day"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "1n/Appl-3/&/Day", '/', 16, strings.ToUpper,
		"_N/APPL_3/_/DAY", []string{"_N", "APPL_3", "_", "DAY"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "an/appl-3/a/day", '/', 8, nil,
		"an/appl_", []string{"an", "appl_"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "an/a99l/", '/', 8, nil,
		"an/a99l", []string{"an", "a99l"}) {
		t.Error("check case failed")
	}
	if !checkStrictIdentifierPathNormalize(t, "an/appl-3//a//day//////", '/', 16, nil,
		"an/appl_3/a/day", []string{"an", "appl_3", "a", "day"}) {
		t.Error("check case failed")
	}
}
