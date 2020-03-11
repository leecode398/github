/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。*/

package main

import (
    "fmt"
)

func isMatch(s string, p string) bool {
    ls := len(s)
    lp := len(p)
    dp := make([][]bool, ls+1)
    flag := true

    for i := range dp {
        dp[i] = make([]bool, lp+1)
    }
    dp[0][0] = true
    for i := 1; i <= lp; i++ {
        if p[i-1] != '*' {
            flag = false
        }
        dp[0][i] = flag
    }

    for i, a := range s {
        for j, b := range p {
            if a == b || b == '?' {
                dp[i+1][j+1] = dp[i][j]
            } else if b == '*' {
                if dp[i][j+1] {
                    dp[i+1][j+1] = true
                } else if dp[i+1][j] {
                    dp[i+1][j+1] = true
                }
            }
        }
    }
    return dp[ls][lp]
}

func main() {
    s := "aa"
    p := "*"

    fmt.Println(isMatch(s, p))
}