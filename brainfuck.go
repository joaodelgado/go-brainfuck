package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	program          []rune
	tape             []uint8
	tapePos, codePos int
)

func main() {

	var reader *bufio.Reader

	flag.Parse()
	if len(flag.Args()) == 1 {
		input, err := ioutil.ReadFile(flag.Args()[0])
		program = bytes.Runes(input)
		if err != nil {
			fmt.Printf("Error reading file %s", flag.Args()[0])
			os.Exit(1)
		}
	} else {
		reader = bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		program = []rune(input)
	}

	if _, err := run(false); err != nil {
		fmt.Println(err)
	}
}

func run(skip bool) (bool, error) {
	for codePos < len(program) {
		if tapePos >= len(tape) {
			tape = append(tape, []uint8{0}...)
		}

		op := string(program[codePos])

		// if value on tape is 0, jump forward to command after matching ]
		if op == "[" {
			codePos++
			oldPos := codePos
			for {
				cont, err := run(tape[tapePos] == 0)
				if err != nil {
					return false, nil
				}
				if cont {
					codePos = oldPos
				} else {
					break
				}
			}
			// if value on tape is not 0, jump back to command after matching [
		} else if op == "]" {
			return tape[tapePos] != 0, nil
		} else if !skip {
			switch string(op) {
			// move right on tape
			case ">":
				tapePos++
			// move left on tape
			case "<":
				tapePos--
			// increment value on tape
			case "+":
				tape[tapePos]++
			// decrement value on tape
			case "-":
				tape[tapePos]--
			// output value on tape
			case ".":
				fmt.Printf("%s", string(tape[tapePos]))
			// input value on tape
			case ",":
				b := make([]uint8, 1)
				_, err := os.Stdin.Read(b)
				if err != nil {
					return false, fmt.Errorf("Error reading input")
				}
				tape[tapePos] = b[0]
			default:
				// discard
			}
		}
		codePos++
	}
	return false, nil
}
