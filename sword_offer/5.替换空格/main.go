package main

import (
	"fmt"
	"strings"
)

/**
题目: 请实现一个函数，把字符串中的每个空格替换成"%20"。例如，
输大 “We are happy.”，则输出“We%20are2%620happy.”。
*/
func main() {
	str := "We are happy."

	fmt.Println(replaceStr(str))
}

func replaceStr(str string) string {
	sb := strings.Builder{}
	flag := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 0x20 {
			sb.WriteString(str[flag:i])
			sb.WriteString("%20")
			flag = i + 1
		}
	}
	sb.WriteString(str[flag:])

	return sb.String()
}
