package identnormalize_test

import (
	"testing"
)

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
