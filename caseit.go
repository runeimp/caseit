package caseit

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	// titlecase "github.com/igorsobreira/titlecase"
	// titlecase "github.com/KenjiTakahashi/tu/titlecase"
	"github.com/runeimp/titlecase"
)

const Version = "1.0.0"

var (
	spaceRE      = regexp.MustCompile(`\s+`)
	UseCodeBreak = false
)

// Idea modified from
// https://github.com/serenize/snaker/blob/master/snaker.go#L102
var commonInitialismsAndOddballs = map[string]string{
	"ACL":        "ACL",
	"API":        "API",
	"ASCII":      "ASCII",
	"CIA":        "CIA",
	"CPU":        "CPU",
	"CSS":        "CSS",
	"DNS":        "DNS",
	"EOF":        "EOF",
	"ETA":        "ETA",
	"FBI":        "FBI",
	"GPU":        "GPU",
	"GRPC":       "gRPC",
	"GUID":       "GUID",
	"GRAPHQL":    "GraphQL",
	"HTML":       "HTML",
	"HTTP":       "HTTP",
	"HTTPS":      "HTTPS",
	"ID":         "ID",
	"IP":         "IP",
	"IRS":        "IRS",
	"JAVASCRIPT": "JavaScript",
	"JSON":       "JSON",
	"LHS":        "LHS",
	"NSA":        "NSA",
	"OAUTH":      "OAuth",
	"OPENID":     "OpenID",
	"OS":         "OS",
	"QPS":        "QPS",
	"RAM":        "RAM",
	"RHS":        "RHS",
	"RPC":        "RPC",
	"SLA":        "SLA",
	"SMTP":       "SMTP",
	"SQL":        "SQL",
	"SSH":        "SSH",
	"TCP":        "TCP",
	"TLS":        "TLS",
	"TTL":        "TTL",
	"UDP":        "UDP",
	"UI":         "UI",
	"UID":        "UID",
	"URI":        "URI",
	"URL":        "URL",
	"UTF8":       "UTF8",
	"UTF16":      "UTF16",
	"UTF32":      "UTF32",
	"UUID":       "UUID",
	"VM":         "VM",
	"VR":         "VR",
	"XML":        "XML",
	"XMPP":       "XMPP",
	"XSRF":       "XSRF",
	"XSS":        "XSS",
}

func CamelCase(s string) (r string) {
	return DelimiterCase(s, "", "Force", true)
}

func CamelCoderCase(s string) (r string) {
	return DelimiterCase(s, "", "Coder", true)
}

func capitalize(s string) (r string) {
	for _, word := range strings.Fields(s) {
		r = r + " " + capitalizeWord(word)
	}
	return r[1:]
}

func capitalizeWord(w string) string {
	upper := strings.ToUpper(w)
	if word, ok := commonInitialismsAndOddballs[upper]; ok == true {
		return word
	}
	word := strings.ToLower(w)
	word = strings.Title(word)
	return word
}

func DelimiterCase(s, d, m string, f bool) (r string) {
	r = PreprocessString(s)

	switch strings.ToLower(m) {
	case "coder":
		r = capitalize(r)
	case "force":
		r = strings.ToLower(r)
		r = strings.Title(r)
	case "lower":
		r = strings.ToLower(r)
	case "smart":
		r = titlecase.Convert(r)
	case "title":
		r = strings.Title(r)
	case "upper":
		r = strings.ToUpper(r)
	default:
		e := fmt.Errorf("Unknown mode: %q", m)
		// fmt.Printf("ERROR: %s\n", e.Error())
		panic(e.Error())
		return r
	}

	if f {
		r = fmt.Sprintf("%s%s", strings.ToLower(r[0:1]), r[1:])
	}

	r = strings.ReplaceAll(r, " ", d)

	return r
}

func DotCoderCase(s string) (r string) {
	return DelimiterCase(s, ".", "Coder", false)
}

func DotLowerCase(s string) (r string) {
	return DelimiterCase(s, ".", "lower", false)
}

func DotForceCase(s string) (r string) {
	return DelimiterCase(s, ".", "Force", false)
}

func DotTitleCase(s string) (r string) {
	return DelimiterCase(s, ".", "Title", false)
}

