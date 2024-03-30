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
	str := SimplifyString(generateStr(), STRING_LIMIT)
	switch difficulty {
	case Easy:
		return easyString(str)
	case Medium:
		return medString(str)
	case Hard:
		return hardString(str)
	case Super:
		return superString(str)
	default:
		return str
	}
}

func generateStr() string {
	return test
}

func superString(str string) string {
	lastChar := ' '
	super := []rune{}
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
				super = append(super, rune(rand.Intn(11)+48))
			} else if shouldAdd == 1 {
				super = append(super, rune(rand.Intn(15)+33))
			} else if shouldAdd == 2 {
				super = append(super, rune(rand.Intn(9)+58))
			} else if shouldAdd == 3 {
				super = append(super, rune(rand.Intn(6)+91))
			} else if shouldAdd == 4 {
				super = append(super, rune(rand.Intn(4)+123))
			}
		}
		super = append(super, newChar)
		lastChar = newChar
	}
	return string(super)
}

func hardString(str string) string {
	lastChar := ' '
	hard := []rune{}
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
				hard = append(hard, rune(rand.Intn(8)+58))
			} else if shouldAdd == 3 {
				hard = append(hard, rune(rand.Intn(6)+91))
			} else if shouldAdd == 4 {
				hard = append(hard, rune(rand.Intn(4)+123))
			}
		}
		hard = append(hard, newChar)
		lastChar = newChar
	}
	return string(hard)
}

func medString(str string) string {
	lastChar := ' '
	med := []rune{}
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
				med = append(med, '.')
			} else if shouldAdd == 1 {
				med = append(med, ';')
			} else if shouldAdd == 2 {
				med = append(med, ':')
			} else if shouldAdd == 3 {
				med = append(med, ',')
			}
		}
		med = append(med, newChar)
		lastChar = newChar
	}
	return string(med)
}

func easyString(str string) string {
	easy := []rune{}
	for _, char := range str {
		newChar := char
		if isUpperLeter(char) {
			newChar = shiftToLower(char)
		}
		easy = append(easy, newChar)
	}
	return string(easy)
}

func SimplifyString(str string, limit int) string {
	simplified := []rune{}
	localLimit := limit
	for i, a := range str {
		if i < localLimit {
			if isLowerLetter(a) || isSpace(a) || isUpperLeter(a) {
				simplified = append(simplified, a)
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
