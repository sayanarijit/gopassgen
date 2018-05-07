package gopassgen

/***
Version		: 0.1.0
Author		: Arijit Basu <sayanarijit@gmail.com>
Docs		: https://github.com/sayanarijit/gopassgen#README

USAGE:
	p := gopassgen.NewPolicy()

	p.MinDigits = 5
	p.MinCapsAlpha = 2
	p.MinSpclChars = 2

	password := gopassgen.Generate(p)
***/

import (
	"math/rand"
	"time"
)

type Policy struct {
	MinLength     int
	MaxLength     int
	MinCapsAlpha  int
	MinSmallAlpha int
	MinDigits     int
	MinSpclChars  int
}

func NewPolicy() Policy {
	p := Policy{
		MinLength:     6,
		MaxLength:     16,
		MinCapsAlpha:  0,
		MinSmallAlpha: 0,
		MinDigits:     0,
		MinSpclChars:  0,
	}
	return p
}

func Shuffle(s []byte) {
	rand.Seed(time.Now().UnixNano())
	n := len(s)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Generate(p Policy) string {
	if p.MaxLength > 0 {
		if p.MinCapsAlpha+p.MinSmallAlpha+p.MinDigits+p.MinSpclChars > p.MaxLength {
			panic("Max length is not sufficient")
		}
	}

	CapsAlpha := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	SmallAlpha := []byte("abcdefghijklmnopqrstuvwxyz")
	Digits := []byte("0123456789")
	SpclChars := []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")
	AllChars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	Shuffle(CapsAlpha)
	Shuffle(SmallAlpha)
	Shuffle(Digits)
	Shuffle(SpclChars)

	passwd := CapsAlpha[:p.MinCapsAlpha]

	if p.MinSmallAlpha > 0 {
		passwd = append(passwd, SmallAlpha[:p.MinSmallAlpha]...)
	}
	if p.MinDigits > 0 {
		passwd = append(passwd, Digits[:p.MinDigits]...)
	}
	if p.MinSpclChars > 0 {
		passwd = append(passwd, SpclChars[:p.MinSpclChars]...)
	}

	if len(passwd) < p.MinLength {
		requiredMore := p.MinLength - len(passwd)
		Shuffle(AllChars)
		passwd = append(passwd, AllChars[:requiredMore]...)
	}

	Shuffle(passwd)

	return string(passwd)
}
