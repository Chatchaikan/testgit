package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Time Out")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	fmt.Println("Echo...")
	io.Copy(out, in)

}