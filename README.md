# gopassgen

[![GoDoc](https://godoc.org/github.com/sayanarijit/gopassgen?status.svg)](https://godoc.org/github.com/sayanarijit/gopassgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/sayanarijit/gopassgen)](https://goreportcard.com/report/github.com/sayanarijit/gopassgen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/sayanarijit/gopassgen/blob/master/LICENSE)

To use it as a command-line tool or web app, install [gopgcli](https://github.com/sayanarijit/gopgcli) or [gopgweb](https://github.com/sayanarijit/gopgweb)

### Install

```bash
go get github.com/sayanarijit/gopassgen
```

### Available policy options

| Policy                       | Configuration  | Type    | Default                      |
| ---------------------------- | -------------  | ------- | ---------------------------- |
| Maximum length               | MaxLength      | int     | 16                           |
| Minimum length               | MinLength      | int     | 6                            |
| Minimum capital letters      | MinCapsAlpha   | int     | 0                            |
| Minimum small letters        | MinSmallAlpha  | int     | 0                            |
| Minimum digits               | MinDigits      | int     | 0                            |
| Minimum special characters   | MinSpclChars   | int     | 0                            |
| Permitted capital letters    | CapsAlphaPool  | string  | `ABCDEFGHIJKLMNOPQRSTUVWXYZ` |
| Permitted small letters      | SmallAlphaPool | string  | `abcdefghijklmnopqrstuvwxyz` |
| Permitted digits             | DigitPool      | string  | `0123456789`                 |
| Permitted special characters | SpclCharPool   | string  | `!@#$%^&*()-_=+,.?/:;{}[]~`  |

### Example Usage

#### Generate a 16 character long password with minimum 2 digits, 2 special characters, 1 capital and 1 small letter

```go
package main

import (
    "fmt"
    "github.com/sayanarijit/gopassgen"
)

func main() {

    p := gopassgen.NewPolicy()

    p.MaxLength = 16      // Maximum total length
    p.MinLength = 16      // Minimum total length
    p.MinDigits = 2       // Minimum digits
    p.MinSpclChars = 2    // Minimum special characters
    p.MinCapsAlpha = 1    // Minimum capital letters
    p.MinSmallAlpha = 1   // Minimum small letters

    password := gopassgen.Generate(p)

    fmt.Println(password)
}
```

#### Quickly generate random password of given length using given characters

```go
package main

import (
    "fmt"
    "github.com/sayanarijit/gopassgen"
)

func main() {

    bsPassword := gopassgen.CreateRandom([]byte("ABCDwxyz1234$%^&"), 8) // Returns bytes array

    fmt.Println(string(bsPassword))
}
```

#### Generate password by shuffling given characters

```go
package main

import (
    "fmt"
    "github.com/sayanarijit/gopassgen"
)

func main() {

    bsPassword := []byte("ABCDwxyz1234$%^&")

    gopassgen.Shuffle(bsPassword)

    fmt.Println(string(bsPassword))
}
```
