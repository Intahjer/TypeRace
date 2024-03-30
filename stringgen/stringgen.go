package stringgen

import (
	c "TypeGame/constants"
	"math/rand"
)

type Difficulty int

const (
	Easy   Difficulty = 0
	Medium Difficulty = 1
	Hard   Difficulty = 2
)

var test = "HAHAHAHAHA:123515643526:[]./.HAHAH DASDASFDA ASFK ASJFASLKFJAS FASLF KAsdjfasklfjaslkfasjsaFASFSJAFAfsaflkdjasdklfashfasASJFASFKASFJASLFKJDKLGJ SDKGDJSFg lk;sbdfslkj jfalks;j;alsfjas;fjasflaksjf;lsfjas;flkjasf;lksajfs;lakfjas;lkfjsa;lkf jko joifhjjngfjkngvjskdlnvhuybndnbajksfglhujiaerkhnugjvbnzdsjfkgnbljkasdhglkjbvnsuidjhfhgnjklheqagjndjklafvhnljskdhgbuijsdfhguiaerjhjgjkdhgfjksdlhgfskjhgsldkj"
var STRING_LIMIT = 100

func GetString(Difficulty) string {
	return SimplifyString(test, STRING_LIMIT)
}

func UpperString(str string) string {
	lastChar := ' '
	upper := []rune{}
	for idx, char := range str {
		newChar := char
		if (isSpace(lastChar) && isLowerLetter(char)){
			bool shouldAdd := rand.Intn(1);
 		if(shouldAdd){
			newChar = shiftToUpper(newChar)
		}
			upper = append(upper, newChar)
	lastChar = newChar
	}
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

func SimpleName(str string) string {
	simpleName := SimplifyString(str)
	if simpleName == "" {
		return c.DEFAULT_NAME
	}
	return simpleName
}
