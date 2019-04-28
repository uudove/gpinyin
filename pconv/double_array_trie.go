package pconv

import (
	"container/list"
	"reflect"
)

type Node struct {
	code  int
	depth int
	left  int
	right int
}

type DoubleArrayTrie struct {
	check []int
	base  []int

	used         []bool
	size         int
	allocSize    int
	key          []string
	keySize      int
	progress     int
	nextCheckPos int
	errno        int
}

func (this *DoubleArrayTrie) build(key []string) int {
	if key == nil {
		return 0
	}
	this.key = key
	this.keySize = len(key)
	this.progress = 0

	// resize array
	this.resize(65536)

	this.base[0] = 1
	this.nextCheckPos = 0

	rootNode := new(Node) // pointer
	rootNode.left = 0
	rootNode.right = this.keySize
	rootNode.depth = 0

	siblings := list.New()

	this.fetch(rootNode, siblings)
	this.insert(siblings)

	this.used = nil
	this.key = nil

	return this.errno
}

func (this *DoubleArrayTrie) commonPrefixSearch(key string) []int {
	keyChars := []rune(key)
	keyLength := len(keyChars)
	result := make([]int, 0)

	b := this.base[0]
	n := 0
	p := 0

	for i := 0; i < keyLength; i++ {
		p = b
		n = this.base[p]

		if b == this.check[p] && n < 0 {
			result = append(result, -n-1)
		}

		p = b + int(keyChars[i]) + 1
		if b == this.check[p] {
			b = this.base[p]
		} else {
			return result
		}
	}

	p = b
	n = this.base[p]

	if b == this.check[p] && n < 0 {
		result = append(result, -n-1)
	}

	return result
}

func (this *DoubleArrayTrie) resize(newSize int) int {
	base2 := make([]int, newSize)
	check2 := make([]int, newSize)
	used2 := make([]bool, newSize)

	if this.allocSize > 0 {
		copy(base2, this.base)
		copy(check2, this.check)
		copy(used2, this.used)
	}

	this.base = base2
	this.check = check2
	this.used = used2

	this.allocSize = newSize
	return this.allocSize
}

func (this *DoubleArrayTrie) fetch(parent *Node, siblings *list.List) int {
	if this.errno < 0 {
		return 0
	}

	prev := 0
	for i := parent.left; i < parent.right; i++ {
		keyRune := []rune(this.key[i])
		if len(keyRune) < parent.depth {
			continue
		}

		cur := 0
		if len(keyRune) != parent.depth {
			cur = int(keyRune[parent.depth]) + 1
		}

		if prev > cur {
			this.errno = -3
			return 0
		}

		if cur != prev || siblings.Len() == 0 {
			tmpNode := new(Node)
			tmpNode.depth = parent.depth + 1
			tmpNode.code = cur
			tmpNode.left = i
			if siblings.Len() != 0 {
				last := siblings.Back()
				n := toNode(last)
				if n != nil {
					n.right = i
				}
			}

			siblings.PushBack(tmpNode)
		}

		prev = cur
	}

	if siblings.Len() != 0 {
		last := siblings.Back()
		n := toNode(last)
		if n != nil {
			n.right = parent.right
		}
	}
	return siblings.Len()
}

func (this *DoubleArrayTrie) insert(siblings *list.List) int {
	if this.errno < 0 {
		return 0
	}
	firstNode := toNode(siblings.Front())
	if firstNode == nil {
		return 0
	}
	lastNode := toNode(siblings.Back())
	if lastNode == nil {
		return 0
	}

	begin := 0
	var pos int
	if firstNode.code+1 > this.nextCheckPos {
		pos = firstNode.code
	} else {
		pos = this.nextCheckPos - 1
	}

	nonzeroNum := 0
	first := 0

	if this.allocSize <= pos {
		this.resize(pos + 1)
	}

	for {
		pos++

		if this.allocSize <= pos {
			this.resize(pos + 1)
		}

		if this.check[pos] != 0 {
			nonzeroNum++
			continue
		} else if first == 0 {
			this.nextCheckPos = pos
			first = 1
		}

		begin = pos - firstNode.code

		if this.allocSize <= begin+lastNode.code {
			// progress can be zero
			var tmp float64
			if 1.05 > float64(this.keySize)/float64(this.progress+1) {
				tmp = 1.05
			} else {
				tmp = float64(this.keySize) / float64(this.progress+1)
			}

			this.resize(int(float64(this.allocSize) * tmp))
		}

		if this.used[begin] {
			continue
		}

		outerFlag := false

		for e := siblings.Front(); e != nil; e = e.Next() {
			n := toNode(e.Value)
			if n == nil {
				continue
			}
			if n == firstNode { // start from index 1
				continue
			}
			if this.check[begin+n.code] != 0 {
				outerFlag = true
			}
		}
		if outerFlag {
			continue
		}
		break
	}

	// -- Simple heuristics --
	// if the percentage of non-empty contents in check between the
	// index
	// 'next_check_pos' and 'check' is greater than some constant value
	// (e.g. 0.9),
	// new 'next_check_pos' index is written by 'check'.
	if float64(nonzeroNum)/float64(pos-this.nextCheckPos+1) >= 0.95 {
		this.nextCheckPos = pos
	}

	this.used[begin] = true
	if this.size <= begin+lastNode.code+1 {
		this.size = begin + lastNode.code + 1
	}

	for e := siblings.Front(); e != nil; e = e.Next() {
		n := toNode(e.Value)
		if n == nil {
			continue
		}
		this.check[begin+n.code] = begin
	}

	for e := siblings.Front(); e != nil; e = e.Next() {
		n := toNode(e.Value)
		if n == nil {
			continue
		}
		newSiblings := list.New()

		if this.fetch(n, newSiblings) == 0 {
			this.base[begin+n.code] = -n.left - 1

			this.progress++
			// if (progress_func_) (*progress_func_) (progress,
			// keySize);
		} else {
			h := this.insert(newSiblings)
			this.base[begin+n.code] = h
		}
	}

	return begin
}

func toNode(item interface{}) *Node {
	if item == nil {
		return nil
	}
	t := reflect.TypeOf(item)
	if t.Kind() == reflect.Ptr {
		v, ok := item.(Node)
		if ok {
			return &v
		}
	} else if t.Kind() == reflect.Struct {
		v, ok := item.(Node)
		if ok {
			return &v
		}
	}
	return nil
}
