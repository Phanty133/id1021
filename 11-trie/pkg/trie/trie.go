package trie

type Trie struct {
	root *Node
}

type Node struct {
	next  []*Node
	valid bool
}

func NewNode() *Node {
	return &Node{
		next:  make([]*Node, 27),
		valid: false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func GetCharCode(char rune) int {
	// Skips Q and W

	switch char {
	case 'å':
		return 24
	case 'ä':
		return 25
	case 'ö':
		return 26
	default:
		if char >= 'q' {
			char -= 1
		}

		if char >= 'w' {
			char -= 1
		}

		return int(char) - 'a'
	}
}

func GetCharFromCode(code int) rune {
	switch code {
	case 24:
		return 'å'
	case 25:
		return 'ä'
	case 26:
		return 'ö'
	default:
		if code >= 16 {
			code += 1
		}

		if code >= 22 {
			code += 1
		}

		return rune(code + 'a')
	}
}

func GetPossibleKeys(key rune) []rune {
	idx := getIdx(key)
	possibleKeys := make([]rune, 3)

	for i := 0; i < 3; i++ {
		possibleKeys[i] = GetCharFromCode(idx*3 + i)
	}

	return possibleKeys
}

func getBranchIdxesFromKey(key rune) []int {
	keys := GetPossibleKeys(key)
	branchIdxes := make([]int, 3)

	for i, key := range keys {
		branchIdxes[i] = GetCharCode(key)
	}

	return branchIdxes
}

func getIdx(key rune) int {
	return int(key-'0') - 1
}

func getKey(idx int) rune {
	return rune(idx + 1 + '0')
}

func getKeyFromChar(char rune) rune {
	return getKey(GetCharCode(char) / 3)
}

func WordToSequence(word string) string {
	seq := ""

	for _, char := range word {
		seq += string(getKeyFromChar(char))
	}

	return seq
}

func (t *Trie) AddWord(word string) {
	curr := t.root

	for _, char := range word {
		idx := GetCharCode(char)

		if curr.next[idx] == nil {
			curr.next[idx] = NewNode()
		}

		curr = curr.next[idx]
	}

	curr.valid = true
}

func (t *Trie) Lookup(seq string) []string {
	// Gets key sequence and returns all possible words
	output := make([]string, 0)
	t.root.Lookup(seq, "", &output)

	return output
}

func (t *Node) Lookup(seq string, path string, output *[]string) {
	// Gets key sequence and returns all possible words.

	if len(seq) == 0 {
		if t.valid {
			*output = append(*output, path)
		}

		return
	}

	idx := getBranchIdxesFromKey(rune(seq[0]))

	for _, i := range idx {
		if t.next[i] != nil {
			t.next[i].Lookup(seq[1:], path+string(GetCharFromCode(i)), output)
		}
	}
}
