package main

import (
	"fmt"

	"github.com/thoohv5/algorithm/finishing/tire/trie"
)

/**
概念：
前缀树、也称字典树（Trie），是N叉树的一种特殊形式，前缀树的每一个节点代表一个字符串（前缀）。每一个节点会有多个子节点，通往不同子节点的路径上有着不同的字符。子节点代表的字符串是由节点本身的原始字符串，以及通往该子节点路径上所有的字符组成的。
*/

/**
使用场景：
1、自动补全（单词自动补全）
2、拼写检查（检查单词是否拼写正确）
3、IP路由（最长前缀匹配）
4、九宫格打字预测（根据前缀预测单词）
*/

func main() {
	sensWords := trie.New()
	sensWords.AddWord("傻逼")
	sensWords.AddWord("sb")
	sensWords.AddWord("你懂得")

	fmt.Println([]rune("傻逼"))
	fmt.Println([]rune("sb"))
	fmt.Println([]rune("你懂的"))
	fmt.Println([]rune("你说傻逼，你懂得，哈哈哈，sb,sss"))

	fmt.Println(sensWords.Replace("你说傻逼，你懂得，哈哈哈，sb,sss"))
}
