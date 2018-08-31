package woma

type rPhrase struct {
	lenght int
	words  [][]rune
}

func (p Phrase) asRunes() *rPhrase {
	if p == nil {
		return nil
	}
	rphrase := rPhrase{
		words: make([][]rune, len(p))}
	for i := range p {
		chars := []rune(p[i])
		rphrase.words[i] = chars
		rphrase.lenght += len(chars)
	}
	return &rphrase
}

func similarityQuotient(a, b *rPhrase) float32 {

	u := a.words
	v := b.words
	if len(u) < len(v) {
		u, v = v, u
	}
	n := len(u)

	// TODO: memory allocation can be optimized; but this
	// is a first version
	edits := make([][]*int, n)
	for i := range u {
		edits[i] = make([]*int, n)
		for j := range v {
			r := levenshtein(u[i], v[j])
			edits[i][j] = &r
		}
	}

	minEditCount := max(a.lenght, b.lenght)
	for offset := 0; offset < n; offset++ {
		editCount := 0
		for i := 0; i < n; i++ {
			j := (i + offset) % n
			val := edits[i][j]
			if val == nil {
				editCount += len(u[i])
			} else {
				editCount += *val
			}
		}
		minEditCount = min(minEditCount, editCount)
	}

	length := float32(max(a.lenght, b.lenght))
	return float32(minEditCount) / length
}

// SearchList is a list of words that can be searched with string similarity
// algorithms.
type SearchList struct {
	texts   []string
	phrases []*rPhrase
}

// NewSearchList creates a new search list.
func NewSearchList(texts []string) *SearchList {
	list := SearchList{
		texts:   texts,
		phrases: make([]*rPhrase, len(texts))}
	for i := range texts {
		phrase := Parse(texts[i])
		phrase = phrase.WithoutStopwords()
		list.phrases[i] = phrase.asRunes()
	}
	return &list
}

func (s *SearchList) Find(text string) (float32, string) {

	p := Parse(text).WithoutStopwords().asRunes()
	minScore := float32(1.0)
	match := -1
	for i := range s.phrases {
		score := similarityQuotient(p, s.phrases[i])
		if match == -1 || score < minScore {
			minScore = score
			match = i
		}
	}

	if match == -1 {
		return 1.0, ""
	}
	return minScore, s.texts[match]
}
