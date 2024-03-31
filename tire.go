package daf

import "strings"

// DFATrie 树结构体
type DFATrie struct {
	root *DFANode
}

// 创建敏感词树
func newTrie() *DFATrie {
	return &DFATrie{root: newNode()}
}

// 判断是否有敏感词
func (t *DFATrie) search(text string) bool {
	currentNode := t.root
	// 判断子节点存不存在相同字，若到底都存在则存在敏感词
	for _, ch := range text {
		node, ok := currentNode.Children[ch]
		if !ok {
			break
		}
		if node.isEnd {
			return true
		}
		currentNode = node
	}
	return false
}

// 找到第一个敏感词并且返回
func (t *DFATrie) searchOne(text string) (string, bool) {
	currentNode := t.root
	word := ""
	for _, ch := range text {
		node, ok := currentNode.Children[ch]
		if !ok {
			break
		}
		word += string(ch)
		if node.isEnd {
			// 当找到一个标记为结尾的节点时，返回true和当前累积的敏感词
			return word, true
		}
		currentNode = node
	}
	// 如果遍历完给定的文本还没有找到敏感词，返回false和空字符串
	return "", false
}

// 找到所有的敏感词
func (t *DFATrie) searchAll(text string) []string {
	var foundWords []string
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		currentNode := t.root
		word := ""
		for j := i; j < len(runes); j++ {
			ch := runes[j]
			node, ok := currentNode.Children[ch]
			if !ok {
				break
			}
			word += string(ch)
			if node.isEnd {
				// 发现一个敏感词，将其添加到结果列表中
				foundWords = append(foundWords, word)
				// 这里不需要立即break，因为可能存在更长的敏感词
			}
			currentNode = node
		}
	}

	// 去除重复的敏感词
	uniqueWords := removeDuplicates(foundWords)

	return uniqueWords
}

// 函数去除字符串切片中的重复项
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	var result []string

	for v := range elements {
		if encountered[elements[v]] == true {
			// 不添加重复元素
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// Insert 插入敏感词
func (t *DFATrie) insert(words string) {
	words = strings.ToLower(words)
	currentNode := t.root
	runes := []rune(words)
	for _, char := range runes {
		if _, ok := currentNode.Children[char]; !ok {
			currentNode.Children[char] = &DFANode{Children: make(map[rune]*DFANode)}
		}
		currentNode = currentNode.Children[char]
	}
	currentNode.isEnd = true
}

// 修改敏感词
func (t *DFATrie) update(oldWords, newWords string) {
	RemoveWord(oldWords)
	t.insert(newWords)
}
