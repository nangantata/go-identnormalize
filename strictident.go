package identnormalize

import (
	"unicode"
)

// StrictIdentifier normalize given original identifier into identifier
// matches pattern `[a-zA-Z_]([a-zA-Z0-9_]*)`.
func StrictIdentifier(originalIdent string, maxIdentifierLength int) string {
	result := make([]rune, 0, len(originalIdent))
	for idx, ch := range originalIdent {
		if idx >= maxIdentifierLength {
			break
		}
		if (ch > unicode.MaxASCII) ||
			(!(unicode.IsLetter(ch) || unicode.IsDigit(ch))) ||
			((idx == 0) && unicode.IsDigit(ch)) {
			ch = '_'
		}
		result = append(result, ch)
	}
	return string(result)
}

// StrictIdentifierPath normalize given originalIdentPath into normalizedIdentPath and normalizedFragments
// with each fragement matches pattern `[a-zA-Z_]([a-zA-Z0-9_]*)`.
func StrictIdentifierPath(
	originalIdentPath string, separatorCh rune, maxIdentifierPathLength int,
	identTransFunc IdentifierTransformFunc) (normalizedIdentPath string, normalizedFragments []string) {
	if maxIdentifierPathLength <= 0 {
		return
	}
	if identTransFunc == nil {
		identTransFunc = noopIdentTransform
	}
	lenOrigIdentPath := len(originalIdentPath)
	resultPath := make([]rune, 0, lenOrigIdentPath)
	resultFrags := make([]rune, 0, lenOrigIdentPath)
	normalizedFragments = make([]string, 0, 4)
	for _, ch := range originalIdentPath {
		if len(resultPath) >= maxIdentifierPathLength {
			break
		}
		if ch == separatorCh {
			if len(resultFrags) == 0 {
				continue
			}
			frag := identTransFunc(string(resultFrags))
			normalizedFragments = append(normalizedFragments, frag)
			resultFrags = resultFrags[:0]
			resultPath = append(resultPath, separatorCh)
			continue
		}
		if (ch > unicode.MaxASCII) ||
			(!(unicode.IsLetter(ch) || unicode.IsDigit(ch))) ||
			((len(resultFrags) == 0) && unicode.IsDigit(ch)) {
			ch = '_'
		}
		resultFrags = append(resultFrags, ch)
		resultPath = append(resultPath, ch)
	}
	if len(resultFrags) > 0 {
		frag := identTransFunc(string(resultFrags))
		normalizedFragments = append(normalizedFragments, frag)
	}
	if resultSize := len(resultPath); (resultSize > 0) && (resultPath[resultSize-1] == separatorCh) {
		resultPath = resultPath[:resultSize-1]
	}
	normalizedIdentPath = identTransFunc(string(resultPath))
	return
}
