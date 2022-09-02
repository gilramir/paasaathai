package paasaathai

import (
	"unicode"
)

// From Unicode Thai code points
// https://www.unicode.org/charts/PDF/U0E00.pdf
const (
	/* ก */ THAI_CHARACTER_KO_KAI = 0x0E01
	/* ข */ THAI_CHARACTER_KHO_KHAI = 0x0E02
	/* ฃ */ THAI_CHARACTER_KHO_KHUAT = 0x0E03
	/* ค */ THAI_CHARACTER_KHO_KHWAI = 0x0E04
	/* ฅ */ THAI_CHARACTER_KHO_KHON = 0x0E05
	/* ฆ */ THAI_CHARACTER_KHO_RAKHANG = 0x0E06
	/* ง */ THAI_CHARACTER_NGO_NGU = 0x0E07
	/* จ */ THAI_CHARACTER_CHO_CHAN = 0x0E08
	/* ฉ */ THAI_CHARACTER_CHO_CHING = 0x0E09
	/* ช */ THAI_CHARACTER_CHO_CHANG = 0x0E0A
	/* ซ */ THAI_CHARACTER_SO_SO = 0x0E0B
	/* ฌ */ THAI_CHARACTER_CHO_CHOE = 0x0E0C
	/* ญ */ THAI_CHARACTER_YO_YING = 0x0E0D
	/* ฎ */ THAI_CHARACTER_DO_CHADA = 0x0E0E
	/* ฏ */ THAI_CHARACTER_TO_PATAK = 0x0E0F
	/* ฐ */ THAI_CHARACTER_THO_THAN = 0x0E10
	/* ฑ */ THAI_CHARACTER_THO_NANGMONTHO = 0x0E11
	/* ฒ */ THAI_CHARACTER_THO_PHUTHAO = 0x0E12
	/* ณ */ THAI_CHARACTER_NO_NEN = 0x0E13
	/* ด */ THAI_CHARACTER_DO_DEK = 0x0E14
	/* ต */ THAI_CHARACTER_TO_TAO = 0x0E15
	/* ถ */ THAI_CHARACTER_THO_THUNG = 0x0E16
	/* ท */ THAI_CHARACTER_THO_THAHAN = 0x0E17
	/* ธ */ THAI_CHARACTER_THO_THONG = 0x0E18
	/* น */ THAI_CHARACTER_NO_NU = 0x0E19
	/* บ */ THAI_CHARACTER_BO_BAIMAI = 0x0E1A
	/* ป */ THAI_CHARACTER_PO_PLA = 0x0E1B
	/* ผ */ THAI_CHARACTER_PHO_PHUNG = 0x0E1C
	/* ฝ */ THAI_CHARACTER_FO_FA = 0x0E1D
	/* พ */ THAI_CHARACTER_PHO_PHAN = 0x0E1E
	/* ฟ */ THAI_CHARACTER_FO_FAN = 0x0E1F
	/* ภ */ THAI_CHARACTER_PHO_SAMPHAO = 0x0E20
	/* ม */ THAI_CHARACTER_MO_MA = 0x0E21
	/* ย */ THAI_CHARACTER_YO_YAK = 0x0E22
	/* ร */ THAI_CHARACTER_RO_RUA = 0x0E23
	/* ฤ */ THAI_CHARACTER_RU = 0x0E24
	/* ล */ THAI_CHARACTER_LO_LING = 0x0E25
	/* ฦ */ THAI_CHARACTER_LU = 0x0E26
	/* ว */ THAI_CHARACTER_WO_WAEN = 0x0E27
	/* ศ */ THAI_CHARACTER_SO_SALA = 0x0E28
	/* ษ */ THAI_CHARACTER_SO_RUSI = 0x0E29
	/* ส */ THAI_CHARACTER_SO_SUA = 0x0E2A
	/* ห */ THAI_CHARACTER_HO_HIP = 0x0E2B
	/* ฬ */ THAI_CHARACTER_LO_CHULA = 0x0E2C
	/* อ */ THAI_CHARACTER_O_ANG = 0x0E2D
	/* ฮ */ THAI_CHARACTER_HO_NOKHUK = 0x0E2E
	/* ฯ */ THAI_CHARACTER_PAIYANNOI = 0x0E2F
	/* ะ */ THAI_CHARACTER_SARA_A = 0x0E30
	/* $ั */ THAI_CHARACTER_MAI_HAN_AKAT = 0x0E31
	/* า */ THAI_CHARACTER_SARA_AA = 0x0E32
	/* ำ */ THAI_CHARACTER_SARA_AM = 0x0E33
	/* $ิ */ THAI_CHARACTER_SARA_I = 0x0E34
	/* $ี */ THAI_CHARACTER_SARA_II = 0x0E35
	/* $ึ */ THAI_CHARACTER_SARA_UE = 0x0E36
	/* $ื */ THAI_CHARACTER_SARA_UEE = 0x0E37
	/* $ุ */ THAI_CHARACTER_SARA_U = 0x0E38
	/* $ู */ THAI_CHARACTER_SARA_UU = 0x0E39
	/* $ฺ */ THAI_CHARACTER_PHINTHU = 0x0E3A
	/* ฿ */ THAI_CURRENCY_SYMBOL_BAHT = 0x0E3F
	/* เ */ THAI_CHARACTER_SARA_E = 0x0E40
	/* แ */ THAI_CHARACTER_SARA_AE = 0x0E41
	/* โ */ THAI_CHARACTER_SARA_O = 0x0E42
	/* ใ */ THAI_CHARACTER_SARA_AI_MAIMUAN = 0x0E43
	/* ไ */ THAI_CHARACTER_SARA_AI_MAIMALAI = 0x0E44
	/* ๅ */ THAI_CHARACTER_LAKKHANGYAO = 0x0E45
	/* ๆ */ THAI_CHARACTER_MAIYAMOK = 0x0E46
	/* $็ */ THAI_CHARACTER_MAITAIKHU = 0x0E47
	/* $่ */ THAI_CHARACTER_MAI_EK = 0x0E48
	/* $้ */ THAI_CHARACTER_MAI_THO = 0x0E49
	/* $๊ */ THAI_CHARACTER_MAI_TRI = 0x0E4A
	/* $๋ */ THAI_CHARACTER_MAI_CHATTAWA = 0x0E4B
	/* $์ */ THAI_CHARACTER_THANTHAKHAT = 0x0E4C
	/* $ํ */ THAI_CHARACTER_NIKHAHIT = 0x0E4D
	/* $๎ */ THAI_CHARACTER_YAMAKKAN = 0x0E4E
	/* ๏ */ THAI_CHARACTER_FONGMAN = 0x0E4F
	/* ๐ */ THAI_DIGIT_ZERO = 0x0E50
	/* ๑ */ THAI_DIGIT_ONE = 0x0E51
	/* ๒ */ THAI_DIGIT_TWO = 0x0E52
	/* ๓ */ THAI_DIGIT_THREE = 0x0E53
	/* ๔ */ THAI_DIGIT_FOUR = 0x0E54
	/* ๕ */ THAI_DIGIT_FIVE = 0x0E55
	/* ๖ */ THAI_DIGIT_SIX = 0x0E56
	/* ๗ */ THAI_DIGIT_SEVEN = 0x0E57
	/* ๘ */ THAI_DIGIT_EIGHT = 0x0E58
	/* ๙ */ THAI_DIGIT_NINE = 0x0E59
	/* ๚ */ THAI_CHARACTER_ANGKHANKHU = 0x0E5A
	/* ๛ */ THAI_CHARACTER_KHOMUT = 0x0E5B
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

// We don't consdier THAI_CHARACTER_SARA_AM to be upper; we call it mid
// We do consider THAI_CHARACTER_MAITAIKHU to be an upper vowel
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
func RuneIsMidPositionVowel(r rune) bool {
	return r == THAI_CHARACTER_SARA_A ||
		r == THAI_CHARACTER_SARA_AA ||
		r == THAI_CHARACTER_SARA_AM ||
		r == THAI_CHARACTER_LAKKHANGYAO
}

// We consdier THAI_CHARACTER_PHINTHU to be a lower position vowel
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
	return r == THAI_CHARACTER_THANTHAKHAT ||
		r == THAI_CHARACTER_NIKHAHIT ||
		r == THAI_CHARACTER_YAMAKKAN
}

func RuneIsMidPositionSign(r rune) bool {
	return r == THAI_CHARACTER_PAIYANNOI ||
		r == THAI_CURRENCY_SYMBOL_BAHT ||
		r == THAI_CHARACTER_MAIYAMOK ||
		r == THAI_CHARACTER_FONGMAN ||
		r == THAI_CHARACTER_ANGKHANKHU ||
		r == THAI_CHARACTER_KHOMUT
}
