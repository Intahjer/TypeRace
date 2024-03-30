package stringgen

import (
	"math/rand"
)

type Difficulty string

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
	Super  Difficulty = "Super"
)

var test = "HAHAHAHAHA:12 3515643526:[]./.HAHAH DASDA SFDA ASFK ASJFAS LKFJAS FASLF KAsdjfaskl fjaslkfasjsaFASFSJ AFAfsaflkdja sdklfashfasASJFASFKASFJASLFKJDKLGJ SDKGDJSFg lk;sbdfslkj jfalks;j;alsfjas;fjasflaksjf;lsfjas;flkjasf;lksajfs;lakfjas;lkfjsa;lkf jko joifhjjngfjkngvjskdlnvhuybndnbajksfglhujiaerkhnugjvbnzdsjfkgnbljkasdhglkjbvnsuidjhfhgnjklheqagjndjklafvhnljskdhgbuijsdfhguiaerjhjgjkdhgfjksdlhgfskjhgsldkj"
var STRING_LIMIT = 300

func GetString(difficulty Difficulty) string {
	str := SimplifyString(test, STRING_LIMIT)
	if difficulty == Medium {
		str = medString(str)
	} else if difficulty == Hard {
		str = hardString(str)
	} else if difficulty == Super {
		str = superString(str)
	}
	return str
}

func superString(str string) string {
	lastChar := ' '
	upper := []rune{}
	for _, char := range str {
		newChar := char
		if isSpace(lastChar) && isLowerLetter(newChar) {
			shouldAdd := rand.Intn(2)
			if shouldAdd == 1 {
				newChar = shiftToUpper(newChar)
			}
		} else if isSpace(newChar) {
			shouldAdd := rand.Intn(5)
			if shouldAdd == 0 {
				upper = append(upper, rune(rand.Intn(11)+48))
			} else if shouldAdd == 1 {
				upper = append(upper, rune(rand.Intn(15)+33))
			} else if shouldAdd == 2 {
				upper = append(upper, rune(rand.Intn(9)+58))
			} else if shouldAdd == 3 {
				upper = append(upper, rune(rand.Intn(6)+91))
			} else if shouldAdd == 4 {
				upper = append(upper, rune(rand.Intn(4)+123))
			}
		}
		upper = append(upper, newChar)
		lastChar = newChar
	}
	return string(upper)
}

func hardString(str string) string {
	lastChar := ' '
	upper := []rune{}
	for _, char := range str {
		newChar := char
		if isSpace(lastChar) && isLowerLetter(newChar) {
			shouldAdd := rand.Intn(2)
			if shouldAdd == 1 {
				newChar = shiftToUpper(newChar)
			}
		} else if isSpace(newChar) {
			shouldAdd := rand.Intn(5)
			if shouldAdd == 2 {
				upper = append(upper, rune(rand.Intn(9)+58))
			} else if shouldAdd == 3 {
				upper = append(upper, rune(rand.Intn(6)+91))
			} else if shouldAdd == 4 {
				upper = append(upper, rune(rand.Intn(4)+123))
			}
		}
		upper = append(upper, newChar)
		lastChar = newChar
	}
	return string(upper)
}

func medString(str string) string {
	lastChar := ' '
	upper := []rune{}
	for _, char := range str {
		newChar := char
		if isSpace(lastChar) && isLowerLetter(newChar) {
			shouldAdd := rand.Intn(2)
			if shouldAdd == 1 {
				newChar = shiftToUpper(newChar)
			}
		} else if isSpace(newChar) {
			shouldAdd := rand.Intn(5)
			if shouldAdd == 0 {
				upper = append(upper, '.')
			} else if shouldAdd == 1 {
				upper = append(upper, ';')
			} else if shouldAdd == 2 {
				upper = append(upper, ':')
			} else if shouldAdd == 3 {
				upper = append(upper, ',')
			}
		}
		upper = append(upper, newChar)
		lastChar = newChar
	}
	return string(upper)
}

func SimplifyString(str string, limit int) string {
	simplified := []rune{}
	localLimit := STRING_LIMIT
	for i, a := range str {
		if i < localLimit {
			if isLowerLetter(a) || isSpace(a) {
				simplified = append(simplified, a)
			} else if isUpperLeter(a) {
				simplified = append(simplified, shiftToLower(a))
			} else {
				localLimit++
			}
		}
	}
	return string(simplified)
}

func isLowerLetter(a rune) bool {
	return a > 96 && a < 123
}

func isSpace(a rune) bool {
	return a == 32
}

func isUpperLeter(a rune) bool {
	return a > 64 && a < 91
}

func shiftToLower(a rune) rune {
	return a + 32
}

func shiftToUpper(a rune) rune {
	return a - 32
}
