package prm

// Default stack what will be used by ctx
var dStack = []SkipFunc{
	SkipIfFirst,
	SkipIfPrev,
	SkipIfAbbrev,
	SkipIfApos,
}

// SetStack are allowing to set
// a new runes skipping stack
func SetStack(stack []SkipFunc) {
	dStack = stack
}

// SkipFunc type for runes skipping checks.
// If function is returing a true value then
// the letter (with index of j) will be skiped
type SkipFunc func(i, j int, lttrs []rune, sep rune) bool

// SkipIfFirst checks if a letter is first
// it will remove itself from the stack if a passed
// letters index isn't zero
func SkipIfFirst(i, j int, lttrs []rune, sep rune) (skip bool) {
	if i == 0 {
		skip = true
		return
	}

	return
}

// SkipIfPrev checks if a previous letter is a separator
func SkipIfPrev(i, j int, lttrs []rune, sep rune) (skip bool) {
	if lttrs[i-1] == sep {
		skip = true
	}

	return
}

// SkipIfAbbrev checks if a previous letter is an uppercase ASCII
// and next one is an uppercase rune too
func SkipIfAbbrev(i, j int, lttrs []rune, sep rune) (skip bool) {
	if isUpper(lttrs[i-1]) && (len(lttrs)-1 > j && isUpper(lttrs[j+1])) {
		skip = true
	}

	return
}

// SkipIfApos checks if a previous rune was a letter and current one is an apostrophe
func SkipIfApos(i, j int, lttrs []rune, sep rune) (skip bool) {
	if !notLoN(lttrs[i-1]) && lttrs[j] == '\'' {
		skip = true
	}

	return
}
