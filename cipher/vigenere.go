package cipher

import (
	"errors"
	"fmt"
	"unicode"
)

type Vigenere struct {
	Alphabet    []rune
	TabulaRecta [][]rune
	key         string
}

func New(characters, key string) *Vigenere {
	alphabet := []rune(characters)
	tabulaRecta := generateTabulaRecta(alphabet)

	vigenere := Vigenere{alphabet, tabulaRecta, key}
	return &vigenere
}

func generateTabulaRecta(alphabet []rune) [][]rune {
	alphabetLen := len(alphabet)
	tabulaRecta := make([][]rune, alphabetLen)

	for j := 0; j < alphabetLen; j++ {
		tabulaRecta[j] = make([]rune, alphabetLen)
		for i := 0; i < alphabetLen-j; i++ {
			tabulaRecta[j][i] = alphabet[j+i]
		}
		for i := 0; i < j; i++ {
			tabulaRecta[j][alphabetLen-j+i] = alphabet[i]
		}
	}

	return tabulaRecta
}

func (v Vigenere) Encrypt(plainText string) string {
	textLength := len(plainText)
	keyLength := len(v.key)

	text := []rune(plainText)

	encryped := make([]rune, textLength)

	for i := 0; i < textLength; i++ {
		keyChar := rune(v.key[i%keyLength])
		textChar := text[i]
		if textChar == ' ' {
			encryped[i] = textChar
			continue
		}
		lower := false
		if isLower(textChar) {
			textChar = unicode.ToUpper(textChar)
			lower = true
		}
		e := v.getChar(textChar, keyChar)
		if lower {
			e = unicode.ToLower(e)
		}
		encryped[i] = e
	}

	return string(encryped)
}

func (v Vigenere) getChar(textChar, keyChar rune) rune {
	sliceIndex, err := atIndex(v.Alphabet, keyChar)
	if err != nil {
		panic(fmt.Sprintf("Key character: %s is not present in alphabet: %q", string(keyChar), v.Alphabet))
	}

	row := v.TabulaRecta[sliceIndex]
	charIndex, err := atIndex(v.Alphabet, textChar)
	if err != nil {
		panic(fmt.Sprintf("Text char: %s is not present in alphabet: %q", string(textChar), v.Alphabet))
	}

	return row[charIndex]
}

func atIndex(t []rune, item rune) (int, error) {
	for i, el := range t {
		if el == item {
			return i, nil
		}
	}
	return -1, errors.New("Character not found")
}

// https://stackoverflow.com/a/59293875
func isLower(r rune) bool {
	if !unicode.IsLower(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}
