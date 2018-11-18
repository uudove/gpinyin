package pconv

import (
	"container/list"
	"reflect"
)

type node struct {
	code  int
	depth int
	left  int
	right int
}

type doubleArrayTrie struct {
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

func (t *doubleArrayTrie) build(key []string) int {
	if key == nil {
		return 0
	}
	t.key = key
	t.keySize = len(key)
	t.progress = 0

	t.base[0] = 1
	t.nextCheckPos = 0

	rootNode := new(node) // pointer
	rootNode.left = 0
	rootNode.right = t.keySize
	rootNode.depth = 0

	siblings := list.New()

	t.fetch(rootNode, siblings)
	t.insert(siblings)

	t.used = nil
	t.key = nil

	return t.errno
}

func (t *doubleArrayTrie) commonPrefixSearch(key string) []int {
	keyChars := []rune(key)
	keyLength := len(keyChars)
	result := make([]int, keyLength)

	b := t.base[0]
	n := 0
	p := 0

	for i := 0; i < keyLength; i++ {
		p = b
		n = t.base[p]

		if b == t.check[p] && n < 0 {
			result = append(result, -n-1)
		}

		p = b + int(keyChars[i]) + 1
		if b == t.check[p] {
			b = t.base[p]
		} else {
			return result
		}
	}

	p = b
	n = t.base[p]

	if b == t.check[p] && n < 0 {
		result = append(result, -n-1)
	}

	return result
}

func (t *doubleArrayTrie) resize(newSize int) int {
	base2 := make([]int, newSize)
	check2 := make([]int, newSize)
	used2 := make([]bool, newSize)

	if t.allocSize > 0 {
		copy(base2, t.base)
		copy(check2, t.check)
		copy(used2, t.used)
	}

	t.base = base2
	t.check = check2
	t.used = used2

	t.allocSize = newSize
	return t.allocSize
}

func (t *doubleArrayTrie) fetch(parent *node, siblings *list.List) int {
	if t.errno < 0 {
		return 0
	}

	prev := 0
	for i := parent.left; i < parent.right; i++ {
		k := []rune(t.key[i])
		if len(k) < parent.depth {
			continue
		}

		cur := 0
		if len(k) != parent.depth {
			cur = int(k[parent.depth]) + 1
		}

		if prev > cur {
			t.errno = -3
			return 0
		}

		if cur != prev || siblings.Len() == 0 {
			tmpNode := new(node)
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

func (t *doubleArrayTrie) insert(siblings *list.List) int {
	if t.errno < 0 {
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
	if firstNode.code+1 > t.nextCheckPos {
		pos = firstNode.code
	} else {
		pos = t.nextCheckPos - 1
	}

	nonzeroNum := 0
	first := 0

	if t.allocSize <= pos {
		t.resize(pos + 1)
	}

	for {
		pos++

		if t.allocSize <= pos {
			t.resize(pos + 1)
		}

		if t.check[pos] != 0 {
			nonzeroNum++
			continue
		} else if first == 0 {
			t.nextCheckPos = pos
			first = 1
		}

		begin = pos - firstNode.code

		if t.allocSize <= begin+lastNode.code {
			// progress can be zero
			var tmp float64
			if 1.05 > float64(1.0)*float64(t.keySize)/float64(t.progress+1) {
				tmp = 1.05
			} else {
				tmp = float64(1.0) * float64(t.keySize) / float64(t.progress+1)
			}

			t.resize(int(float64(t.allocSize) * tmp))
		}

		if t.used[begin] {
			continue
		}

		outerFlag := false

		for e := siblings.Front(); e != nil; e = e.Next() {
			n := toNode(e.Value)
			if n == nil {
				continue
			}
			if t.check[begin+n.code] != 0 {
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
	if float64(1.0)*float64(nonzeroNum)/float64(pos-t.nextCheckPos+1) >= 0.95 {
		t.nextCheckPos = pos
	}

	t.used[begin] = true
	if t.size <= begin+lastNode.code+1 {
		t.size = begin + lastNode.code + 1
	}

	for e := siblings.Front(); e != nil; e = e.Next() {
		n := toNode(e.Value)
		if n == nil {
			continue
		}
		t.check[begin+n.code] = begin
	}

	for e := siblings.Front(); e != nil; e = e.Next() {
		n := toNode(e.Value)
		if n == nil {
			continue
		}
		newSiblings := list.New()

		if t.fetch(n, newSiblings) == 0 {
			t.base[begin+n.code] = -n.left - 1

			t.progress++
			// if (progress_func_) (*progress_func_) (progress,
			// keySize);
		} else {
			h := t.insert(newSiblings)
			t.base[begin+n.code] = h
		}
	}

	return begin
}

func getListValue(li *list.List, index int) interface{} {
	if index < 0 || index >= li.Len() {
		return nil
	}

	i := 0
	for e := li.Front(); e != nil; e = e.Next() {
		if i == index {
			return e.Value
		}
	}
	return nil
}

func toNode(item interface{}) *node {
	t := reflect.TypeOf(item)
	if t.Kind() == reflect.Ptr {
		v, ok := item.(node)
		if ok {
			return &v
		}
	} else if t.Kind() == reflect.Struct {
		v, ok := item.(node)
		if ok {
			return &v
		}
	}
	return nil
}
