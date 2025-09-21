# MdConverter

`MdConverter` is a simple CLI application in Go for converting Markdown files.

---

## Installation

### Clone the repository

```bash
git clone https://github.com/Zizu-oswald/MdConverter.git
```

### Download the binary

Go to [Releases](https://github.com/<your-username>/md2pdf/releases) and download the binary for your platform.

---

## Usage

Convert Markdown → PDF:

```bash
MdConverter convert input.md -o output.pdf
```

Convert Markdown → HTML:

```bash
MdConverter convert input.md -o output.html
```

### Options

| Flag | Description                                                    |
| ---- | -------------------------------------------------------------- |
| `-o` | Set the output file name (defaults to source file name + .pdf) |
| `-h` | Show help                                                      |

---

## Examples

Markdown → PDF:

```bash
MdConverter convert notes.md -o notes.pdf
```

Markdown → HTML:

```bash
MdConverter convert report.md -o report.html
```

---

## Build from source

```bash
git clone https://github.com/Zizu-oswald/MdConverter.git
cd MdConverter
go build .
```

---
