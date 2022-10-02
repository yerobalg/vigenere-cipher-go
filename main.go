package VegnereCipher

import (
	"fmt"
	"regexp"
	"errors"
)

func vegnereCipher(plainText string, key string) (string, error) {
	if !regexp.MustCompile(`^[A-Z]+$`).MatchString(plainText) || !regexp.MustCompile(`^[A-Z]+$`).MatchString(key) {
		return "", errors.New("Only uppercase letters are allowed")
	}

	// if key length is less than the plainText length, then we need to repeat the key until it is the same length as the plainText
	for len(key) < len(plainText) {
		key += key
	}

	// if key length is greater than the plainText length, then we need to truncate the key to the same length as the plainText
	if len(key) > len(plainText) {
		key = key[:len(plainText)]
	}

	fmt.Println(key)

	// create a slice of runes for the key and plainText
	byteKey := []rune(key)
	bytePlainText := []byte(plainText)
	decipheredText := []byte{}

	// decipher the plainText
	for i, char := range bytePlainText {
		plainTextInt := int(char) - 65
		keyInt := int(byteKey[i]) - 65
		encryptedChar := byte(cipherFormula(plainTextInt, keyInt) + 65)
		decipheredText = append(decipheredText, encryptedChar)
	}

	return string(decipheredText), nil
}

func vegnereDecipher(cipherText string, key string) (string, error) {
	if !regexp.MustCompile(`^[A-Z]+$`).MatchString(cipherText) || !regexp.MustCompile(`^[A-Z]+$`).MatchString(key) {
		return "", errors.New("Only uppercase letters are allowed")
	}
	
	// if key length is less than the cipherText length, then we need to repeat the key until it is the same length as the cipherText
	for len(key) < len(cipherText) {
		key += key
	}

	// if key length is greater than the cipherText length, then we need to truncate the key to the same length as the cipherText
	if len(key) > len(cipherText) {
		key = key[:len(cipherText)]
	}

	fmt.Println(key)

	// create a slice of runes for the key and cipherText
	byteKey := []rune(key)
	byteCipherText := []byte(cipherText)
	decipheredText := []byte{}

	// decipher the cipherText
	for i, char := range byteCipherText {
		cipherTextInt := int(char) - 65
		keyInt := int(byteKey[i]) - 65
		encryptedChar := byte(dechipherFormula(cipherTextInt, keyInt) + 65)
		fmt.Println(byte(dechipherFormula(cipherTextInt, keyInt)))
		decipheredText = append(decipheredText, encryptedChar)
	}

	return string(decipheredText), nil
}

func cipherFormula(char int, key int) int {
	return (char + key) % 26
}

func dechipherFormula(char int, key int) int {
	return (char - key + 26) % 26
}
