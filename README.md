# lc3-disassembler

A simple CLI tool that disassembles LC-3 binary instructions.

##### Disclaimer: This tool was made in preparation for an exam at DTU. Intended for educational purposes only.

---

## ✨ Features

- Disassemble LC-3 binary instructions to LC-3 assembly
- that's it ¯\\_( ͡° ͜ʖ ͡°)\_/¯

---

## 📦 Installation

1. Make sure you have [Go](https://golang.org/doc/install) installed.
2. Clone the repository.
3. Build into an executable `go build -o lc3d.exe`
4. Run the executable to generate the input file (will be generated as /data/input.txt).
5. Paste binary LC-3 instructions into the input file. One instruction per line.
6. Run the `disassemble` command. The output will be generated as /data/output.txt.

---

## 🧠 Usage

```bash
lc3d.exe [command]
```

🔧 Available Commands

- `disassemble` - Disassemble the LC-3 binary instructions in the input file and write the output to the output file.
- `help` - Show help message

---

📄 License
MIT © Birk Nordbo
