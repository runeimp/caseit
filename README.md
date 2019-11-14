CaseIt
======

Case conversion library and command line tool written in Go


Case Conversions Implemented
----------------------------

* NoDelimiter
	* [x] `lowercase`
	* [x] `camelCase`
	* [x] `PascalCase`
	* [x] `UPPERCASE`
* dot.case
	* [x] `dot.lower.case`
	* [x] `Dot.Title.Case`
	* [x] `DOT.UPPER.CASE`
* kebab-case
	* [x] `kebab-lower-case`
	* [x] `Kebab-Title-Case`
	* [x] `KEBAB-UPPER-CASE`
* Separate Words
	* [x] `separate words lower case`
	* [x] `Separate Words Title Case`
	* [ ] `Separate Words Smart Title Case` (words like "a", "the", etc. are not title cased)
	* [x] `SEPARATE WORDS UPPER CASE`
* `snake_case`
	* [x] `snake_case`
	* [x] `Snake_Title_Case`
	* [x] `SNAKE_UPPER_CASE`


Features
--------

* Case conversion
* No external libraries. This may change with a later version but in general striving to be as pure as reasonable.
* The only tool I've seen that can do `Snake_Title_Case` which is why I created this.

ToDo:
-----

* [ ] Add all the casing options
* [ ] Add more options to `caseit` command line tool
* [ ] Add custom delimiter support for input
* [ ] Add custom delimiter support for output
* [ ] Add unit tests
* [ ] <s>Update</s> Create the docs
* [ ] Update `caseit.Separatem()` to handle things like "ID or other ALL CAPS casing" to maintain the caps when possible
* [ ] Setup cross-compiling for releases

