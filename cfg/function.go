package cfg

import "github.com/numacci/go-algorithm/datastruct"

const stackCap = 10000

type Function[T any] struct {
	CFGs []*CFG[T]
}

func NewFunction[T any](hs []*T, preds func(*T) []*T, succs func(*T) []*T) Function[T] {
	workList := datastruct.NewStack(stackCap)
	visited := datastruct.NewSet()

	for _, v := range hs {
		workList.Push(v)
	}

	cfgs := make([]*CFG[T], 0)

	for !workList.IsEmpty() {
		b := make([]*T, 0)

		s := workList.Top().(*T)
		workList.Pop()
		if visited.Contains(s) {
			continue
		}

		for {
			visited.Insert(s)
			b = append(b, s)

			succ := succs(s)
			if len(succs(s)) != 1 || len(preds(succ[0])) > 1 {
				for _, v := range succs(s) {
					workList.Push(v)
				}
				break
			}

			s = succs(s)[0]
		}

		cfgs = append(cfgs, &CFG[T]{Nodes: b})
	}

	return Function[T]{CFGs: cfgs}
}

// TODO:
func (f Function[T]) GetFromHead() {

}
