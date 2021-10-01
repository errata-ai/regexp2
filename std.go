package regexp2

func CompileStd(re string) (*Regexp, error) {
	return Compile(re, Multiline)
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
