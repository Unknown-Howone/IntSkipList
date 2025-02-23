package IntSkipList

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	r = rand.New(src)
}

const (
	maxLevel = 16 // 跳表最大层数
)

type Node struct {
	Val  int
	Next []*Node // 存放每一层的下一个节点
}

type Skiplist struct {
	Head  *Node
	Level int
}

func newNode(v, level int) *Node {
	return &Node{
		Val:  v,
		Next: make([]*Node, level),
	}
}

func New() Skiplist {
	return Skiplist{
		Head:  newNode(-1, maxLevel), // 这个地方物理头节点就一个,但是相对于每层都创建了逻辑头节点
		Level: 1,
	}
}

func (sl *Skiplist) Search(target int) bool {
	prev := sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		prev = closePrev(prev, i, target)
		if prev.Next[i] != nil && prev.Next[i].Val == target {
			return true
		}
	}
	return false
}

func (sl *Skiplist) Add(num int) {
	level := randomLevel()
	if level > sl.Level {
		sl.Level = level
	}
	node := newNode(num, level)
	prev := sl.Head
	for i := sl.Level - 1; i >= 0; i-- { // 也可以 for i := level - 1; i >= 0; i-- { 因为add操作时level是已知的
		prev = closePrev(prev, i, num)
		if i < level {
			node.Next[i] = prev.Next[i]
			prev.Next[i] = node
		}
	}
}

func (sl *Skiplist) Erase(num int) bool {
	foundNum := false
	prev := sl.Head
	for i := sl.Level - 1; i >= 0; i-- {
		prev = closePrev(prev, i, num)
		if prev.Next[i] != nil && prev.Next[i].Val == num {
			prev.Next[i].Next[i], prev.Next[i] = nil, prev.Next[i].Next[i]
			foundNum = true
		}
	}
	for sl.Level > 1 && sl.Head.Next[sl.Level-1] == nil {
		sl.Level--
	}
	return foundNum
}

// closePrev 找到在同一层当中，比num小的最大的节点(作为前驱)
func closePrev(prev *Node, level, num int) *Node {
	for prev.Next[level] != nil && prev.Next[level].Val < num {
		prev = prev.Next[level]
	}
	return prev
}

func randomLevel() int {
	level := 1
	for level < maxLevel && r.Float64() < 0.5 { // 一半的概率上升level
		level++
	}
	return level
}
