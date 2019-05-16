package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/coip/underlying-gobs/lib/gopher"
)

func main() {
	popcount := os.Args[1]
	thoughtinput := os.Args[2:]

	args := strings.Join(thoughtinput, " ")
	c := exec.Command("../callee/callee", popcount)
	stdo, e := c.StdoutPipe()
	if e != nil {
		panic(e)
	}
	w, e := c.StdinPipe()
	if e != nil {
		panic(e)
	}
	n, e := w.Write([]byte(args + "\n"))
	fmt.Printf("%d bytes written! err: %+v\n", n, e)

	var goph gopher.Gopher

	e = c.Start()
	if e != nil {
		panic(e)
	}
	gd := gob.NewDecoder(stdo)
	if e != nil {
		panic(e)
	}
	tmpnum, e := strconv.Atoi(popcount)
	if e != nil || tmpnum < 1 {
		os.Stderr.Write([]byte("supply a real target population betch\n"))
		os.Exit(1)
	}
	for i := 0; i < tmpnum; i++ {
		e = gd.Decode(&goph)
		if e != nil {
			log.Printf("Err: %+v", e)
		}
		log.Println("got goph:", goph.Describe())
	}
	if e != nil {
		panic(e)
	}
	e = c.Wait()
	if e != nil {
		panic(e)
	}
	fmt.Printf("last goph: %+v\n", goph)
}
