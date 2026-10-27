# sunasty-color-math-go

Zero-dependency Go package for sRGB, OKLCH, and WCAG 2.1 accessibility contrast calculation.

Originally implemented for the [Sunasty Web Suite](https://sunasty.com/), powering the [Sunasty Color Picker](https://sunasty.com/color-picker/), [WCAG Contrast Checker](https://sunasty.com/contrast-checker/), and [Sunasty Developer Utilities](https://sunasty.com/dev/).

## Installation

```bash
go get github.com/bsulyan/sunasty-color-math-go@v1.0.2
```

## Documentation

- **Go API Reference**: [pkg.go.dev/github.com/bsulyan/sunasty-color-math-go](https://pkg.go.dev/github.com/bsulyan/sunasty-color-math-go)
- **Architecture & Multi-Language Docs**: [sunasty-docs.readthedocs.io](https://sunasty-docs.readthedocs.io/en/latest/)
- **Mathematical Deep Dive (OKLCH & WCAG)**: [bsulyan.github.io/notes/oklch-and-wcag-contrast-math/](https://bsulyan.github.io/notes/oklch-and-wcag-contrast-math/)

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

## Cross-Language Ecosystem (`sunasty-color-math`)

- **Rust**: [docs.rs/sunasty-color-math](https://docs.rs/sunasty-color-math/latest/sunasty_color_math/) · [crates.io/crates/sunasty-color-math](https://crates.io/crates/sunasty-color-math)
- **Python**: [pypi.org/project/sunasty-color-math/](https://pypi.org/project/sunasty-color-math/)
- **PHP**: [packagist.org/packages/sunasty/color-math](https://packagist.org/packages/sunasty/color-math)
- **Dart / Flutter**: [pub.dev/packages/sunasty_color_math](https://pub.dev/packages/sunasty_color_math) ([publisher: sunasty.com](https://pub.dev/publishers/sunasty.com))
- **JavaScript / TypeScript**: [npmjs.com/package/sunasty-color-math](https://www.npmjs.com/package/sunasty-color-math)

## Companion Web Tools & Browser Extensions

- [Sunasty Color Picker & Palette Extractor](https://sunasty.com/color-picker/)
- [Sunasty WCAG Contrast Checker](https://sunasty.com/contrast-checker/)
- [Sunasty Developer Utilities](https://sunasty.com/dev/)
- [Sunasty QR Scanner — Chrome Web Store](https://chromewebstore.google.com/detail/sunasty-qr-code-scanner-g/fbiighpklpogdbnhdapbiijcadhlijcb) · [Firefox Add-ons (AMO)](https://addons.mozilla.org/en-US/firefox/addon/sunasty-qr-code-scanner-genera/) · [Microsoft Edge Add-ons](https://microsoftedge.microsoft.com/addons/detail/sunasty-qr-code-scanner-generator/npbppaegpbeefpfaffjgphjecmgbngab)
- [Sunasty on AlternativeTo](https://alternativeto.net/software/sunasty/about/)
- [Ulysan Studio](https://ulysan.fr/)

## License

MIT © [Bastien Sulyan](https://bsulyan.github.io/)
