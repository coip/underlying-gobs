package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/coip/underlying-gobs/lib/gopher"

)

var (
	names = []string{
		"Mallory",
		"Willa",
		"Maurise",
		"Bernete",
		"Carly",
		"Evered",
		"Guss",
		"Jobyna",
		"Jena",
		"Stinky",
		"Sybila",
		"Angie",
		"Bryn",
		"Emmalyn",
	}
)

var population = int32(0)
var poplock sync.Mutex

// FROM https://blog.klauspost.com/defer-fun/
func Un(f func()) {
	f()
}

func Lock(x sync.Locker) func() {
	x.Lock()
	return func() { x.Unlock() }
}

///////////////////////////
func main() {
	seedthought, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	seedthought = seedthought[:len(seedthought)-1]
	gd := gob.NewEncoder(os.Stdout)

	targetPopulation := os.Args[1]
	tmpnum, e := strconv.Atoi(targetPopulation)
	if e != nil || tmpnum < 1 {
		os.Stderr.Write([]byte("supply a real target population betch\n"))
		os.Exit(1)
	}
	for i := 0; i < tmpnum; i++ {
		f := Lock(&poplock)
		g := gopher.MakeGopher(names[population], seedthought, population)
		atomic.AddInt32(&population, 1)
		Un(f)
		e := gd.Encode(g)
		if e != nil {
			log.Printf("Err: %+v", e)
		}
	}

	os.Exit(0)
}
