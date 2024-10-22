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
	keyLength := len(key)
	for i, letter := range message {
		if letter == ' ' {
			result += " "
		} else {
			keyLetter := key[i%keyLength]
			result += shiftByLetter(string(letter), string(keyLetter))
		}
	}
	return result
}

// Scytale cipher
//
// Encrypts a message by simulating a scytale cipher.
//
// A scytale is a cylinder around which you can wrap a long strip of
//
//	parchment that contained a string of apparent gibberish. The parchment,
//	when read using the scytale, would reveal a message due to every nth
//	letter now appearing next to each other, revealing a proper message.
//
// This encryption method is obsolete and should never be used to encrypt
//
//	data in production settings.
//
// You may read more about the method here:
//
//	https://en.wikipedia.org/wiki/Scytale
//
// You may follow this algorithm to implement a scytale-style cipher:
//  1. Take a message to be encoded and a "shift" number.
//     For this example, we will use "INFORMATION_AGE" as
//     the message and 3 as the shift.
//  2. Check if the length of the message is a multiple of
//     the shift. If it is not, add additional underscores
//     to the end of the message until it is.
//     In this example, "INFORMATION_AGE" is already a multiple of 3,
//     so we will leave it alone.
//  3. This is the tricky part. Construct the encoded message.
//     For each index i in the encoded message, use the character at the index
//     (i // shift) + (len(message) // shift) * (i % shift) of the raw message.
//     If this number doesn't make sense, you can play around with the cipher at
//     https://dencode.com/en/cipher/scytale to try to understand it.
//  4. Return the encoded message. In this case,
//     the output should be "IMNNA_FTAOIGROE".
//
// Example
// scytaleCipher("INFORMATION_AGE", 3) -> "IMNNA_FTAOIGROE"
// scytaleCipher("INFORMATION_AGE", 4) -> "IRIANMOGFANEOT__"
// scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8) -> "AOTSRIOALRH_EMRNGIMA_PTT"
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - the message, encoded appropriately.

func scytaleCipher(message string, shift int) string {
	messageLen := len(message)
	if messageLen%shift != 0 {
		// Add underscores to make the message length a multiple of the shift
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
//
// Decrypts a message that was originally encrypted with the `scytaleCipher` function above.
//
// Example:
// scytaleDecipher('IMNNA_FTAOIGROE', 3) -> 'INFORMATION_AGE'
// scytaleDecipher('AOTSRIOALRH_EMRNGIMA_PTT', 8) -> 'ALGORITHMS_ARE_IMPORTANT'
// scytaleDecipher('IRIANMOGFANEOT__', 4) -> 'INFORMATION_AGE_'
//
// Params:
// - message, a string of uppercase English letters and underscores. Underscores represent spaces.
// - shift, a positive integer that does not exceed the length of the message.
//
// Returns:
// - The message, decoded appropriately.
func scytaleDecipher(message string, shift int) string {
	// Replace this with your code
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
