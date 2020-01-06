package main

import (
	"fmt"
	"os"

	"github.com/runeimp/caseit"

	arg "github.com/alexflint/go-arg"
)

/*
 * CONSTANTS
 */
const (
	AppDesc         = "Processes input string(s) providing conversion to almost any code standard desired, including proper title casing."
	AppName         = "CaseIt"
	AppAbbreviation = "CI"
	AppVerMajor     = 1
	AppVerMinor     = 0
	AppVerPatch     = 0
	AppVerPre       = ""
	CLIName         = "caseit"
)

/*
 * DERIVED CONSTANTS AS VARIABLES
 */
var (
	AppVersion = fmt.Sprintf("%d.%d.%d%s", AppVerMajor, AppVerMinor, AppVerPatch, AppVerPre)
	AppMeta    = AppMetaData{
		CLI:      CLIName,
		Desc:     AppDesc,
		Label:    fmt.Sprintf("%s v%s", AppName, AppVersion),
		Name:     AppName,
		Version:  AppVersion,
		VerMajor: AppVerMajor,
		VerMinor: AppVerMinor,
		VerPatch: AppVerPatch,
	}
	falsePointer = &[]bool{false}[0]
	truePointer  = &[]bool{true}[0]
)

/*
 * TYPES
 */
type (
	// AppMetaData defines meta-data about an application
	AppMetaData struct {
		CLI      string
		Desc     string
		Label    string
		Name     string
		VerMajor uint
		VerMinor uint
		VerPatch uint
		VerPre   string
		Version  string
	}

	AllCMD struct {
		Input  []string `arg:"positional" help:"The input string(s) to process"`
		All    bool     `arg:"-a" help:"Use all options. This is the default."`
		Camel  bool     `arg:"-C" help:"camelCase - Commonly used in Perl, JavaScript, and many other languages for variables and functions. Also for private properties in Go."`
		Force  bool     `arg:"-f" help:"Forced Title Case - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower  bool     `arg:"-l" help:"lower case - Commonly used for most of a sentence."`
		Title  bool     `arg:"-t" help:"Title Case - Words that are all caps are left all caps."`
		Coder  bool     `arg:"-c" help:"Coder Title Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Smart  bool     `arg:"-s" help:"Smart Title Case - Uses code based on Python Titlecase which is smart about small words, etc."`
		Pascal bool     `arg:"-p" help:"PascalCase - Words in all caps will be forced to initial letter upper and the rest lower. Common in Pascal and for public properties in Go."`
		Upper  bool     `arg:"-u" help:"UPPER CASE - Commonly used for movie titles."`
	}

	CamelCMD struct {
		Input  []string `arg:"positional" help:"The input string(s) to process"`
		All    bool     `arg:"-a" help:"Use all options"`
		Force  bool     `arg:"-f" help:"Forced TitleNodilCase - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower  bool     `arg:"-l" help:"titlenodilcase - Commonly used for environment variables."`
		Camel  *bool    `arg:"-C" help:"camelCase - Commonly used in Perl, JavaScript, and many other languages for variables and functions. Also for private properties in Go. This is the default."`
		Title  bool     `arg:"-t" help:"TitleNodilCase - Words that are all caps are left all caps."`
		Coder  bool     `arg:"-c" help:"Coder TitleNodilCase - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Pascal bool     `arg:"-p" help:"PascalCase - Words in all caps will be forced to initial letter upper and the rest lower. Common in Pascal and for public properties in Go."`
		Smart  bool     `arg:"-s" help:"Smart TitleNodilCase - Uses code based on Python Titlecase which is smart about small words, etc."`
		Upper  bool     `arg:"-u" help:"TITLENODILCASE - Commonly used for environment variables."`
	}

	DelimitedCMD struct {
		Input     []string `arg:"positional" help:"The input string(s) to process"`
		Delimiter string   `arg:"-d" default:"*" help:"The delimiter to use between words`
		All       bool     `arg:"-a" help:"Use all options"`
		Force     bool     `arg:"-f" help:"Forced Title*Kebab*Case - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower     bool     `arg:"-l" help:"lower*kebab*case - Commonly used for CSS class names and file names on POSIX systems. This is the default."`
		Title     bool     `arg:"-t" help:"Title*Kebab*Case - Words that are all caps are left all caps."`
		Coder     bool     `arg:"-c" help:"Coder Title*Kebab*Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Smart     bool     `arg:"-s" help:"Smart Title*Kebab*Case - Uses code based on Python Titlecase which is smart about small words, etc."`
		Upper     bool     `arg:"-u" help:"UPPER*KEBAB*CASE - Commonly used for important file names on POSIX systems?"`
	}

	DotCMD struct {
		Input []string `arg:"positional" help:"The input string(s) to process"`
		All   bool     `arg:"-a" help:"Use all options"`
		Force bool     `arg:"-f" help:"Forced Title.Dot.Case - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower bool     `arg:"-l" help:"lower.dot.case - Commonly to JavaScript and JSON dot notation."`
		Title bool     `arg:"-t" help:"Title.Dot.Case - Words that are all caps are left all caps."`
		Coder bool     `arg:"-c" help:"Coder Title.Dot.Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Smart bool     `arg:"-s" help:"Smart Title.Dot.Case - Uses code based on Python Titlecase which is smart about small words, etc."`
		Upper bool     `arg:"-u" help:"UPPER.DOT.CASE - Commonly used for ...?"`
	}

	KebabCMD struct {
		Input []string `arg:"positional" help:"The input string(s) to process"`
		All   bool     `arg:"-a" help:"Use all options"`
		Force bool     `arg:"-f" help:"Forced Title-Kebab-Case - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower bool     `arg:"-l" help:"lower-kebab-case - Commonly used for CSS class names and file names on POSIX systems. This is the default."`
		Title bool     `arg:"-t" help:"Title-Kebab-Case - Words that are all caps are left all caps."`
		Coder bool     `arg:"-c" help:"Coder Title-Kebab-Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Smart bool     `arg:"-s" help:"Smart Title-Kebab-Case - Uses code based on Python Titlecase which is smart about small words, etc."`
		Upper bool     `arg:"-u" help:"UPPER-KEBAB-CASE - Commonly used for important file names on POSIX systems?"`
	}

	SnakeCMD struct {
		Input []string `arg:"positional" help:"The input string(s) to process"`
		All   bool     `arg:"-a" help:"Use all options"`
		Force bool     `arg:"-f" help:"Forced Title_Snake_Case - Words in all caps will be forced to initial letter upper and the rest lower. Commonly used for database names."`
		Lower bool     `arg:"-l" help:"Python's traditional_snake_case. This is the default."`
		Title bool     `arg:"-t" help:"Title_Snake_Case - Words that are all caps are left all caps. Commonly used for database names."`
		Coder bool     `arg:"-c" help:"Coder Title_Snake_Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased. Commonly used for database names."`
		Smart bool     `arg:"-s" help:"Smart Title_Snake_Case - Uses code based on Python Titlecase which is smart about small words, etc. Commonly used for database names."`
		Upper bool     `arg:"-u" help:"UPPER_SNAKE_CASE or CONSTANTS_CASE in many languages."`
	}

	SpaceCMD struct {
		Input []string `arg:"positional" help:"The input string(s) to process"`
		All   bool     `arg:"-a" help:"Use all options"`
		Force bool     `arg:"-f" help:"Forced Title Case - Words in all caps will be forced to initial letter upper and the rest lower."`
		Lower bool     `arg:"-l" help:"lower case - Commonly used for most of a sentence."`
		Title bool     `arg:"-t" help:"Title Case - Words that are all caps are left all caps."`
		Coder bool     `arg:"-c" help:"Coder Title Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased."`
		Smart bool     `arg:"-s" help:"Smart Title Case - Uses code based on Python Titlecase which is smart about small words, etc. This is the default."`
		Upper bool     `arg:"-u" help:"UPPER CASE - Commonly used for movie titles."`
	}

	appArgs struct {
		AllCMD *AllCMD       `arg:"subcommand:all" help:"All commands used"`
		Camel  *CamelCMD     `arg:"subcommand:camel" help:"Words are not separated by a delimiter"`
		Char   *DelimitedCMD `arg:"subcommand:char" help:"Character specified separated words"`
		Dot    *DotCMD       `arg:"subcommand:dot" help:"Dot separated words"`
		Kebab  *KebabCMD     `arg:"subcommand:kebab" help:"Hyphen separated words"`
		Snake  *SnakeCMD     `arg:"subcommand:snake" help:"Underscore separated words"`
		Word   *SpaceCMD     `arg:"subcommand:word" help:"Space separated words"`
	}
)

