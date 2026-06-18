# Go Reloaded

A collection of Go programming exercises focused on string manipulation, text processing, and input handling. Each drill demonstrates different techniques for parsing and transforming text input in Go.

## Project Structure

```
go-reloaded/
├── main.go                                    # Main project - binary to decimal conversion
├── drill-2-convert-hex/                       # Convert hexadecimal to decimal
├── drill-3-convert-bin/                       # Convert binary to decimal
├── drill-4-fix-punctuation/                   # Attach punctuation to preceding words
├── dril-5-multi-word/                         # Apply case modifiers (up/low/cap)
└── drill-7-Fix Single-Quote-Formatting/       # Fix spacing around single quotes
```

## Drills Overview

### Main Project
Converts binary numbers to decimal. Input format: `<binary_number> (bin)`

Example:
```bash
echo "1010 (bin)" | go run main.go
# Output: 10
```

### Drill 2: Convert Hexadecimal
Converts hexadecimal numbers to decimal format using the `[hex]` tag.

**Key Function:** `convertHex()`

Example:
```bash
echo "FF [hex]" | go run drill-2-convert-hex/main.go
# Output: 255 (removed [hex] tag)
```

**Features:**
- Parses hexadecimal strings
- Uses `strconv.ParseInt()` with base 16
- Removes the `[hex]` tag from output

### Drill 3: Convert Binary
Converts binary numbers to decimal format using the `[bin]` tag.

**Key Function:** `convertBin()`

Example:
```bash
echo "1111 [bin]" | go run drill-3-convert-bin/main.go
# Output: 15
```

**Features:**
- Parses binary strings
- Uses `strconv.ParseInt()` with base 2
- Removes the `[bin]` tag from output

### Drill 4: Fix Punctuation
Attaches standalone punctuation marks to the preceding word.

**Key Function:** `fixPunctuation()`

Example:
```bash
echo "Hello , world !" | go run drill-4-fix-punctuation/main.go
# Output: Hello, world!
```

**Supported Punctuation:** `! ? , ; : .`

**Features:**
- Removes spaces between words and punctuation
- Uses `strings.ContainsAny()` for punctuation detection
- Preserves word order and content

### Drill 5: Multi-Word Case Modifiers
Applies case transformations to words using tags like `[up]`, `[low]`, `[cap]` with optional word counts.

**Key Function:** `applyCaseModifiers()`

Example:
```bash
echo "hello world [up,2]" | go run dril-5-multi-word/main.go
# Output: HELLO WORLD
```

**Supported Modifiers:**
- `[up]` - Convert last word to UPPERCASE
- `[low]` - Convert last word to lowercase
- `[cap]` - Capitalize last word
- `[up,n]`, `[low,n]`, `[cap,n]` - Apply to last n words

**Features:**
- Parses modifier tags with optional counts
- Applies transformations to specified number of preceding words
- Removes tags from output

### Drill 7: Fix Single-Quote Formatting
Removes unwanted spaces around single quotes using regular expressions.

**Key Function:** `fixSingleQuotes()`

Example:
```bash
echo "She ' s here" | go run "drill-7-Fix Single-Quote-Formatting/main.go"
# Output: She's here
```

**Features:**
- Removes spaces after opening single quotes: `' ` → `'`
- Removes spaces before closing single quotes: ` '` → `'`
- Uses `regexp.MustCompile()` for pattern matching
- Uses anonymous functions for replacements

## Key Go Concepts Used

- **Input Handling:** `bufio.Scanner` for reading multi-line input
- **String Processing:** `strings` package functions (Fields, Split, Join, etc.)
- **Type Conversion:** `strconv` package for base conversions (binary, hexadecimal, decimal)
- **Regular Expressions:** `regexp` package for pattern matching
- **Text Parsing:** Identifying and processing special markers and tags

## Running the Programs

Each drill can be run independently:

```bash
# Navigate to the drill directory
cd drill-2-convert-hex

# Run with input
echo "FF [hex]" | go run main.go
```

Or build and run as an executable:

```bash
go build -o converter main.go
echo "FF [hex]" | ./converter
```

## Learning Outcomes

Through these drills, you'll practice:

1. ✅ Reading and processing standard input
2. ✅ String tokenization and manipulation
3. ✅ Number base conversions
4. ✅ Pattern matching and validation
5. ✅ Text transformation algorithms
6. ✅ Error handling with `strconv` conversions
7. ✅ Regular expressions in Go
8. ✅ Functional programming patterns (anonymous functions)

## Requirements

- Go 1.16 or higher
- Standard Go library (no external dependencies)

## License

This is an educational project for learning Go programming concepts.
