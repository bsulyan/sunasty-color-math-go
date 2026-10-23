# sunasty-color-math-go

Zero-dependency Go package for sRGB, OKLCH, and WCAG 2.1 accessibility contrast calculation.

Originally implemented for the [Sunasty Color Picker](https://sunasty.com/color-picker/) client-side tool and the [Sunasty Developer Utilities](https://sunasty.com/dev/).

## Installation

```bash
go get github.com/bsulyan/sunasty-color-math-go
```

## Documentation

Full API documentation is hosted on [pkg.go.dev/github.com/bsulyan/sunasty-color-math-go](https://pkg.go.dev/github.com/bsulyan/sunasty-color-math-go).

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/bsulyan/sunasty-color-math-go"
)

func main() {
	teal, _ := colormath.ParseHex("#0F766E")
	white := colormath.RGB{R: 255, G: 255, B: 255}

	ratio := colormath.ContrastRatio(teal, white)
	wcag := colormath.CheckWCAG(ratio)

	fmt.Printf("Contrast: %.2f:1\n", ratio)
	fmt.Printf("Passes WCAG AA: %v\n", wcag.AANormal)
}
```

## Related Tools

- [Sunasty Color Picker & Palette Extractor](https://sunasty.com/color-picker/)
- [Sunasty Developer Utilities](https://sunasty.com/dev/)
- [Ulysan](https://ulysan.fr)

## License

MIT © [Bastien Sulyan](https://ulysan.fr)
