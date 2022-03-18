package cfg

import "fmt"

var i = 0

type CFG[T any] struct {
	Nodes []*T
}

func (cfg CFG[T]) Print(p func(*T)) {
	fmt.Printf("%v: \n", i)
	for _, v := range cfg.Nodes {
		p(v)
	}
	i++
	fmt.Println("")
}
