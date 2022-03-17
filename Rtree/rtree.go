
package rtree

type Numerics interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64
}

type rect[T Numerics] struct {
	start	[]T
	end 	[]T
}

type rnode[T Numerics, OBJ any] struct {
	rect	 rect[T] 
	children []*rnode[T, OBJ]
	object	 *OBJ
}

type Rtree[T Numerics, OBJ any] struct {
	root 	rnode[T, OBJ]
	dim 	uint
	max	 	uint
}

func New[T Numerics, OBJ any](dim uint, max uint) Rtree[T, OBJ]{

	noneRect := rect[T]{ start:make([]T, dim), end:make([]T, dim) }
	rootNode := rnode[T, OBJ]{ rect:noneRect, children:[]*rnode[T, OBJ]{}, object:nil }

	return Rtree[T, OBJ]{ root:rootNode, dim:dim, max:max }
}

func (rnode *rnode[T, OBJ]) isObject() bool{
	return rnode.object != nil
}

func (rnode *rnode[T, OBJ]) isInsideOf(targnode *rnode[T, OBJ]) bool {
	return rnode.rect.isInsideOf(&targnode.rect)
}

func (rnode *rnode[T, OBJ]) isOutsideOf(targnode *rnode[T, OBJ]) bool {
	return rnode.rect.isOutsideOf(&targnode.rect)
}

func (rrect *rect[T]) isInsideOf(targrect *rect[T]) bool{
	for i:=0; i < len(rrect.start); i++{
		if rrect.start[i] < targrect.start[i] || rrect.end[i] > targrect.end[i]{
			return false
		}
	}
	return true
}

func (rrect *rect[T]) isOutsideOf(targrect *rect[T]) bool{
	for i:=0; i < len(rrect.start); i++{
		if rrect.end[i] < targrect.start[i] || targrect.end[i] < rrect.start[i]{
			return true
		}
	}
	return false
}


func min[T Numerics](a, b T) T{
	if a < b { return a }
	return b
}

func max[T Numerics](a, b T) T{
	if a > b { return a }
	return b 
}


func (rnode *rnode[T, OBJ]) getSizeExpandedBy(givennode *rnode[T, OBJ]) T{
	var expanded T = 1
	for i := 0; i < len(rnode.rect.start); i++ {
		expanded *= max(rnode.rect.end[i], givennode.rect.end[i]) - min(rnode.rect.start[i], givennode.rect.start[i]) 
	}
	return expanded
}


func (rnode *rnode[T, OBJ]) getSize() T{
	var size T = 1
	for i := 0; i < len(rnode.rect.start); i++ {
		size *= rnode.rect.end[i] - rnode.rect.start[i]
	}
	return size
}

func expandedRect[T Numerics](a, b *rect[T]) rect[T]{
	dim := len(a.start)

	rect := rect[T]{ start:make([]T, dim), end:make([]T, dim) }

	for i := 0; i < dim; i++ {
		rect.start[i] = min(a.start[i], b.start[i]) 
		rect.end[i]   = max(a.end[i],   b.end[i])
	}

	return rect
}

func (node *rnode[T, OBJ]) store(objnode *rnode[T, OBJ], max uint){

	if node.isObject() == true {
		newnode := *node
		*node = rnode[T, OBJ]{
			rect	: expandedRect(&node.rect, &objnode.rect),
			children: []*rnode[T, OBJ]{&newnode, objnode},
			object	: nil,
		}
	} else {

		if uint(len(node.children)) < max {
			node.children = append(node.children, objnode)
		} else {
			min_expantion_size := T(0)
			min_size := T(0)
			node_ad := -1
		
			for i := 0; i<len(node.children); i++ {
				size := node.children[i].getSize()
				expantion_size := node.children[i].getSizeExpandedBy(objnode) - size
				if expantion_size < 0 { expantion_size = 0 }
				chosen := false
				if min_expantion_size > expantion_size || node_ad == -1{
					chosen = true
				} else if min_expantion_size == expantion_size && min_size > size{
					chosen = true
				}

				if chosen {
					node_ad = i
					min_expantion_size = expantion_size
					min_size = size
				}
			}

			node.children[node_ad].store(objnode, max)

		}

		node.rect = expandedRect(&node.rect, &objnode.rect)

	}

}

func (node *rnode[T, OBJ]) getDepth() int{
	s := 0
	for i := range node.children{
		s = max(s, node.children[i].getDepth()+1)
	}
	return s
}

func (node *rnode[T, OBJ]) isPoint() bool{

	for i := range node.rect.start{
		if node.rect.start[i] != node.rect.end[i] {
			return false
		}
	}
	return true
}

func (rtree *Rtree[T, OBJ]) Store(start []T, end []T, obj OBJ){

	if uint(len(start)) != rtree.dim || uint(len(end)) != rtree.dim {
		panic("NO!")
	}

	objrect := rect[T]{ start:make([]T, rtree.dim), end:make([]T, rtree.dim) }

	for i := 0; i < int(rtree.dim); i++ {
		objrect.start[i], objrect.end[i]  = min(start[i], end[i]), max(start[i], end[i])
	}

	objnode := rnode[T, OBJ]{
		rect: objrect, 
		children : nil,
		object : &obj,
	}

	if len(rtree.root.children) == 0 {
		rtree.root.rect = objnode.rect
	}

	rtree.root.store(&objnode, rtree.max)

}

