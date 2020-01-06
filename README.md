CaseIt
======

Case conversion library and command line tool written in Go

CLI Tool v1.0.0  
Library v1.0.0


Case Conversions Implemented
----------------------------

### Word Casing

* lower case -- All words in statement are lower cased.
* Title Case (simple) -- All words have the first character upper cased. All other letters are left as they were initially.
* Title Case (forced) -- All letters in a word after the first are lower cased.
* Title Case (coder) -- Like simple title case but also checks if a word is a common initialism such as gRPC, OAuth, OpenID, etc. and enforces the initialism.
* Title Case (smart) -- Like simple title case but also lower cases small words such as "a", "the", etc. Uses my [titlecase][] library which is based on [Python Titlecase][].
* UPPER CASE -- All words upper cased.

### Word Delimiters

Each of the word casing options are then paired with a word separation option

* CamelCase AKA No Delimiter
	* `lowercase`
	* `camelCase` -- Initial word is always lower cased.
	* `TitleCase`
	* `ForcedTitleCase`
	* `CoderTitleCase`
	* `SmartTitleCase`
	* `PascalCase` -- The same a `ForcedTitleCase`
	* `UPPERCASE`
* char\*case - Any delimiter for those who simply need _more options_. `*` is the default delimiter if you don't specify one with `-d` or `--delimiter`.
	* `dot*lower*case`
	* `Dot*Title*Case`
	* `Dot*Forced*Title*Case`
	* `Dot*Coder*Title*Case`
	* `Dot*Smart*Title*Case`
	* `DOT*UPPER*CASE`
* dot.case
	* `dot.lower.case`
	* `Dot.Title.Case`
	* `Dot.Forced.Title.Case`
	* `Dot.Coder.Title.Case`
	* `Dot.Smart.Title.Case`
	* `DOT.UPPER.CASE`
* kebab-case
	* `kebab-lower-case`
	* `Kebab-Title-Case`
	* `Kebab-Forced-Title-Case`
	* `Kebab-Coder-Title-Case`
	* `Kebab-Smart-Title-Case`
	* `KEBAB-UPPER-CASE`
* `snake_case`
	* `snake_case`
	* `Snake_Title_Case`
	* `Snake_Forced_Title_Case`
	* `Snake_Coder_Title_Case`
	* `Snake_Smart_Title_Case`
	* `SNAKE_UPPER_CASE`
* Word Case AKA Space Case
	* `separate words lower case`
	* `Separate Words Title Case`
	* `Separate Words Forced Title Case`
	* `Dot.Coder.Title.Case`
	* `Separate Words Smart Title Case` (words like "a", "the", etc. are not title cased)
	* `SEPARATE WORDS UPPER CASE`


Features
--------

* Case conversion
* Word delimitation
* No external libraries. OK, that's not entirely true. But the one external library is my own [titlecase][] package.  :angel:
* I wanted a tool that could also do `Snake_Title_Case`. I couldn't find one so I created this tool.


Example
-------

### Help

```text
$ caseit
Processes input string(s) providing conversion to almost any code standard desired, including proper title casing.
CaseIt v1.0.0
Usage: caseit

Options:
  --help, -h             display this help and exit
  --version              display version and exit

Commands:
  all                    All commands used
  camel                  Words are not separated by a delimiter
  char                   Character specified separated words
  dot                    Dot separated words
  kebab                  Hyphen separated words
  snake                  Underscore separated words
  word                   Space separated words
```


#### Subcommand Help

```text
$ caseit snake -h
Processes input string(s) providing conversion to almost any code standard desired, including proper title casing.
CaseIt v1.0.0
Usage: caseit snake [--all] [--force] [--lower] [--title] [--coder] [--smart] [--upper] [INPUT [INPUT ...]]

Positional arguments:
  INPUT                  The input string(s) to process

Options:
  --all, -a              Use all options
  --force, -f            Forced Title_Snake_Case - Words in all caps will be forced to initial letter upper and the rest lower. Commonly used for database names.
  --lower, -l            Python's traditional_snake_case. This is the default.
  --title, -t            Title_Snake_Case - Words that are all caps are left all caps. Commonly used for database names.
  --coder, -c            Coder Title_Snake_Case - Words that are common initialisms are properly cased for the initialism. Otherwise force title cased. Commonly used for database names.
  --smart, -s            Smart Title_Snake_Case - Uses code based on Python Titlecase which is smart about small words, etc. Commonly used for database names.
  --upper, -u            UPPER_SNAKE_CASE or CONSTANTS_CASE in many languages.
  --help, -h             display this help and exit
  --version              display version and exit
```


