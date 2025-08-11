# go-reloaded

## Description
`go-reloaded` is a Go-based text processing tool designed to perform specific text modifications on an input file and write the results to an output file. The tool processes text transformations such as converting hexadecimal and binary numbers to decimal, modifying word cases (uppercase, lowercase, capitalized), handling punctuation spacing, and adjusting articles ("a" to "an") based on vowel rules.

This project is part of the curriculum at Tomorrow School, aimed at reinforcing Go programming skills, file handling, string manipulation, and unit testing. The tool is designed to follow good coding practices and includes unit tests for robust functionality.

## Features
- **Number Conversion**:
  - Replaces hexadecimal numbers marked with `(hex)` with their decimal equivalents (e.g., `42 (hex)` → `66`).
  - Replaces binary numbers marked with `(bin)` with their decimal equivalents (e.g., `10 (bin)` → `2`).
- **Case Modification**:
  - Converts words before `(up)` to uppercase (e.g., `go (up)` → `GO`).
  - Converts words before `(low)` to lowercase (e.g., `SHOUTING (low)` → `shouting`).
  - Capitalizes words before `(cap)` (e.g., `bridge (cap)` → `Bridge`).
  - Supports modifying multiple words with `(up, <number>)`, `(low, <number>)`, or `(cap, <number>)` (e.g., `so exciting (up, 2)` → `SO EXCITING`).
- **Punctuation Handling**:
  - Ensures single punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) are close to the previous word with a space before the next word (e.g., `there ,and` → `there, and`).
  - Preserves groups of punctuation like `...` or `!?` (e.g., `thinking ...` → `thinking...`).
  - Formats quotation marks (`'`) around words without spaces (e.g., `' awesome '` → `'awesome'`).
- **Article Adjustment**:
  - Converts `a` to `an` before words starting with a vowel (a, e, i, o, u) or `h` (e.g., `a amazing` → `an amazing`).

## Installation
1. Ensure you have Go installed (version 1.16 or later recommended).
2. Clone the repository:
   ```bash
   git clone https://01.tomorrow-school.ai/git/azzz/go-reloaded
   ```
3. Navigate to the project directory:
   ```bash
   cd go-reloaded
   ```

## Usage
Run the program with an input file and an output file as arguments:
```bash
go run . <input_file> <output_file>
```

### Example
**Input (`sample.txt`)**:
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
```

**Command**:
```bash
go run . sample.txt result.txt
```

**Output (`result.txt`)**:
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

### Additional Examples
**Input (`sample.txt`)**:
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
**Output (`result.txt`)**:
```
Simply add 66 and 2 and you will see the result is 68.
```

**Input (`sample.txt`)**:
```
Punctuation tests are ... kinda boring ,what do you think ?
```
**Output (`result.txt`)**:
```
Punctuation tests are... kinda boring, what do you think?
```

## Project Structure
- `main.go`: Entry point of the program, handles file input/output and orchestrates text processing.
- `processor.go`: Contains the core text processing logic for transformations.
- `*_test.go`: Unit test files for validating functionality.
- `README.md`: This file, providing project documentation.

## Development
- The project uses only standard Go packages, ensuring compatibility and simplicity.
- Follow Go best practices for clean, maintainable code.
- Write unit tests in files named `*_test.go` to verify each transformation function.
- Test your code locally with various input files to ensure robustness.

### Running Tests
To run unit tests:
```bash
go test ./...
```

## Contributing
This project is audited by peers at Tomorrow School. When contributing:
1. Create your own test cases to validate your code and for auditing others.
2. Ensure code adheres to Go style guidelines (use `gofmt` and `golint`).
3. Submit changes via pull requests to the repository: [https://01.tomorrow-school.ai/git/azzz/go-reloaded](https://01.tomorrow-school.ai/git/azzz/go-reloaded).

## Learning Objectives
- Master Go file system (fs) API for reading and writing files.
- Develop skills in string and number manipulation.
- Practice writing unit tests for robust code.
- Understand peer code review and auditing processes.

## License
This project is for educational purposes as part of the Tomorrow School curriculum.