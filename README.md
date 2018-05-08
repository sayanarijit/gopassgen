# gopassgen

[![GoDoc](https://godoc.org/github.com/sayanarijit/gopassgen?status.svg)](https://godoc.org/github.com/sayanarijit/gopassgen)

Policy based password generator library for Golang

### Install

```bash
go get github.com/sayanarijit/gopassgen
```

### Available policy options

| Policy option              | Variable      | Type | Default |
| -------------------------- | ------------- | ---- | ------- |
| Minimum length             | MinLength     | int  | 6       |
| Maximum length             | MaxLength     | int  | 16      |
| Minimum capital letters    | MinCapsAlpha  | int  | 0       |
| Minimum small letters      | MinSmallAlpha | int  | 0       |
| Minimum digits             | MinDigits     | int  | 0       |
| Minimum special characters | MinSpclChars  | int  | 0       |

### Example Usage

* Generate a random password with 10 minimum digits and 5 special characters

```go
package main

import (
    "fmt"
    "github.com/sayanarijit/gopassgen"
)

func main() {

    p := gopassgen.NewPolicy()

    p.MinDigits = 10                     // Minimum digits
    p.MinSpclChars = 5                   // Minimum special characters

    password := gopassgen.Generate(p)

    fmt.Println(password)
}
```

* Quickly generate random password of given length using given characters

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

* Generate password by shuffling given characters

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