func DotSmartCase(s string) (r string) {
	return DelimiterCase(s, ".", "Smart", false)
}

func DotUpperCase(s string) (r string) {
	return DelimiterCase(s, ".", "UPPER", false)
}

func KebabCoderCase(s string) (r string) {
	return DelimiterCase(s, "-", "Coder", false)
}

func KebabCase(s string) (r string) {
	return DelimiterCase(s, "-", "lower", false)
}

func KebabForceCase(s string) (r string) {
	return DelimiterCase(s, "-", "Force", false)
}

func KebabTitleCase(s string) (r string) {
	return DelimiterCase(s, "-", "Title", false)
}

func KebabSmartCase(s string) (r string) {
	return DelimiterCase(s, "-", "Smart", false)
}

func KebabUpperCase(s string) (r string) {
	return DelimiterCase(s, "-", "UPPER", false)
}

func NoDelimiterCamelCase(s string) (r string) {
	return CamelCase(s)
}

func NoDelimiterCamelCoderCase(s string) (r string) {
	return DelimiterCase(s, "", "Coder", true)
}

func NoDelimiterLowerCase(s string) (r string) {
	return DelimiterCase(s, "", "lower", true)
}

func NoDelimiterPascalCase(s string) (r string) {
	return PascalCase(s)
}

func NoDelimiterUpperCase(s string) (r string) {
	return DelimiterCase(s, "", "UPPER", false)
}

func PascalCase(s string) (r string) {
	return DelimiterCase(s, "", "Force", false)
}

func PascalCoderCase(s string) (r string) {
	return DelimiterCase(s, "", "Coder", false)
}

func PreprocessString(s string) (r string) {
	p := ' '

	for _, c := range s {
		r += Separatem(p, c)
		p = c
	}

	r = strings.TrimSpace(r)

	return spaceRE.ReplaceAllString(r, " ")
}

/*
MEETS CRITERIA to continue with Diagnostic Evaluation--Sufficient Emotional/Behavioral Needs have been Identified Above
*/

// Separatem takes p (previous), and c (current) and compares them to determine how to separate them, if necessary.
func Separatem(p, c rune) (r string) {
	// If p and c are different cases return them concatenated with a space between
	// If c is a dot, hyphen, underscore, slash, colon, or semicolon then convert it to a space
	// NOTE: Consider scenarios with other delimiters and where to handle them. Probably not here.
	switch c {
	case '.', '-', '_', '/', ':', ';', '=', '+', '*':
		// Replace with space
		// fmt.Printf("%q is a delimiter\n", c)
		r = " "
	case 0x27, ',', '(', ')', '[', ']', '{', '}', '<', '>':
		// Drop
		r = ""
	case '&':
		// Symbol to word
		r = "and"
	default:
		if UseCodeBreak && unicode.IsLetter(p) && unicode.IsLetter(c) {
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
	return DelimiterCase(s, "_", "lower", false)
}

func SnakeCoderCase(s string) (r string) {
	return DelimiterCase(s, "_", "Coder", false)
}

func SnakeForceCase(s string) (r string) {
	return DelimiterCase(s, "_", "Force", false)
}

func SnakeTitleCase(s string) (r string) {
	return DelimiterCase(s, "_", "Title", false)
}

func SnakeSmartCase(s string) (r string) {
	return DelimiterCase(s, "_", "Smart", false)
}

func SnakeUpperCase(s string) (r string) {
	return DelimiterCase(s, "_", "UPPER", false)
}

func SpaceCoderCase(s string) (r string) {
	return DelimiterCase(s, " ", "Coder", false)
}

func SpaceForceCase(s string) (r string) {
	return DelimiterCase(s, " ", "Force", false)
}

func SpaceLowerCase(s string) (r string) {
	return DelimiterCase(s, " ", "lower", false)
}

func SpaceTitleCase(s string) (r string) {
	return DelimiterCase(s, " ", "Title", false)
}

func SpaceSmartCase(s string) (r string) {
	return DelimiterCase(s, " ", "Smart", false)
}

func SpaceUpperCase(s string) (r string) {
	r = PreprocessString(s)
	return strings.ToUpper(r)
}
