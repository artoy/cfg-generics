package cfg

import "fmt"

type CFG[T any] struct {
	Nodes []*T
}

func (cfg CFG[T]) Print(p func(*T)) {
	for i, v := range CFG[T].Nodes {
		fmt.Printf("%v: \n", i)
		p(v)
		fmt.Println("")
	}
}
