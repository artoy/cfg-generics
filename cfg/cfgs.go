package cfg

import "github.com/numacci/go-algorithm/datastruct"

const stackCap = 10000

func NewCFGs[T any](hs []*T, succs func(*T) []*T, preds func(*T) []*T) []*CFG[T] {
	workList := datastruct.NewStack(stackCap)
	visited := datastruct.NewSet()

	for _, v := range hs {
		workList.Push(v)
	}

	cfgs := make([]*CFG[T], 0)

	for !workList.IsEmpty() {
		b := make([]*T, 0)

		s := workList.Top().(*T)
		if visited.Contains(s) {
			continue
		}

		for {
			visited.Insert(s)
			b = append(b, s)

			if len(succs(s)) != 1 || len(preds(succs(s)[0])) > 1 {
				for _, v := range succs(s) {
					workList.Push(v)
				}
				break
			}

			s = succs(s)[0]
		}

		cfgs = append(cfgs, &CFG[T]{Nodes: b})
	}

	return cfgs
}
