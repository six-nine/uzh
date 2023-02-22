package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/six-nine/uzh/lexer"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fileName := argsWithoutProg[0]

	l := lexer.New()

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf := make([]byte, 4*1024) //the chunk size
		n, err := r.Read(buf)       //loading chunk into buffer
		buf = buf[:n]
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
		}
		for i := 0; i < n; i++ {
			l.AddChar(buf[i])
		}
	}

	tokens := l.ExtractTokens()

	for _, tok := range tokens {
		fmt.Println(tok.Type, "\t", tok.Literal)
	}
}
