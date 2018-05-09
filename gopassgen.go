package gopassgen

/*
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopassgen#gopassgen

Usage:
	p := gopassgen.NewPolicy()

	p.MaxLength = 16    // Maximum total length
	p.MinLength = 16    // Minimum total length
	p.MinDigits = 2     // Minimum digits
	p.MinSpclChars = 2  // Minimum special characters
	p.MinCapsAlpha = 1  // Minimum capital letters
	p.MinSmallAlpha = 1 // Minimum small letters

	password := gopassgen.Generate(p)
*/

import (
	"math/rand"
	"time"
)

// Policy of password to be passed in Generate() function
type Policy struct {
	MinLength      int    // Minimum length of password
	MaxLength      int    // Maximum length of password
	MinCapsAlpha   int    // Minimum length of capital letters
	MinSmallAlpha  int    // Minimum length of small letters
	MinDigits      int    // Minimum length of digits
	MinSpclChars   int    // Minimum length of special characters
	CapsAlphaPool  string // Permitted capital letters
	SmallAlphaPool string // Permitted small letters
	DigitPool      string // Permitted digits
	SpclCharPool   string // Permitted special characters
}

// NewPolicy returns a default password policy which can be modified
func NewPolicy() Policy {
	p := Policy{
		MinLength:      6,
		MaxLength:      16,
		MinCapsAlpha:   0,
		MinSmallAlpha:  0,
		MinDigits:      0,
		MinSpclChars:   0,
		CapsAlphaPool:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		SmallAlphaPool: "abcdefghijklmnopqrstuvwxyz",
		DigitPool:      "0123456789",
		SpclCharPool:   "!@#$%^&*()-_=+,.?/:;{}[]~",
	}
	return p
}

func random(min, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// CreateRandom returns a random byte string of given length from given byte string
func CreateRandom(bs []byte, length int) []byte {
	filled := make([]byte, length)
	max := len(bs)

	for i := 0; i < length; i++ {
		Shuffle(bs)
		filled[i] = bs[random(0, max)]
	}

	return filled
}

// Shuffle the given byte string
func Shuffle(bs []byte) {
	n := len(bs)
	for i := n - 1; i > 0; i-- {
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(i + 1)
		bs[i], bs[j] = bs[j], bs[i]
	}
}

// Generate a new password based on given policy
func Generate(p Policy) string {

	// Character length based policies should not be negative
	if p.MinLength < 0 || p.MaxLength < 0 || p.MinCapsAlpha < 0 ||
		p.MinSmallAlpha < 0 || p.MinDigits < 0 || p.MinSpclChars < 0 {
		panic("Character length should not be negative")
	}

	collectiveMinLength := p.MinCapsAlpha + p.MinSmallAlpha + p.MinDigits + p.MinSpclChars

	// Min length is the collective min length
	if collectiveMinLength > p.MinLength {
		p.MinLength = collectiveMinLength
	}

	// Max length should be greater than collective minimun length
	if p.MinLength > p.MaxLength {
		panic("Minimum length cannot be greater than maximum length")
	}

	if p.MaxLength == 0 {
		return ""
	}

	capsAlpha := []byte(p.CapsAlphaPool)
	smallAlpha := []byte(p.SmallAlphaPool)
	digits := []byte(p.DigitPool)
	spclChars := []byte(p.SpclCharPool)
	allChars := []byte(p.CapsAlphaPool + p.SmallAlphaPool + p.DigitPool + p.SpclCharPool)

	passwd := CreateRandom(capsAlpha, p.MinCapsAlpha)

	passwd = append(passwd, CreateRandom(smallAlpha, p.MinSmallAlpha)...)
	passwd = append(passwd, CreateRandom(digits, p.MinDigits)...)
	passwd = append(passwd, CreateRandom(spclChars, p.MinSpclChars)...)

	passLen := len(passwd)

	if passLen < p.MaxLength {
		randLength := random(p.MinLength, p.MaxLength)
		passwd = append(passwd, CreateRandom(allChars, randLength-passLen)...)
	}

	Shuffle(passwd)

	return string(passwd)
}
