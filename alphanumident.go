package identnormalize

import (
	"unicode"
)

// StrictIdentifier normalize given original identifier into identifier
// matches pattern `[a-z]([a-z0-9]*)`.
// Characters out of acceptable range will be removed.
func AlphabetNumberOnlyIdentifier(originalIdent string, maxIdentifierLength int) string {
	result := make([]rune, 0, len(originalIdent))
	for _, ch := range originalIdent {
		currentResultSize := len(result)
		if currentResultSize >= maxIdentifierLength {
			break
		}
		if (ch > unicode.MaxASCII) ||
			(!(unicode.IsLetter(ch) || unicode.IsDigit(ch))) ||
			((currentResultSize == 0) && (!unicode.IsLetter(ch))) {
			continue
		}
		result = append(result, ch)
	}
	return string(result)
}

// StrictIdentifierPath normalize given originalIdentPath into normalizedIdentPath and normalizedFragments
// with each fragement matches pattern `[a-z]([a-z0-9]*)`.
// Characters out of acceptable range will be removed.
func AlphabetNumberOnlyIdentifierPath(
	originalIdentPath string, separatorCh rune, maxIdentifierPathLength int,
	identTransFunc IdentifierTransformFunc) (normalizedIdentPath string, normalizedFragments []string) {
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
			((len(resultFrags) == 0) && (!unicode.IsLetter(ch))) {
			continue
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
