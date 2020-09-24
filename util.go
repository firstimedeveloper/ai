package main

import "errors"

type pair struct {
	pID personID
	mID movieID
}

type Node struct {
	State  pair
	Parent pair
	Action pair
}

// Frontier Queue
type Frontier struct {
	frontier []Node
}

func (f *Frontier) Add(n Node) {
	f.frontier = append(f.frontier, n)
}

func (f *Frontier) Contains(p pair) bool {
	for _, n := range f.frontier {
		if n.State == p {
			return true
		}
	}
	return false
}

func (f *Frontier) Empty() error {
	if len(f.frontier) == 0 {
		return errors.New("frontier is empty")
	}
	return nil
}

func (f *Frontier) Peek() (pair, error) {
	if f.Empty() != nil {
		return pair{}, f.Empty()
	}
	return f.frontier[0].State, nil
}

func (f *Frontier) Remove() error {
	if f.Empty() != nil {
		return f.Empty()
	}
	f.frontier = f.frontier[1:]
	return nil
}
