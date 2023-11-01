# go-lox
a byte-code Lox interpreter written in Go (Golang)

[![Go](https://github.com/Krawabbel/go-lox/actions/workflows/go.yml/badge.svg)](https://github.com/Krawabbel/go-lox/actions/workflows/go.yml)

[![Go Coverage](https://github.com/Krawabbel/go-lox/wiki/coverage.svg)](https://raw.githack.com/wiki/Krawabbel/go-lox/coverage.html)

The Lox interpreter found in this repository is an adaptation of the bytecode VM implementation in C proposed by Robert Nystrom in his highly recommended book 'Crafting Interpreters'.

## Build and Run

* Clone the repository, e.g. with ```git clone https://github.com/Krawabbel/go-lox.git``` or simply download it [here](https://github.com/Krawabbel/go-lox/archive/refs/heads/main.zip).

* Build go-lox with ```go build```.

## Run Lox Programs

This interpreter features a REPL (read-evaluate-print-loop) and script-runner which can be run by running

Linux and Mac: ```./go-lox ["path/to/script.lox"]```

Windows: ```./go-lox ["path/to/script.lox"]```

If no path to a script is provided, the REPL will be started.

## Lox Cheat Sheet

Lox is a minimalistic programming language in the tradition of C.

Some notable peculiarities:
* All statements must be terminated by a ';'.
* All numbers are 32-bit floats.
* All statements in if-else-clauses, for- and while-loops must be surrounded by curly braces ('{' and '}'). 

### Hello World

```print "Hello, World";```

### Key Words (reserved)

* and
* class
* else 
* false
* for
* fun 
* if
* nil
* or
* print
* return
* super
* this
* true
* var
* while

## Resources

* Crafting Interpreters (TOC): http://craftinginterpreters.com/contents.html
* Crafting Interpreters on github: https://github.com/munificent/craftinginterpreters/tree/master
* other Lox implementations: https://github.com/munificent/craftinginterpreters/wiki/Lox-implementations
