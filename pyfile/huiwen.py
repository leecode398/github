#!/usr/bin/env python
#coding=utf-8

class Solution:
    def longestPalindrome(self, s: str) -> str:
        #动态规划
        l = len(s)
        maxlen = 0
        dp = [[0]*l for i in range(l)]
        rs = ""

        for i in range(l):
            for j in range(i+1):
                if i-j <= 1:
                    if s[i] == s[j]:
                        dp[j][i] = 1
                        if maxlen < i-j+1:
                            rs = s[j:i+1]
                            maxlen = i-j+1
                else:
                    if s[i] == s[j] and dp[j+1][i-1]:
                        dp[j][i] = 1
                        if maxlen < i-j+1:
                            rs = s[j:i+1]
                            maxlen = i-j+1
        return rs

def main():
    Solution1 = Solution()
    s = "aaabaaaa"
    print(Solution1.longestPalindrome(s))

if __name__ == "__main__":
    main()
