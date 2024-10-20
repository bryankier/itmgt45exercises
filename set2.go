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
	fmt.Println(vigenereCipher("A C", "KEY"))
}

// Shift letter
func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
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
	if letter == " " {
		return " "
	}
	shift := int(letterShift[0] - 'A')
	return shiftLetter(letter, shift)
}

// Vigenere cipher
//
// Encrypt a message using a keyphrase instead of a static number.
// Every letter in the message is shifted by the number represented by the respective letter in the key.
// Spaces are ignored.
//
// Example
// vigenereCipher("A C", "KEY") -> "K A"
//
// If needed, the keyphrase is extended to match the length of the key.
// If the key is "KEY" and the message is "LONGTEXT", the key will be extended to "KEYKEYKE".
//
// Params:
// - message, a string of uppercase english letters and/or spaces
// - key, a string of uppercase English letters with no spaces. Will not exceed the length of the message.
//
// Returns:
// - the message, shifted appropriately
func vigenereCipher(message string, key string) string {
	result := ""
	keyLen := len(key)
	for i, letter := range message {
		if letter == ' ' {
			result += " "
		} else {
			keyLetter := key[i%keyLen]
			result += shiftByLetter(string(letter), string(keyLetter))
		}
	}
	return result
}
