// Placeholder
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Shift Letter
	fmt.Println(shiftLetter("A", 0))
	fmt.Println(shiftLetter("A", 2))
	fmt.Println(shiftLetter("Z", 24))
	fmt.Println(shiftLetter("X", 5))
	fmt.Println(shiftLetter("Y", 1))
	//Caeser Cipher
	fmt.Println(caesarCipher("Ssample", 5))
	// Shift by letter
	fmt.Println(shiftByLetter("A", "A"))
	fmt.Println(shiftByLetter("A", "C"))
	fmt.Println(shiftByLetter("B", "K"))
	fmt.Println(shiftByLetter(" ", " "))
	// Vigenere cipher
	fmt.Println(vigenereCipher("A CON", "KEY"))
	// Scytale cipher
	fmt.Println(scytaleCipher("INFORMATION_AGE", 3))
	fmt.Println(scytaleCipher("INFORMATION_AGE", 4))
	fmt.Println(scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8))
	// Scytale decipher
	fmt.Println(scytaleDecipher("IMNNA_FTAOIGROE", 3))
	fmt.Println(scytaleDecipher("AOTSRIOALRH_EMRNGIMA_PTT", 8))
	fmt.Println(scytaleDecipher("IRIANMOGFANEOT__", 4))
}

// Shift letter
func shiftLetter(letter string, shift int) string {
	if letter == "" {
		return ""
	}
	letter = strings.ToUpper(letter)
	start := 'A'
	shifted := (int(int32(letter[0])) - int(start) + shift) % 26
	return string(start + int32(shifted))
}

// Caesar cipher
func caesarCipher(message string, shift int) string {
	ciphered := ""
	for _, number := range message {
		ciphered += shiftLetter(string(number), shift)
	}
	return ciphered
}

// Shift by letter
func shiftByLetter(letter string, letterShift string) string {
	if letter == "" {
		return ""
	}
	shift := int(letterShift[0] - 'A')
	return shiftLetter(letter, shift)
}

// Vigenere cipher
func vigenereCipher(message string, key string) string {
	result := ""
	keyLength := len(key)
	for i, letter := range message {
		if letter == ' ' {
			result += " "
		} else {
			letterShift := key[i%keyLength]
			result += shiftByLetter(string(letter), string(letterShift))
		}
	}
	return result
}

// Scytale cipher
func scytaleCipher(message string, shift int) string {
	messageLen := len(message)
	if messageLen%shift != 0 {
		for i := 0; i < shift-(messageLen%shift); i++ {
			message += "_"
		}
	}

	encoded := make([]rune, len(message))
	for i := 0; i < len(message); i++ {
		row := i / shift
		col := i % shift
		newIndex := col*(len(message)/shift) + row
		encoded[i] = rune(message[newIndex])
	}
	return string(encoded)
}

// Scytale decipher
func scytaleDecipher(message string, shift int) string {
	decoded := make([]rune, len(message))
	messageLen := len(message)
	for i := 0; i < messageLen; i++ {
		row := i / (messageLen / shift)
		col := i % (messageLen / shift)
		newIndex := col*shift + row
		decoded[i] = rune(message[newIndex])
	}
	return string(decoded)
}
