# gopassgen

Policy based password generator based on GO

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
