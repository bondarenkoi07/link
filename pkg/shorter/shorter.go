package shorter

const (
	evenDivider = 'Z'
	evenStart   = 'A'
	oddDivider  = 'z'
	oddStart    = 'a'
	base        = evenDivider - evenStart
)

type Shorter struct {
	seeds      []rune
	blockSize  int
	blockCount int
}

func (sh Shorter) Short(url string) (short string) {
	return sh.encode([]rune(url))
}

func (sh Shorter) encode(url []rune) string {
	var hash = make([]rune, sh.blockCount*sh.blockSize)

	for i, r := range url {
		hash[i%len(hash)] = sh.start(i/sh.blockSize+1) +
			(hash[i%len(hash)]+
				r%sh.seeds[(i/sh.blockSize+1)%len(sh.seeds)])%base
	}

	return string(hash)
}

func (sh Shorter) start(block int) rune {
	if block%2 == 0 {
		return evenStart
	}

	return oddStart
}
