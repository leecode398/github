#include <stdio.h>
#include <string.h>
int main()
{
	int a[5] = {1,2,3,4,5};
	for(int i : a)
	{
		printf("%d\t",i);
	}
	return 0;
}
