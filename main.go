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

var commands = func() map[string]func(v ...any) {
	return map[string]func(v ...any){
		"--help": help,
		"-h":     help,
		"--out":  out,
		"-o":     out,
	}
}

func help(params ...any) {
	fmt.Println(`
Usage: grab [OPTION] ... [FILE] ...
Grab FILE(s) content and show it through standard output

--help, -h	display this help and exit
--out, -o	output the file content to stdout

Examples:

grab -o example.json
grab -o example.json example.txt
	`)
}

func out(params ...any) {
	for _, param := range params {
		fileName, _ := param.(string)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		buffer := bytes.NewBuffer(make([]byte, bufSize))
		_, err = reader.Read(buffer.Bytes())
		if err != nil {
			panic(err)
		}

		fmt.Println("Filename:", file.Name())
		fmt.Println(buffer.String())
	}
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
			cmd(arg)
		}
	} else {
		cmd()
	}
}