### snake_case

#### Default Case Only (lower for snake_case)

```bash
$ caseit snake "This OAuth test" " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
Source: "This OAuth test"
        " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
    snake_case:
 lower: this_oauth_test

 lower: all_caps_to_go_thru_1_6_of_the_1_2_with_the_correct_cia_id_over_grpc

```

#### Force Case, Coder Case, and Smart Case

```bash
$ caseit snake -f -c -s "This OAuth test" " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
Source: "This OAuth test"
        " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
    snake_case:
 Force: This_Oauth_Test
 Coder: This_OAuth_Test
 Smart: This_OAuth_Test

 Force: All_Caps_To_Go_Thru_1_6_Of_The_1_2_With_The_Correct_Cia_Id_Over_Grpc
 Coder: All_Caps_To_Go_Thru_1_6_Of_The_1_2_With_The_Correct_CIA_ID_Over_gRPC
 Smart: ALL_CAPS_to_Go_Thru_1_6_of_the_1_2_with_the_Correct_CIA_ID_Over_GRPC

```

### The Kitchen Sink

#### All cases and all styles

```bash
$ caseit all "This OAuth test" " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
Source: "This OAuth test"
        " ALL CAPS to go thru 1--6 of the 1/2 with the correct CIA ID over GRPC"
    CamelCase:
 lower: thisoauthtest
 camel: thisOauthTest
 Force: ThisOauthTest
 Title: ThisOAuthTest
 Coder: ThisOAuthTest
 Smart: ThisOAuthTest
Pascal: ThisOauthTest
 UPPER: THISOAUTHTEST

 lower: allcapstogothru16ofthe12withthecorrectciaidovergrpc
 camel: allCapsToGoThru16OfThe12WithTheCorrectCiaIdOverGrpc
 Force: AllCapsToGoThru16OfThe12WithTheCorrectCiaIdOverGrpc
 Title: ALLCAPSToGoThru16OfThe12WithTheCorrectCIAIDOverGRPC
 Coder: AllCapsToGoThru16OfThe12WithTheCorrectCIAIDOvergRPC
 Smart: ALLCAPStoGoThru16ofthe12withtheCorrectCIAIDOverGRPC
Pascal: AllCapsToGoThru16OfThe12WithTheCorrectCiaIdOverGrpc
 UPPER: ALLCAPSTOGOTHRU16OFTHE12WITHTHECORRECTCIAIDOVERGRPC

    char*case:
 lower: this.oauth.test
 Force: This.Oauth.Test
 Title: This.OAuth.Test
 Coder: This.OAuth.Test
 Smart: This.OAuth.Test
 UPPER: THIS.OAUTH.TEST

 lower: all.caps.to.go.thru.1.6.of.the.1.2.with.the.correct.cia.id.over.grpc
 Force: All.Caps.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.Cia.Id.Over.Grpc
 Title: ALL.CAPS.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.CIA.ID.Over.GRPC
 Coder: All.Caps.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.CIA.ID.Over.gRPC
 Smart: ALL.CAPS.to.Go.Thru.1.6.of.the.1.2.with.the.Correct.CIA.ID.Over.GRPC
 UPPER: ALL.CAPS.TO.GO.THRU.1.6.OF.THE.1.2.WITH.THE.CORRECT.CIA.ID.OVER.GRPC

    dot.case:
 lower: this.oauth.test
 Force: This.Oauth.Test
 Title: This.OAuth.Test
 Coder: This.OAuth.Test
 Smart: This.OAuth.Test
 UPPER: THIS.OAUTH.TEST

 lower: all.caps.to.go.thru.1.6.of.the.1.2.with.the.correct.cia.id.over.grpc
 Force: All.Caps.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.Cia.Id.Over.Grpc
 Title: ALL.CAPS.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.CIA.ID.Over.GRPC
 Coder: All.Caps.To.Go.Thru.1.6.Of.The.1.2.With.The.Correct.CIA.ID.Over.gRPC
 Smart: ALL.CAPS.to.Go.Thru.1.6.of.the.1.2.with.the.Correct.CIA.ID.Over.GRPC
 UPPER: ALL.CAPS.TO.GO.THRU.1.6.OF.THE.1.2.WITH.THE.CORRECT.CIA.ID.OVER.GRPC

    kebab-case:
 lower: this-oauth-test
 Force: This-Oauth-Test
 Title: This-OAuth-Test
 Coder: This-OAuth-Test
 Smart: This-OAuth-Test
 UPPER: THIS-OAUTH-TEST

 lower: all-caps-to-go-thru-1-6-of-the-1-2-with-the-correct-cia-id-over-grpc
 Force: All-Caps-To-Go-Thru-1-6-Of-The-1-2-With-The-Correct-Cia-Id-Over-Grpc
 Title: ALL-CAPS-To-Go-Thru-1-6-Of-The-1-2-With-The-Correct-CIA-ID-Over-GRPC
 Coder: All-Caps-To-Go-Thru-1-6-Of-The-1-2-With-The-Correct-CIA-ID-Over-gRPC
 Smart: ALL-CAPS-to-Go-Thru-1-6-of-the-1-2-with-the-Correct-CIA-ID-Over-GRPC
 UPPER: ALL-CAPS-TO-GO-THRU-1-6-OF-THE-1-2-WITH-THE-CORRECT-CIA-ID-OVER-GRPC

    snake_case:
 lower: this_oauth_test
 Force: This_Oauth_Test
 Title: This_OAuth_Test
 Coder: This_OAuth_Test
 Smart: This_OAuth_Test
 UPPER: THIS_OAUTH_TEST

 lower: all_caps_to_go_thru_1_6_of_the_1_2_with_the_correct_cia_id_over_grpc
 Force: All_Caps_To_Go_Thru_1_6_Of_The_1_2_With_The_Correct_Cia_Id_Over_Grpc
 Title: ALL_CAPS_To_Go_Thru_1_6_Of_The_1_2_With_The_Correct_CIA_ID_Over_GRPC
 Coder: All_Caps_To_Go_Thru_1_6_Of_The_1_2_With_The_Correct_CIA_ID_Over_gRPC
 Smart: ALL_CAPS_to_Go_Thru_1_6_of_the_1_2_with_the_Correct_CIA_ID_Over_GRPC
 UPPER: ALL_CAPS_TO_GO_THRU_1_6_OF_THE_1_2_WITH_THE_CORRECT_CIA_ID_OVER_GRPC

    Word Case:
 lower: this oauth test
 Force: This Oauth Test
 Title: This OAuth Test
 Coder: This OAuth Test
 Smart: This OAuth Test
 UPPER: THIS OAUTH TEST

 lower: all caps to go thru 1 6 of the 1 2 with the correct cia id over grpc
 Force: All Caps To Go Thru 1 6 Of The 1 2 With The Correct Cia Id Over Grpc
 Title: ALL CAPS To Go Thru 1 6 Of The 1 2 With The Correct CIA ID Over GRPC
 Coder: All Caps To Go Thru 1 6 Of The 1 2 With The Correct CIA ID Over gRPC
 Smart: ALL CAPS to Go Thru 1 6 of the 1 2 with the Correct CIA ID Over GRPC
 UPPER: ALL CAPS TO GO THRU 1 6 OF THE 1 2 WITH THE CORRECT CIA ID OVER GRPC

```



