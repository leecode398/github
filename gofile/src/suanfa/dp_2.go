/*给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。*/

package main

import (
    "fmt"
)

func isMatch(s string, p string) bool {
    ls := len(s)
    lp := len(p)
    dp := make([][]bool, ls+1)

    for i := range dp {
        dp[i] = make([]bool, lp+1)
    }
    dp[0][0] = true
    for i, c := range p {
        if c == '*' && dp[0][i-1] {  //当s[i] == * 时只需要看dp[0][i-1] (即s[i-2]位置)是否为true, 例如 missyou 与 missa*you 其中a*不用考虑
            dp[0][i+1] = true
        }
    }

    for i, a := range s {
        for j, b := range p {
            if a == b || b == '.' {
                dp[i+1][j+1] = dp[i][j]
            } else if b == '*' {
                if p[j-1] != s[i] && p[j-1] != '.' {
                    dp[i+1][j+1] = dp[i+1][j-1]
                } else {
                    dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]
                }
            }
        }
    }
    return dp[ls][lp]
}

func main() {
    s := "missssaable"
    p := "mis*a*.le"

    fmt.Println(isMatch(s, p))
}