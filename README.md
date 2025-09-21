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
mdconverter convert input.md -o output.pdf
```

Convert Markdown → HTML:

```bash
mdconverter convert input.md -o output.html
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
mdconverter convert notes.md -o notes.pdf
```

Markdown → HTML:

```bash
mdconverter convert report.md -o report.html
```

---

## Build from source

```bash
git clone https://github.com/Zizu-oswald/MdConverter.git
cd MdConverter
go build -o mdconverter
```

---
## Adding to PATH

To make `MdConverter` available as a global command, add the binary to your system `PATH`.

---

###  Linux

#### Option 1: Install system-wide

```bash
sudo cp mdconverter /usr/local/bin/
```

#### Option 2: User-only installation

```bash
mkdir -p $HOME/bin
cp mdconverter $HOME/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Now you can run:

```bash
mdconverter -h
```

---

### Windows

1. Copy `mdconverter.exe` to a folder, e.g.:

   ```
   C:\tools\mdconverter\
   ```

2. Open:

   * **Win + R → sysdm.cpl**
   * Tab **Advanced → Environment Variables**
   * In **System variables**, find and edit `Path`
   * Add:

     ```
     C:\tools\mdconverter
     ```

3. Restart PowerShell or CMD.

Now you can run:

```powershell
mdconverter -h
```
