package gopassgen

import (
	"fmt"
	"testing"
)

func TestRandomPassword(t *testing.T) {
	p := NewPolicy()
	password := Generate(p)

	fmt.Println("Random password:", password)

	if len(password) < 6 || len(password) > 16 {
		t.Errorf("Length is expected to be between 6 and 16, but got %v", len(password))
	}
}

func TestMin15Digits(t *testing.T) {
	p := NewPolicy()
	p.MinDigits = 15
	password := Generate(p)

	fmt.Println("Min 15 digits password:", password)

	if len(password) < 6 || len(password) > 16 {
		t.Errorf("Length is expected to be between 6 and 16, but got %v", len(password))
	}
}

func TestMin10Digits5SpclChars(t *testing.T) {
	p := NewPolicy()
	p.MinDigits = 10
	p.MinSpclChars = 5
	password := Generate(p)

	fmt.Println("Min 10 digits, 5 special chars password:", password)

	if len(password) < 6 || len(password) > 16 {
		t.Errorf("Length is expected to be between 6 and 16, but got %v", len(password))
	}
}
