package tree

type Tree struct {
	left *Tree
	right *Tree
	value int
}

func New(value int, left *Tree, right *Tree) *Tree {
	return &Tree{ left, right, value}
}

func (t *Tree) Value() int {
	return t.value
}

func (t *Tree) Add(value int) {
	if t==nil {
		t=New(value,nil,nil)
	}
	switch {
	case value<t.value:
		if t.left==nil {
			t.left=New(value,nil,nil)
		} else {
			t.left.Add(value)
		}
	default:
		if t.right==nil {
			t.right=New(value,nil,nil)
		} else {
			t.right.Add(value)
		}
	}
}


func (t *Tree) Find(value int) *Tree {
	if t==nil {
		return nil
	}
	switch {
	case value==t.value:
		return t
	case value<t.value:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

func (t *Tree) PreOrderTraversal() []int {
	s:=[]int{}
	if t==nil {
		return s
	}
	s=append(s,t.left.PreOrderTraversal()...)
	s=append(s,t.value)
	s=append(s,t.right.PreOrderTraversal()...)
	return s
}

func (t *Tree) InOrderTraversal() []int {
	s:=[]int{}
	if t==nil {
		return s
	}
	s=append(s,t.value)
	s=append(s,t.left.InOrderTraversal()...)
	s=append(s,t.right.InOrderTraversal()...)
	return s
}

func (t *Tree) PostOrderTraversal() []int {
	s:=[]int{}
	if t==nil {
		return s
	}
	s=append(s,t.left.PostOrderTraversal()...)
	s=append(s,t.right.PostOrderTraversal()...)
	s=append(s,t.value)
	return s
}