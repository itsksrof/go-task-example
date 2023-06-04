package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const ARGS_LENGTH = 3

var bufSize = 2 << 20 // Default buffer size is roughly 2MBs

var commands = func() map[string]func(v ...any) string {
	return map[string]func(v ...any) string{
		"--help": help,
		"-h":     help,
		"--out":  out,
		"-o":     out,
	}
}

func help(params ...any) string {
	return `Usage: grab [OPTION] ... [FILE] ...
Grab FILE(s) content and show it through standard output

--help, -h	display this help and exit
--out, -o	output the file content to stdout

Examples:

grab -o example.json
grab -o example.json example.txt`
}

func out(params ...any) string {
	buffer := bytes.NewBuffer(make([]byte, bufSize))

	for _, param := range params {
		fileName, _ := param.(string)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		_, err = reader.Read(buffer.Bytes())
		if err != nil {
			panic(err)
		}
	}

	return buffer.String()
}

func main() {
	if val, ok := os.LookupEnv("BUF_SIZE"); ok {
		n, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		bufSize = n << 20
	}

	cmd, ok := commands()[os.Args[1]]
	if !ok {
		panic("command not found, try --help to list the available commands")
	}

	if len(os.Args) >= ARGS_LENGTH {
		for _, arg := range os.Args[2:] {
			fmt.Println(cmd(arg))
		}
	} else {
		fmt.Println(cmd())
	}
}
