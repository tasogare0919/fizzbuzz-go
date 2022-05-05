package main

import (
	"strings"
	"testing"
)

func TestFizz(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(15)) {
		t.Errorf("Correct")
	}
	if !strings.EqualFold("Fizz", FizzBuzz(3)) {
		t.Errorf("Correct")
	}
	if !strings.EqualFold("Buzz", FizzBuzz(5)) {
		t.Errorf("Correct")
	}
	if !strings.EqualFold("98", FizzBuzz(98)) {
		t.Errorf("Correct")
	}
}
