# Go-Reloaded

A text completion and auto-correction tool written in Go. This tool processes a text file, applies specific formatting rules (hex/bin conversion, casing, and punctuation), and outputs the result to a new file.

## Features

- **Hex/Binary Conversion**: Converts `(hex)` or `(bin)` tags to their decimal equivalents.
- **Case Manipulation**: 
  - Supports `(up)`, `(low)`, and `(cap)` for single words.
  - Supports multi-word versions like `(up, 2)`, `(low, 3)`, etc.
- **Punctuation Formatting**: 
  - Standardizes spacing for `.,!?;:`.
  - Handles groups of punctuation like `...` or `!?`.
- **Quote Formatting**: Properly places single quotes `'` around words or phrases without internal spacing.
- **Article Correction**: Automatically changes `a` to `an` before words starting with vowels or the letter `h`.

## Installation

Ensure you have [Go](https://go.dev) installed on your machine.

1. Clone the repository:
   ```bash
   git clone <your-repo-link>
   cd go-reloaded
