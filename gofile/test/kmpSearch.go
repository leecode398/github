package main
import (
    "fmt"
)
func strStr(haystack string, needle string) int {
    if len(needle) == 0 || len(needle) > len(haystack){
        return 0
    }
    next := NextArray(needle)
    return KmpSearch(haystack, needle, next)
}

func NextArray(needle string) []int {
    l := len(needle)
    next := make([]int, l)
    next[0] = -1
    k := -1
    i := 0
    for i < l-1 {
        if k == -1 || needle[k] == needle[i] {
            i++
            k++
            if needle[i] != needle[k] {
                next[i] = k
            } else {
                next[i] = next[k]
            }
        } else {
            k = next[k]
        }
    }
    return next
}

func KmpSearch(haystack string, needle string, next []int) int {
    l1 := len(haystack)
    l2 := len(needle)
    i, j := 0, 0
    for i <= l1 && j < l2 {
        if j == -1 || haystack[i] == needle[j] {
            i++
            j++
        } else {
            j = next[j]
        }
    }
    if j == l2 {
        return i - l2
    }
    return -1
}

func main() {
    haystack := "bbcabcdababcdabcdabde"
    needle := "abcdabd"
    next := NextArray(needle)
    KmpSearch(haystack, needle, next)
}