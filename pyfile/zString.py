
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        l = len(s)
        rs = list()
        j = 0
        for i in range(numRows):
            while i < l:
                print(s[i])
                rs[j] = s[i]
                j += 1
                i += (num-1)

def main():
    Solution1 = Solution()
    s = "PAYPALISHIRING"
    numRows = 3
    print(Solution1.convert(s, numRows))

if __name__ == "__main__":
    main()
