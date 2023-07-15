package trie

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	Root *Node
}

func InitTrie() *Trie {
	result := &Trie{Root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		// 'a' is code 97
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		// 'a' is code 97
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}
