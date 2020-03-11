money=[1,5,10,20,50,100]
n = int(input())
li=[]
for i in range(n+1):
    li.append(0)
li[0]=1
for i in money:
    for j in range(n+1):
        if j>=i:
            li[j]=li[j]+li[j-i]
    print(li)
