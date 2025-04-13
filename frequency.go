package caesarcipher

type Values map[rune]int
type Fractions map[rune]float32

type Frequency struct {
	Name   string
	Values map[rune]int
}

func NewFrequency(name string) *Frequency {
	return &Frequency{
		Values: map[rune]int{
			'a': 0,
			'b': 0,
			'c': 0,
			'd': 0,
			'e': 0,
			'f': 0,
			'g': 0,
			'h': 0,
			'i': 0,
			'j': 0,
			'k': 0,
			'l': 0,
			'm': 0,
			'n': 0,
			'o': 0,
			'p': 0,
			'q': 0,
			'r': 0,
			's': 0,
			't': 0,
			'u': 0,
			'v': 0,
			'w': 0,
			'x': 0,
			'y': 0,
			'z': 0,
		},
	}
}

func (f *Frequency) Merge(f2 Frequency) {
	for k, v := range f2.Values {
		//we only merge existing keys
		if _, ok := f.Values[k]; ok {
			f.Values[k] += v
		}
	}
}

func (f *Frequency) ToFractions() Fractions {
	fractions := map[rune]float32{}
	total := float32(0)
	for _, v := range f.Values {
		total += float32(v)
	}
	for k, v := range f.Values {
		fractions[k] = float32(v) / total
	}
	return fractions
}
