package identnormalize

// IdentifierTransformFunc map given originalIdent into resultIdent.
// Examples of transform functions are: strings.ToUpper, string.ToLower.
type IdentifierTransformFunc func(originalIdent string) (resultIdent string)

func noopIdentTransform(v string) string {
	return v
}
