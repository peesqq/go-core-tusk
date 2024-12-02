package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func createVariables() (int, int, int, float64, string, bool, complex64) {
	decimal := 42
	octal := 052
	hexadecimal := 0x2A
	floatVal := 3.14
	str := "GoLang"
	boolean := true
	complexNum := complex(1, 2)
	return decimal, octal, hexadecimal, floatVal, str, boolean, complex64(complexNum)
}

func printVariableTypes(variables ...interface{}) {
	for i, v := range variables {
		fmt.Printf("Variable %d: %v, Type: %s\n", i+1, v, reflect.TypeOf(v))
	}
}

func convertAndCombine(variables ...interface{}) string {
	strs := make([]string, len(variables))
	for i, v := range variables {
		strs[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strs, " ")
}

func stringToRunes(s string) []rune {
	return []rune(s)
}

func hashWithSalt(runes []rune, salt string) string {
	middle := len(runes) / 2
	runesWithSalt := append(runes[:middle], append([]rune(salt), runes[middle:]...)...)
	hash := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hash[:])
}

func main() {
	decimal, octal, hexadecimal, floatVal, str, boolean, complexNum := createVariables()

	printVariableTypes(decimal, octal, hexadecimal, floatVal, str, boolean, complexNum)

	combinedStr := convertAndCombine(decimal, octal, hexadecimal, floatVal, str, boolean, complexNum)
	fmt.Println("Combined String:", combinedStr)

	runes := stringToRunes(combinedStr)
	fmt.Println("Runes:", runes)

	hashed := hashWithSalt(runes, str)
	fmt.Println("Hashed with Salt (SHA256):", hashed)
}
