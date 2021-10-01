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

// FindAllString is the 'All' version of FindString; it returns a slice of all
// successive matches of the expression, as defined by the 'All' description
// in the package comment.
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
// A return value of nil indicates no match.
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int {
	m, err := re.FindStringMatch(s)
	if err != nil {
		println(err.Error())
		return nil
	}

	var result [][]int
	for m != nil {
		result = append(result, []int{m.Group.Index, m.Group.Length})

		m, err = re.FindNextMatch(m)
		if err != nil {
			println(err.Error())
			return nil
		}
	}

	return result
}

// FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it
// returns a slice of all successive matches of the expression, as defined by
// the 'All' description in the package comment.
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
