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


References
----------

My main reference in starting this was [Properly Capitalizing a Title][] and a quick review of the other similar products that already existed.


Similar Products
----------------

* [strcase by Ian Coleman][] -- Very close to what I did with CaseIt. There are a few features in both that are present in the other.
* [go-strcase by Adrian Stoewer][] -- This one if fine if you just need `kebab-case`, `snake_case`, `camelCase`, and `PascalCase`.
* [snaker by Serenize UG] -- This one only handles `snake_case`, `camelCase`, and `PascalCase`. But also is aware of common initialisms and words with odd capitalization usage such as ID and OAuth.



[go-strcase by Adrian Stoewer]: https://github.com/stoewer/go-strcase
[Properly Capitalizing a Title]: https://golangcookbook.com/chapters/strings/title/
[snaker by Serenize UG]: https://github.com/serenize/snaker
[strcase by Ian Coleman]: https://github.com/iancoleman/strcase
