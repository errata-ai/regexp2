package regexp2

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
