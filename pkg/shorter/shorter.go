package shorter

const (
	evenDivider rune = 'Z'
	evenStart   rune = 'A'
	oddStart    rune = 'a'
	oddDivider  rune = 'z'
	base             = evenDivider - evenStart
)

type Shorter struct {
	seeds      []rune
	blockSize  int
	blockCount int
}

func NewShorter(blockSize int, blockCount int, seed rune) *Shorter {
	var (
		seeds = make([]rune, blockCount)
		i     rune
	)
	for i = 0; i < rune(blockCount); i++ {
		seeds[i] = seed * (i + 1) % oddDivider
	}

	return &Shorter{blockSize: blockSize, blockCount: blockCount, seeds: seeds}
}

func (sh Shorter) Short(url string) (short string) {
	return sh.encode([]rune(url))
}

func (sh Shorter) encode(url []rune) string {
	var hash = make([]rune, sh.blockCount*sh.blockSize)

	for i, r := range url {
		hash[i%len(hash)] = sh.start(i/sh.blockCount+1) +
			(hash[i%len(hash)]+
				r%sh.seeds[(i/sh.blockCount+1)%len(sh.seeds)])%base
	}

	return string(hash)
}

func (sh Shorter) start(block int) rune {
	if block%2 == 0 {
		return evenStart
	}

	return oddStart
}
