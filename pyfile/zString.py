import math
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        l = len(s)
        if numRows == 1 or l < numRows:
            return s
        i = 0
        m = 1
        ll = math.ceil(l/(2*numRows-2)*(numRows-1))
        rs = [['0']*ll for i in range(numRows)]
        for k in range(ll):
            if k%(numRows-1) == 0:
                m -= 1
                for j in range (numRows):
                    if i < l and m < numRows:
                        rs[m][k] = s[i]
                        i += 1
                        m += 1
                m -= 1
            else:
                if i < l:
                    rs[m-1][k] = s[i]
                    i += 1
                    m -= 1
        retstr = list()
        for i in range(numRows):
            for j in range(ll):
                if rs[i][j] != '0':
                    retstr.append(rs[i][j])
        return "".join(retstr)

def main():
    Solution1 = Solution()
    s = "ABC"
    numRows = 2
    print(Solution1.convert(s, numRows))

if __name__ == "__main__":
    main()
