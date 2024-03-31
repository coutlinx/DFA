package daf

import (
	"fmt"
	"testing"
	"time"
)

func TestDFA(t *testing.T) {
	words := "习近平你干啥，毛泽东你也干啥,asdasfafuckyouFUCKYOUfasfafhlhasghignbasdghfuaaghrhyjykykjejeejerhreheyweruyqbwpogauygvbbmznvagoqpruinopH，我日你仙人,尼玛，你妈"
	start := time.Now() // 获取开始时间
	if sw := SearchWordsAll(words); len(sw) > 0 {
		fmt.Println("存在敏感词", sw)
	} else {
		fmt.Println("不存在敏感词")
	}
	elapsed := time.Since(start) // 计算耗费时间
	fmt.Printf("SearchWordsAll cost: %d ms\n", elapsed.Milliseconds())

}