func (appArgs) Description() string {
	return AppMeta.Desc
}

func (appArgs) Version() string {
	return AppMeta.Label
}

/*
 * LOCAL FUNCTIONS
 */

func rangeAllCase(subCMD *AllCMD, inputList []string, delimiter string, firstWordLower bool) {
	// fmt.Printf("rangeAllCase() | subCMD = %#v\n", subCMD)
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.DelimiterCase(input, delimiter, "lower", firstWordLower))
		}
		if delimiter == "" && (subCMD.All || subCMD.Camel) {
			fmt.Printf(" camel: %s\n", caseit.CamelCase(input))
			// fmt.Printf(" Camel: %s\n", caseit.DelimiterCase(input, delimiter, "Force", true))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.DelimiterCase(input, delimiter, "Force", firstWordLower))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.DelimiterCase(input, delimiter, "Title", firstWordLower))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.DelimiterCase(input, delimiter, "Coder", firstWordLower))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.DelimiterCase(input, delimiter, "Smart", firstWordLower))
		}
		if delimiter == "" && (subCMD.All || subCMD.Pascal) {
			fmt.Printf("Pascal: %s\n", caseit.PascalCase(input))
			// fmt.Printf("Pascal: %s\n", caseit.DelimiterCase(input, delimiter, "Smart", firstWordLower))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.DelimiterCase(input, delimiter, "UPPER", firstWordLower))
		}
		fmt.Println()
	}
}

