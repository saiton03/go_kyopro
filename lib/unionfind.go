package kyopro

type UnionFind struct {
	Par []int
}

func NewUf(N int) *UnionFind {
	ret := &UnionFind{
		Par: make([]int, N),
	}
	return ret
}

func (uf *UnionFind) Root(x int) int {
	if uf.Par[x] == x {
		return x
	}
	uf.Par[x] = uf.Root(uf.Par[x])
	return uf.Par[x]
}

func (uf *UnionFind) Unite(x, y int) {
	rx := uf.Root(x)
	ry := uf.Root(y)
	if rx != ry {
		uf.Par[rx] = ry
	}
}

func (uf *UnionFind) SameRoot(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}
