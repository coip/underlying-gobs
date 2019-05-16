package gopher

import (
	"fmt"
)

//Gopher is arbitrary
type Gopher struct {
	Iteration int32
	Name      string
	Vision    string
}

//MakeGopher returns a Gopher
func MakeGopher(name, seedthought string, iteration int32) Gopher {
	return Gopher{iteration, name, seedthought} //, make(chan string)}
}

//Describe is some func that belongs to a Gopher
func (g Gopher) Describe() string {
	return fmt.Sprintf("i am %s, the %dth of the mity gophs, and am driven with the single word of %s at mind!\n", g.Name, g.Iteration, g.Vision)
}
