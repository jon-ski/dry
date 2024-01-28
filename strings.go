package dry

import (
	"strings"
)

// StrJoinN joins the parts with the separator, but only if the part is not empty.
func StrJoinN(sep string, parts ...string) string {
	var nonEmptyParts []string
	for _, part := range parts {
		if part != "" {
			nonEmptyParts = append(nonEmptyParts, part)
		}
	}
	return strings.Join(nonEmptyParts, sep)
}
