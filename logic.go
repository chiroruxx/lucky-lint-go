package main

var strokeCounts = map[rune]int{
	'A': 3,
	'B': 3,
	'C': 1,
	'D': 2,
	'E': 4,
	'F': 3,
	'G': 2,
	'H': 3,
	'I': 1,
	'J': 1,
	'K': 3,
	'L': 2,
	'M': 4,
	'N': 3,
	'O': 1,
	'P': 2,
	'Q': 2,
	'R': 3,
	'S': 1,
	'T': 2,
	'U': 1,
	'V': 2,
	'W': 4,
	'X': 2,
	'Y': 3,
	'Z': 3,
	'a': 2,
	'b': 2,
	'c': 1,
	'd': 2,
	'e': 2,
	'f': 2,
	'g': 2,
	'h': 2,
	'i': 2,
	'j': 2,
	'k': 3,
	'l': 1,
	'm': 3,
	'n': 2,
	'o': 1,
	'p': 2,
	'q': 2,
	'r': 2,
	's': 1,
	't': 2,
	'u': 2,
	'v': 2,
	'w': 4,
	'x': 2,
	'y': 2,
	'z': 3,
	'-': 1,
	'+': 2,
	'_': 1,
	'.': 1,
}

func calcStrokeCount(name string) int {
	var count int
	for _, r := range name {
		c, ok := strokeCounts[r]
		if !ok {
			continue
		}
		count += c
	}
	return count
}

type luckyLevel int

const (
	luckyLevelNormal luckyLevel = iota
	luckyLevelGood
	luckyLevelGreat
	luckyLevelBest
)

var luckyLevelMap = map[int]luckyLevel{
	15: luckyLevelBest,
	24: luckyLevelBest,
	31: luckyLevelBest,
	32: luckyLevelBest,
	0:  luckyLevelBest,
	11: luckyLevelGreat,
	16: luckyLevelGreat,
	21: luckyLevelGreat,
	23: luckyLevelGreat,
	41: luckyLevelGreat,
	3:  luckyLevelGood,
	5:  luckyLevelGood,
	6:  luckyLevelGood,
	8:  luckyLevelGood,
	13: luckyLevelGood,
	18: luckyLevelGood,
	25: luckyLevelGood,
	29: luckyLevelGood,
	33: luckyLevelGood,
	37: luckyLevelGood,
	44: luckyLevelGood,
	39: luckyLevelGood,
	45: luckyLevelGood,
	47: luckyLevelGood,
	48: luckyLevelGood,
	51: luckyLevelGood,
}

const maxStrokeCount = 52

func getLuckyLevel(strokeCount int) luckyLevel {
	c := strokeCount % maxStrokeCount
	level, ok := luckyLevelMap[c]
	if !ok {
		return luckyLevelNormal
	}
	return level
}
