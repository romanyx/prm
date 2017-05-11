package prm

// ctx struct for parameterizing
type ctx struct {
	lttrs []rune
	sep   rune
	stack []SkipFunc
}

// Parameterize iterates each letter.
// If the letter isn't an ASCII or a number
// it will iterate through []SkipFunc stack passing
// into it previous passed index, current iteration
// index, slice of letters and a separator rune.
// If some of SkipFunc's will return true, then
// []SkipFunc iteration will break and letter
// will be skiped. All non ASCII letters and
// numbers will be replaced as separators
func (ctx *ctx) parameterize() {
	var i int

	for j := 0; j < len(ctx.lttrs); j++ {
		if notLoN(ctx.lttrs[j]) {
			var skip bool

			for _, skipFunc := range ctx.stack {
				if skip = skipFunc(i, j, ctx.lttrs, ctx.sep); skip {
					break
				}
			}

			if skip {
				continue
			}

			ctx.lttrs[j] = ctx.sep
		}

		ctx.lttrs[i] = ctx.lttrs[j]
		i++
	}

	if ctx.lttrs[i-1] == ctx.sep {
		i--
	}

	ctx.lttrs = ctx.lttrs[:i]
}

// Parameterize is parameterizing the string
// using the passed separator, downcasing letters
// and returning string
func Parameterize(s string, sep rune) string {
	ctx := &ctx{}
	ctx.lttrs = make([]rune, 0, len(s))
	ctx.sep = sep
	ctx.stack = dStack

	for _, r := range s {
		ctx.lttrs = append(ctx.lttrs, r)
	}

	ctx.parameterize()

	for i, r := range ctx.lttrs {
		ctx.lttrs[i] = downcase(r)
	}

	return string(ctx.lttrs)
}

// Return true if the rune is a digit or a ASCII letter
func notLoN(r rune) bool {
	if r <= 'z' && r >= 'A' {
		return false
	}

	if r <= '9' && r >= '0' {
		return false
	}

	return true
}

// Return true if the rune is an uppercase ASCII
func isUpper(r rune) bool {
	return r <= 'Z' && r >= 'A'
}

// Downcase ASCII letter
func downcase(r rune) rune {
	if isUpper(r) {
		return r + 32
	}

	return r
}
