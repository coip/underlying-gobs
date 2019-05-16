# stimulatorgator
### chomping on the _gobs_ over `stdin` and `stdout`

first, `go get -v github.com/coip/underlying-gobs`

_(there is an error about no .go files in the toplevel of the project, that is okay.)_

---

then build `callee`:

 `cd $GOPATH/src/github.com/coip/underlying-gobs/callee && go build`

---

then build `caller`:

 `cd $GOPATH/src/github.com/coip/underlying-gobs/caller && go build`

---

then invoke `$GOPATH/src/github.com/coip/underlying-gobs/caller/caller 3 "i like avocados"`


--- 

to have some `callee` Gophers make Gophers to send the `caller` Gopjhers some Gophers over `gob`'d stdin/stdout, and unGob them in caller into a Gopher literal which calls the lib func `Describe` in the `caller`.

## ( ͡~ ͜ʖ ͡°)