func rangeCamelCase(subCMD *CamelCMD, inputList []string) {
	fmt.Println("    CamelCase:")
	delimiter := ""
	firstWordLower := false
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.DelimiterCase(input, delimiter, "lower", firstWordLower))
		}
		// fmt.Printf("rangeCamelCase() | subCMD.Camel = %#v | truePointer = %#v\n", subCMD.Camel, truePointer)
		if subCMD.All || subCMD.Camel != nil && subCMD.Camel == truePointer {
			fmt.Printf(" camel: %s\n", caseit.CamelCase(input))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.DelimiterCase(input, delimiter, "Force", firstWordLower))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.DelimiterCase(input, delimiter, "Title", firstWordLower))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.DelimiterCase(input, delimiter, "Coder", firstWordLower))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.DelimiterCase(input, delimiter, "Smart", firstWordLower))
		}
		if subCMD.All || subCMD.Pascal {
			fmt.Printf("Pascal: %s\n", caseit.PascalCase(input))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.DelimiterCase(input, delimiter, "UPPER", firstWordLower))
		}
		fmt.Println()
	}
}

func rangeDelimitedCase(subCMD *DelimitedCMD, inputList []string, delimiter string, firstWordLower bool) {
	// fmt.Printf("rangeAllCase() | subCMD = %#v\n", subCMD)
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.DelimiterCase(input, delimiter, "lower", firstWordLower))
		}
		// if delimiter == "" && (subCMD.All || subCMD.Camel) {
		// 	fmt.Printf(" camel: %s\n", caseit.CamelCase(input))
		// 	// fmt.Printf(" Camel: %s\n", caseit.DelimiterCase(input, delimiter, "Force", true))
		// }
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.DelimiterCase(input, delimiter, "Force", firstWordLower))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.DelimiterCase(input, delimiter, "Title", firstWordLower))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.DelimiterCase(input, delimiter, "Coder", firstWordLower))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.DelimiterCase(input, delimiter, "Smart", firstWordLower))
		}
		// if delimiter == "" && (subCMD.All || subCMD.Pascal) {
		// 	fmt.Printf("Pascal: %s\n", caseit.PascalCase(input))
		// 	// fmt.Printf("Pascal: %s\n", caseit.DelimiterCase(input, delimiter, "Smart", firstWordLower))
		// }
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.DelimiterCase(input, delimiter, "UPPER", firstWordLower))
		}
		fmt.Println()
	}
}

