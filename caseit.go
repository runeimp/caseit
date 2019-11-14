package caseit

import (
	"fmt"
	"strings"
	"unicode"
)

func CamelCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.Title(r)
	r = fmt.Sprintf("%s%s", strings.ToLower(r[0:1]), r[1:])
	r = strings.ReplaceAll(r, " ", "")

	return r
}

func DotLowerCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToLower(r)
	r = strings.ReplaceAll(r, " ", ".")

	return r
}

func DotTitleCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.Title(r)
	r = strings.ReplaceAll(r, " ", ".")

	return r
}

func DotUpperCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToUpper(r)
	r = strings.ReplaceAll(r, " ", ".")

	return r
}

func KebabCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToLower(r)
	r = strings.ReplaceAll(r, " ", "-")

	return r
}

func KebabTitleCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.Title(r)
	r = strings.ReplaceAll(r, " ", "-")

	return r
}

func KebabUpperCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToUpper(r)
	r = strings.ReplaceAll(r, " ", "-")

	return r
}

func NoDelimiterCamelCase(s string) (r string) {
	return CamelCase(s)
}

func NoDelimiterLowerCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToLower(r)
	r = strings.ReplaceAll(r, " ", "")

	return r
}

func NoDelimiterPascalCase(s string) (r string) {
	return PascalCase(s)
}

func NoDelimiterUpperCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToUpper(r)
	r = strings.ReplaceAll(r, " ", "")

	return r
}

func PascalCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.Title(r)
	r = strings.ReplaceAll(r, " ", "")

	return r
}

func preprocessString(s string) (r string) {
	p := ' '

	for _, c := range s {
		r = r + Separatem(p, c)
		p = c
	}

	return r
}

// Separatem takes p (previous), and c (current) and compares them to determine how to separate them, if necessary.
func Separatem(p, c rune) (r string) {
	// If p and c are different cases return them concatenated with a space between
	// If c is a dot, hyphen, or underscore convert to a space
	// NOTE: Consider scenarios with other delimiters and where to handle them. Probably not here.
	switch c {
	case '.', '-', '_', ' ':
		// fmt.Printf("%q is a delimiter\n", c)
		r = " "
	default:
		if unicode.IsLetter(p) && unicode.IsLetter(c) {
			// if unicode.IsLower(p) && unicode.IsUpper(c) || unicode.IsUpper(p) && unicode.IsLower(c) {
			if unicode.IsLower(p) && unicode.IsUpper(c) {
				// fmt.Printf("%q and %q is are opposite case\n", p, c)
				r = fmt.Sprintf(" %s", string(c))
			} else {
				// fmt.Printf("%q is a letter\n", c)
				r = string(c)
			}
		} else {
			// fmt.Printf("%q or %q is not a letter\n", p, c)
			r = string(c)
		}
	}
	return r
}

func SnakeCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToLower(r)
	r = strings.ReplaceAll(r, " ", "_")

	return r
}

func SnakeTitleCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.Title(r)
	r = strings.ReplaceAll(r, " ", "_")

	return r
}

func SnakeUpperCase(s string) (r string) {
	r = preprocessString(s)
	r = strings.ToUpper(r)
	r = strings.ReplaceAll(r, " ", "_")

	return r
}

func SpaceLowerCase(s string) (r string) {
	r = preprocessString(s)
	return strings.ToLower(r)
}

func SpaceSmartCase(s string) (r string) {
	r = preprocessString(s)
	// r = strings.Title(r)
	r = "Smart Title Case not yet implemented"

	return r
}

func SpaceTitleCase(s string) (r string) {
	r = preprocessString(s)
	return strings.Title(r)
}

func SpaceUpperCase(s string) (r string) {
	r = preprocessString(s)
	return strings.ToUpper(r)
}
