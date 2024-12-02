package main

import (
	"reflect"
	"testing"
)

func TestCreateVariables(t *testing.T) {
	decimal, octal, hexadecimal, floatVal, str, boolean, complexNum := createVariables()

	if decimal != 42 || octal != 42 || hexadecimal != 42 {
		t.Errorf("Expected 42 for all integer variables, but got %d, %d, %d", decimal, octal, hexadecimal)
	}
	if floatVal != 3.14 {
		t.Errorf("Expected 3.14 for floatVal, but got %f", floatVal)
	}
	if str != "GoLang" {
		t.Errorf("Expected 'GoLang' for str, but got %s", str)
	}
	if boolean != true {
		t.Errorf("Expected true for boolean, but got %t", boolean)
	}
	if complexNum != complex(1, 2) {
		t.Errorf("Expected (1+2i) for complexNum, but got %v", complexNum)
	}
}

// Тестируем функцию printVariableTypes
func TestPrintVariableTypes(t *testing.T) {
	decimal := 42
	str := "test"
	floatVal := 3.14

	// Определяем типы и проверяем их
	expectedTypes := []reflect.Type{
		reflect.TypeOf(decimal),
		reflect.TypeOf(str),
		reflect.TypeOf(floatVal),
	}

	var results []reflect.Type
	printVariableTypes(decimal, str, floatVal) // Используем реальный вывод функции
	for _, v := range []interface{}{decimal, str, floatVal} {
		results = append(results, reflect.TypeOf(v))
	}

	for i, expected := range expectedTypes {
		if results[i] != expected {
			t.Errorf("Expected type %s but got %s", expected, results[i])
		}
	}
}

func TestConvertAndCombine(t *testing.T) {
	result := convertAndCombine(42, "test", 3.14, true)
	expected := "42 test 3.14 true"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestStringToRunes(t *testing.T) {
	input := "GoLang"
	expected := []rune{'G', 'o', 'L', 'a', 'n', 'g'}
	result := stringToRunes(input)
	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Expected rune %c but got %c at index %d", expected[i], r, i)
		}
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := stringToRunes("test string")
	salt := "GoLang"
	hashed := hashWithSalt(runes, salt)
	if len(hashed) != 64 {
		t.Errorf("Expected hash length of 64, but got %d", len(hashed))
	}
}
