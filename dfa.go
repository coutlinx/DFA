package daf

import "log"

var SensitiveTire *DFATrie

func init() {
	SensitiveTire = newTrie()
	dir := "C:\\Users\\Administrator\\Desktop\\DAF\\excel"
	// 读取Excel文件
	if err := InitTrie(dir); err != nil {
		log.Fatal(err)
	}
}

// SearchWords DAF算法检查敏感词
func SearchWords(text string) bool {
	return SensitiveTire.search(text)
}

// SearchWordsOne 查询第一个敏感词
func SearchWordsOne(text string) (string, bool) {
	return SensitiveTire.searchOne(text)
}

func SearchWordsAll(text string) []string {
	return SensitiveTire.searchAll(text)
}

// RemoveWord 删除敏感词
func RemoveWord(word string) {
	removeHelper([]rune(word), SensitiveTire.root, 0)
}

// InsertWord 新增敏感词
func InsertWord(word string) {
	SensitiveTire.insert(word)
}

// UpdateWord 修改敏感词
func UpdateWord(oldWord, newWord string) {
	SensitiveTire.update(oldWord, newWord)
}

// 删除的递归方法
func removeHelper(runes []rune, currentNode *DFANode, charIndex int) {
	if charIndex == len(runes) {
		// 递归的基案，当我们到达单词的末尾时
		if currentNode.isEnd {
			currentNode.isEnd = false
		}
		return
	}

	// 当前字符
	char := runes[charIndex]
	// 递归到下一个字符
	if _, ok := currentNode.Children[char]; ok {
		removeHelper(runes, currentNode.Children[char], charIndex+1)
		// 如果当前节点是叶节点，则删除它
		if len(currentNode.Children[char].Children) == 0 && !currentNode.Children[char].isEnd {
			delete(currentNode.Children, char)
		}
	} else {
		// 如果我们找不到字符，则该词不在trie中
		return
	}

	// 清除没有子节点且不是单词结束的节点
	if len(currentNode.Children) == 0 && !currentNode.isEnd {
		delete(currentNode.Children, char)
	}

	return
}
