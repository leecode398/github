//
//  main.c
//  Thread
//
//  Created by lx on 2019/5/12.
//  Copyright © 2019 lx. All rights reserved.
//

#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
void print_a(void);
void print_b(void);

int main(int argc, const char * argv[]) {
    
    pthread_t p1;
    pthread_t p2;
    if(pthread_create(&p1,NULL,print_a,NULL)==-1)
    {
        puts("fail to create pthread t0");
        exit(1);
    }
    if(pthread_create(&p2,NULL,print_b,NULL)==-1)
    {
        puts("fail to create pthread t0");
        exit(1);
    }
    pthread_join(p1, NULL);
    pthread_join(p2, NULL);
    return 0;
}

void print_a()
{
    for(int i=0;i<100;i++)
    {
        printf("aaaa%d\n",i);
//        sleep(1);
    }
}
void print_b()
{
    for(int i=0;i<100;i++)
    {
        printf("bbbb%d\n",i);
//        sleep(1);
    }
}
