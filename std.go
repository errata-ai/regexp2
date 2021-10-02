package regexp2

func CompileStd(s string) (*Regexp, error) {
	return Compile(s, Multiline)
}

func MustCompileStd(s string) *Regexp {
	re, err := CompileStd(s)
	if err != nil {
		panic(err)
	}
	return re
}

func (re *Regexp) MatchStringStd(s string) bool {
	match, err := re.MatchString(s)
	if err != nil {
		panic(err)
	}
	return match
}

// FindAllString is the 'All' version of FindString; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
//
// A return value of nil indicates no match.
func (re *Regexp) FindAllString(s string, n int) []string {
	var result []string

	m, err := re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	for m != nil {
		result = append(result, m.Group.String())

		m, err = re.FindNextMatch(m)
		if err != nil {
			panic(err)
		}
	}

	if n > -1 {
		result = result[:n]
	}

	return result
}

func (re *Regexp) FindAllStringMatches(s string) []*Match {
	var matches []*Match

	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m)
		m, _ = re.FindNextMatch(m)
	}

	return matches
}

// FindAllStringIndex is the 'All' version of FindStringIndex; it returns a
// slice of all successive matches of the expression, as defined by the 'All'
// description in the package comment.
//
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int {
	var result [][]int

	m, err := re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	for m != nil {
		result = append(result,
			[]int{m.Group.Index, m.Group.Index + m.Group.Length})

		m, err = re.FindNextMatch(m)
		if err != nil {
			panic(err)
		}
	}

	return result
}

// FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it
// returns a slice of all successive matches of the expression, as defined by
// the 'All' description in the package comment.
//
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string {
	var result [][]string

	m, err := re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	for m != nil {
		m.populateOtherGroups()

		subs := make([]string, 0, len(m.otherGroups)+1)
		subs = append(subs, m.Group.String())

		for i := 0; i < len(m.otherGroups); i++ {
			subs = append(subs, (&m.otherGroups[i]).String())
		}
		result = append(result, subs)

		m, err = re.FindNextMatch(m)
		if err != nil {
			panic(err)
		}
	}

	return result
}

// FindAllStringSubmatchIndex is the 'All' version of
// FindStringSubmatchIndex; it returns a slice of all successive matches of
// the expression, as defined by the 'All' description in the package
// comment.
//
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int {
	var result [][]int

	m, err := re.FindStringMatch(s)
	if err != nil {
		panic(err)
	}

	for m != nil {
		subs := []int{m.Group.Index, m.Group.Index + m.Group.Length}

		m.populateOtherGroups()
		for i := 0; i < len(m.otherGroups); i++ {
			g := m.otherGroups[i]
			if g.Index+g.Length == 0 {
				g.Index = -1
			}
			subs = append(subs, g.Index)
			subs = append(subs, g.Index+g.Length)
		}
		result = append(result, subs)

		m, err = re.FindNextMatch(m)
		if err != nil {
			panic(err)
		}
	}

	return result
}

// SubexpNames returns the names of the parenthesized subexpressions
// in this StdRegexp. The name for the first sub-expression is names[1],
// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
// Since the StdRegexp as a whole cannot be named, names[0] is always
// the empty string. The slice should not be modified.
func (re *Regexp) SubexpNames() []string {
	results := []string{}
	for i, s := range re.capslist {
		if i == 0 {
			results = append(results, "")
		} else {
			results = append(results, s)
		}
	}
	return results
}
