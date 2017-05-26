// Copyright 2017 romanyx. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package prm are allowing to generate a url string
from a ASCII string

For example:
	prm.Parameterize("Computer world", '-') // "computer-world"

You can set your own skipping stack. By default Parameterize method
wouldn't replace dots in abbreviations as separator.

	prm.Parameterize("C.I.A. and K.G.B.", '-') // "cia-and-kgb"

But this behaviour can be changed by using SetStack method

	stack := []prm.SkipFunc{
		prm.SkipIfFirst,
		prm.SkipIfPrev,
		prm.SkipIfApos,
	}
	prm.SetStack(stack)

Now Parameterize method will behave differently

	prm.Parameterize("C.I.A. and K.G.B.", '-') // "c-i-a-and-k-g-b"

Default stack is: SkipIfFirst, SkipIfPrev, SkipIfAbbrev, SkipIfApos

It will skip if separator is a first letter, if a previous letter
was replaced as separator, if a previous letter and next is a uppercase
and if a letter is an apostrophe

*/
package prm