func rangeDotCase(subCMD *DotCMD, inputList []string) {
	fmt.Println("    dot.case:")
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.DotLowerCase(input))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.DotForceCase(input))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.DotTitleCase(input))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.DotCoderCase(input))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.DotSmartCase(input))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.DotUpperCase(input))
		}
		fmt.Println()
	}
}

func rangeKebabCase(subCMD *KebabCMD, inputList []string) {
	fmt.Println("    kebab-case:")
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.KebabCase(input))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.KebabForceCase(input))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.KebabTitleCase(input))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.KebabCoderCase(input))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.KebabSmartCase(input))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.KebabUpperCase(input))
		}
		fmt.Println()
	}
}

func rangeSnakeCase(subCMD *SnakeCMD, inputList []string) {
	fmt.Println("    snake_case:")
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.SnakeCase(input))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.SnakeForceCase(input))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.SnakeTitleCase(input))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.SnakeCoderCase(input))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.SnakeSmartCase(input))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.SnakeUpperCase(input))
		}
		fmt.Println()
	}
}

func rangeSpaceCase(subCMD *SpaceCMD, inputList []string) {
	fmt.Println("    Word Case:")
	for _, input := range inputList {
		if subCMD.All || subCMD.Lower {
			fmt.Printf(" lower: %s\n", caseit.SpaceLowerCase(input))
		}
		if subCMD.All || subCMD.Force {
			fmt.Printf(" Force: %s\n", caseit.SpaceForceCase(input))
		}
		if subCMD.All || subCMD.Title {
			fmt.Printf(" Title: %s\n", caseit.SpaceTitleCase(input))
		}
		if subCMD.All || subCMD.Coder {
			fmt.Printf(" Coder: %s\n", caseit.SpaceCoderCase(input))
		}
		if subCMD.All || subCMD.Smart {
			fmt.Printf(" Smart: %s\n", caseit.SpaceSmartCase(input))
		}
		if subCMD.All || subCMD.Upper {
			fmt.Printf(" UPPER: %s\n", caseit.SpaceUpperCase(input))
		}
		fmt.Println()
	}
}

func showSource(src []string) {
	if len(src) > 0 {
		fmt.Printf("Source: %q\n", src[0])
	}
	for i, input := range src {
		if i > 0 {
			fmt.Printf("        %q\n", input)
		}
	}
}

/*
 * THE MAIN ENTRY POINT
 */