ToDo:
-----

* [x] Add all the casing options
* [x] Add more options to `caseit` command line tool
* [x] Add custom delimiter support for CLI input
* [x] Add custom delimiter support for library output
* [ ] Add unit tests
* [ ] <s>Update</s> Create the docs
* [x] Update `caseit.Separatem()` to handle things like "ID or other ALL CAPS casing" to maintain the caps when possible
* [ ] Setup cross-compiling for releases


References
----------

My main reference in starting this was [Properly Capitalizing a Title][] and a quick review of the other similar products that already existed.


Similar Products
----------------

* [strcase by Ian Coleman][] -- Very close to what I did with CaseIt. There are a few features in both that are present in the other.
* [go-strcase by Adrian Stoewer][] -- This one is fine if you just need `kebab-case`, `snake_case`, `camelCase`, and `PascalCase`.
* [snaker by Serenize UG] -- This one only handles `snake_case`, `camelCase`, and `PascalCase`. But also is aware of common initialisms and words with odd capitalization usage such as ID and OAuth and inspired my usage of that concept in this app.



[go-strcase by Adrian Stoewer]: https://github.com/stoewer/go-strcase
[Properly Capitalizing a Title]: https://golangcookbook.com/chapters/strings/title/
[Python Titlecase]: https://github.com/ppannuto/python-titlecase
[snaker by Serenize UG]: https://github.com/serenize/snaker
[strcase by Ian Coleman]: https://github.com/iancoleman/strcase
[titlecase]: https://github.com/runeimp/titlecase

