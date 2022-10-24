package paasaathai

import (
	"fmt"
	"strings"
	"unicode"
)

// Are all the runes in this string in the Thai range of Unicode code points?
func StringIsThai(text string) bool {
	for _, r := range text {
		if !unicode.In(r, unicode.Thai) {
			return false
		}
	}
	return true
}

// Is this rune in the Thai range of Unicode code points?
func RuneIsThai(r rune) bool {
	return unicode.In(r, unicode.Thai)
}

// Return a string representation of the runes in a string.
// Latin is returned as-is, Thai is converted to names.
// Everything else is converted to U+xxxx
func StringToRuneNames(text string) string {
	runes := []rune(text)
	names := make([]string, len(runes))

	for i, r := range runes {
		if unicode.In(r, unicode.Latin) && unicode.IsPrint(r) {
			names[i] = string(r)
			continue
		}
		thaiName, has := RuneToThaiName[r]
		if has {
			names[i] = thaiName
		} else {
			names[i] = fmt.Sprintf("U+%04X", r)
		}
	}
	return strings.Join(names, ", ")
}

// Return a string, converting Thai rune names to actual runes.
// Any other name results in an error
func RuneNamesToString(names []string) (string, error) {
	answer := ""
	for i, name := range names {
		if r, has := ThaiNameToRune[name]; has {
			answer += string(r)
		} else {
			return "", fmt.Errorf("'%s' at index %d is not a Thai codepoint name", name, i)
		}
	}
	return answer, nil
}

func RuneToName(r rune) string {
	n, _ := RuneToThaiName[r]
	return n
}

func NameToRune(name string) rune {
	r, _ := ThaiNameToRune[name]
	return r
}

// THAI_CHARACTER_O_ANG is considered a consonant here (in Thai it also acts
// like a vowel)
// THAI_CHARACTER_RU and THAI_CHARACTER_LU are also considered consonants
func RuneIsConsonant(r rune) bool {
	return r >= THAI_CHARACTER_KO_KAI &&
		r <= THAI_CHARACTER_HO_NOKHUK
}

// Includes all of front, upper, mid, and lower positions
func RuneIsVowel(r rune) bool {
	return RuneIsFrontPositionVowel(r) ||
		RuneIsUpperPositionVowel(r) ||
		RuneIsMidPositionVowel(r) ||
		RuneIsLowerPositionVowel(r)
}

func RuneIsUpperPosition(r rune) bool {
	return RuneIsUpperPositionVowel(r) || RuneIsToneMark(r) || RuneIsUpperPositionSign(r)
}

// We don't consider THAI_CHARACTER_SARA_AM to be upper; we call it mid
// We do consider THAI_CHARACTER_MAITAIKHU to be an upper vowel, since Unicode
// does.
func RuneIsUpperPositionVowel(r rune) bool {
	return r == THAI_CHARACTER_MAI_HAN_AKAT ||
		r == THAI_CHARACTER_SARA_I ||
		r == THAI_CHARACTER_SARA_II ||
		r == THAI_CHARACTER_SARA_UE ||
		r == THAI_CHARACTER_SARA_UEE ||
		r == THAI_CHARACTER_MAITAIKHU
}

// We consider THAI_CHARACTER_SARA_AM to be a mid position vowel
// We consider THAI CHARACTER LAKKHANGYAO to be a vowel too, and thus, mid
// position
// "Mid" position is both vertically "mid", and after an initial consonant
func RuneIsMidPositionVowel(r rune) bool {
	return r == THAI_CHARACTER_SARA_A ||
		r == THAI_CHARACTER_SARA_AA ||
		r == THAI_CHARACTER_SARA_AM ||
		r == THAI_CHARACTER_LAKKHANGYAO
}

// We consider THAI_CHARACTER_PHINTHU to be a lower position vowel
func RuneIsLowerPositionVowel(r rune) bool {
	return r == THAI_CHARACTER_SARA_U ||
		r == THAI_CHARACTER_SARA_UU ||
		r == THAI_CHARACTER_PHINTHU
}

func RuneIsFrontPositionVowel(r rune) bool {
	return r == THAI_CHARACTER_SARA_E ||
		r == THAI_CHARACTER_SARA_AE ||
		r == THAI_CHARACTER_SARA_O ||
		r == THAI_CHARACTER_SARA_AI_MAIMUAN ||
		r == THAI_CHARACTER_SARA_AI_MAIMALAI
}

func RuneIsToneMark(r rune) bool {
	return r == THAI_CHARACTER_MAI_EK ||
		r == THAI_CHARACTER_MAI_THO ||
		r == THAI_CHARACTER_MAI_TRI ||
		r == THAI_CHARACTER_MAI_CHATTAWA
}

func RuneIsDigit(r rune) bool {
	return r >= THAI_DIGIT_ZERO && r <= THAI_DIGIT_NINE
}

// Not a character or digit, "sign", currency, repetition, etc.
func RuneIsSign(r rune) bool {
	return RuneIsUpperPositionSign(r) ||
		RuneIsMidPositionSign(r)
}

func RuneIsUpperPositionSign(r rune) bool {
	return r == THAI_CHARACTER_THANTHAKHAT || // used in Thai
		r == THAI_CHARACTER_NIKHAHIT || // used in Sanskrit?
		r == THAI_CHARACTER_YAMAKKAN // used where?
}

func RuneIsMidPositionSign(r rune) bool {
	return r == THAI_CHARACTER_PAIYANNOI ||
		r == THAI_CURRENCY_SYMBOL_BAHT ||
		r == THAI_CHARACTER_MAIYAMOK ||
		r == THAI_CHARACTER_FONGMAN ||
		r == THAI_CHARACTER_ANGKHANKHU ||
		r == THAI_CHARACTER_KHOMUT
}