func main() {
	args := appArgs{}
	p := arg.MustParse(&args)

	if len(os.Args) == 1 {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	/*
	 * Handling CLI Argument Defaults
	 */
	if args.AllCMD != nil {
		if args.AllCMD.All == false && args.AllCMD.Force == false && args.AllCMD.Lower == false &&
			args.AllCMD.Title == false && args.AllCMD.Coder == false && args.AllCMD.Pascal == false &&
			args.AllCMD.Smart == false && args.AllCMD.Upper == false {
			args.AllCMD.All = true
		}
	}
	if args.Camel != nil {
		// fmt.Printf("main() | args.Camel.Camel = %#v | truePointer = %#v\n", args.Camel.Camel, truePointer)
		// Since args.Camel.Camel is now *bool it will always either be nil or *true
		if args.Camel.Camel == nil {
			if args.Camel.All == false && args.Camel.Force == false && args.Camel.Lower == false &&
				args.Camel.Title == false && args.Camel.Coder == false && args.Camel.Pascal == false &&
				args.Camel.Smart == false && args.Camel.Upper == false {
				args.Camel.Camel = truePointer
			} else {
				args.Camel.Camel = falsePointer
			}
		} else {
			args.Camel.Camel = truePointer
		}
		// fmt.Printf("main() | args.Camel.Camel = %#v | truePointer = %#v\n", args.Camel.Camel, truePointer)
	}

	if args.Char != nil {
		if args.Char.All == false && args.Char.Force == false && args.Char.Lower == false && args.Char.Title == false &&
			args.Char.Coder == false && args.Char.Smart == false && args.Char.Upper == false {
			args.Char.Coder = true // NOTE: Maybe this shouldn't be defaulted?
		}
	}

	if args.Dot != nil {
		if args.Dot.All == false && args.Dot.Force == false && args.Dot.Lower == false && args.Dot.Title == false &&
			args.Dot.Coder == false && args.Dot.Smart == false && args.Dot.Upper == false {
			args.Dot.Coder = true // NOTE: Maybe this shouldn't be defaulted? Seems like camelCase is missing and should be the default.
		}
	}

	if args.Kebab != nil {
		if args.Kebab.All == false && args.Kebab.Force == false && args.Kebab.Lower == false && args.Kebab.Title == false &&
			args.Kebab.Coder == false && args.Kebab.Smart == false && args.Kebab.Upper == false {
			args.Kebab.Lower = true
		}
	}

	if args.Snake != nil {
		if args.Snake.All == false && args.Snake.Force == false && args.Snake.Lower == false && args.Snake.Title == false &&
			args.Snake.Coder == false && args.Snake.Smart == false && args.Snake.Upper == false {
			args.Snake.Lower = true
		}
	}

	if args.Word != nil {
		if args.Word.All == false && args.Word.Force == false && args.Word.Lower == false && args.Word.Title == false &&
			args.Word.Coder == false && args.Word.Smart == false && args.Word.Upper == false {
			args.Word.Coder = true // NOTE: Maybe this shouldn't be defaulted?
		}
	}

	/*
	 * Handling The CLI Arguments
	 */
	if args.AllCMD != nil {
		// fmt.Printf("main() | subCMD = %#v\n", args.AllCMD)
		showSource(args.AllCMD.Input)
	}

	if args.AllCMD != nil || args.Camel != nil {
		if args.AllCMD == nil {
			showSource(args.Camel.Input)
			rangeCamelCase(args.Camel, args.Camel.Input)
		} else {
			fmt.Println("    CamelCase:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, "", firstWordLower)
		}
	}

	if args.AllCMD != nil || args.Char != nil {
		if args.AllCMD == nil {
			showSource(args.Char.Input)
			firstWordLower := false
			rangeDelimitedCase(args.Char, args.Char.Input, args.Char.Delimiter, firstWordLower)
		} else {
			fmt.Println("    char*case:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, ".", firstWordLower)
		}
	}

	if args.AllCMD != nil || args.Dot != nil {
		if args.AllCMD == nil {
			showSource(args.Dot.Input)
			rangeDotCase(args.Dot, args.Dot.Input)
		} else {
			fmt.Println("    dot.case:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, ".", firstWordLower)
		}
	}

	if args.AllCMD != nil || args.Kebab != nil {
		if args.AllCMD == nil {
			showSource(args.Kebab.Input)
			rangeKebabCase(args.Kebab, args.Kebab.Input)
		} else {
			fmt.Println("    kebab-case:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, "-", firstWordLower)
		}
	}

	if args.AllCMD != nil || args.Snake != nil {
		if args.AllCMD == nil {
			showSource(args.Snake.Input)
			rangeSnakeCase(args.Snake, args.Snake.Input)
		} else {
			fmt.Println("    snake_case:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, "_", firstWordLower)
		}
	}

	if args.AllCMD != nil || args.Word != nil {
		if args.AllCMD == nil {
			showSource(args.Word.Input)
			rangeSpaceCase(args.Word, args.Word.Input)
		} else {
			fmt.Println("    Word Case:")
			firstWordLower := false
			rangeAllCase(args.AllCMD, args.AllCMD.Input, " ", firstWordLower)
		}
	}
}
