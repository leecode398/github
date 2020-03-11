package main
import(
    "fmt"
)

func CenterSpread(s string, center int) int{
    l := len(s)
    i := center - 1
    j := center +1
    step := 0
    for i >= 0 && j < l && s[i] == s[j] {
        i--
        j++
        step++
    }
    return step
}

// func longestPalindrome(s string) string{
//     l := len(s)
//     if l < 2 {
//         return s
//     }

//     str := "$"
//     for i := 0; i < l; i++ {
//         str += string(s[i])
//         str += "#"
//     }
//     size := 2*l+1
//     max := 1

//     start := 0
//     for i := 0; i < size; i++ {
//         curLen := CenterSpread(str, i)
//         if curLen > max {
//             max = curLen
//             start = (i - max) / 2
//         }
//     }
//     return s[start:start+max]
// }

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func maxnum(a int, b int) int {
    if a < b {
        return b
    }
    return a
}

// func manacher(s string) int {
//     l := len(s)
//     max := 0

//     id := 0
//     mx := 0

//     p := make([]int, l)

//     for i := 1; i < l; i++ {
//         if i < mx {
//             p[i] = min(p[2*id-i], mx-i)
//         } else {
//             p[i] = 1
//         }
//         if s[i-p[i]] == s[i+p[i]] {
//             p[i]++
//         }

//         if mx < i + p[i] {
//             id = i
//             mx = i + p[i]
//         }
//         max = maxnum(max, p[i]-1)
//     }
//     return max
// }

func longestPalindrome(s string) string {
    if s == "" {
        return s
    }
    sByte := []byte(s)
    sNew := []byte{}
    sNew = append(sNew,'$')
    l := len(s)
    for i := 0; i < l; i++ {
        sNew = append(sNew, '#')
        sNew = append(sNew, sByte[i])
    }
    sNew = append(sNew, '#')
    sNew = append(sNew, '\\')



    s1 := string(manacher(sNew))
    return s1
}

func manacher(sNew []byte) []byte {
    p := make(map[int]int)
    maxMid := 0
    p[maxMid] = 1
    maxRight := maxMid + p[maxMid] - 1

    result := 0


    length := len(sNew) - 1

    for i := 1; i < length; i++ {

        if i <= maxRight {
            if p[2*maxMid - i] <= maxRight - i + 1 {
                p[i] = p[2*maxMid-i]
            } else {
                p[i] = maxRight - i + 1
            }
        } else {
            p[i] = 1
        }
        for sNew[i-p[i]] == sNew[i+p[i]] {
            p[i]++
        }
        if i+p[i]-1 > maxRight {
            if p[i] > p[result] {
                result = i
            }
            maxMid = i
            maxRight = i + p[i] -1
        }
    }
    sFinal := []byte{}
    for i := result-p[result]+2; i <= result+p[result]-2; i += 2{
        sFinal = append(sFinal, sNew[i])
    }
    return sFinal
}


func main () {
    s := "cbbd"
    fmt.Println(longestPalindrome(s))
}