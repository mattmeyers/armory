# Armory

Armory is a CLI tool that can generate typed data structures on the fly.

## Installation

```
go get -u github.com/mattmeyers/armory/cmd/armory/...
```

## Usage

```
armory - A CLI data structure generator
  armory [options] <data structure>

data structures:
  set
  stack

options:
  --out string, -o string   the file to write to
  --pkg string, -p string   the file's package (default "main")
  --type string, -t string  the data structure's type (default "int")
  --zero string, -z string  the type's zero value, inferred by default
  --help, -h                display this help message
```

## License

MIT License

Copyright (c) 2020 Matthew Meyers

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
