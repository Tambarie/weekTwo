package main

import (
	"reflect"
	"testing"
)

func TestCalculator(t *testing.T) {
	var inputStrings = []string{"2+3","2-3","3*2","4-5","2/2","5*4","7+9"}
	get := calculator(inputStrings...)
	result := []float64{5,-1,6,-1,1,20,16}
	if !reflect.DeepEqual(result,get) {
		t.Fatalf("expected : %v , got%v", result,get)
	}
}

func TestAddition(t *testing.T) {
	number := []float64{2,3,4}
	result := float64(9)
	get := addition(number)

	if !reflect.DeepEqual(result,get) {
		t.Fatalf("expected : %v , got%v", result,get)
	}
}

func TestSubtraction(t *testing.T) {
	number := []float64{3,4}
	result := float64(-1)
	get := subtraction(number)

	if !reflect.DeepEqual(result,get) {
		t.Fatalf("expected : %v , got%v", result,get)
	}
}

func TestMultiplication(t *testing.T) {
	number := []float64{3,4}
	result := float64(12)
	get := multiplication(number)

	if !reflect.DeepEqual(result,get) {
		t.Fatalf("expected : %v , got%v", result,get)
	}
}

func TestDivision(t *testing.T) {
	number := []float64{6,2}
	result := float64(3)
	get := division(number)

	if !reflect.DeepEqual(result,get) {
		t.Fatalf("expected : %v , got%v", result,get)
	}
}