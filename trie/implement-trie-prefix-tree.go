package main

type TrieNode struct {
	IsEnd    bool
	Children [26]*TrieNode
	Word     string
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		b := int(word[i] - 'a')
		if node.Children[b] == nil {
			node.Children[b] = &TrieNode{}
		}
		node = node.Children[b]
	}
	node.IsEnd = true
	node.Word = word
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		b := int(word[i] - 'a')
		if node.Children[b] == nil {
			return false
		}
		node = node.Children[b]
	}
	return node.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		b := int(prefix[i] - 'a')
		if node.Children[b] == nil {
			return false
		}
		node = node.Children[b]
	}
	return true
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
