#include<iostream>
using namespace std;
#include "stack.h"
#include "bitree.h"
/*template <typename T>
stack<T>::stack()
{
    top = 0;
    size = 0;
}
template <typename T>
bool stack<T>::InitStack()
{
    base = new stackNode[MAX_SIZE];
    if(!base)
    {
        cout<<"申请内存失败"<<endl;
        return false;
    }
    top = base;
    size = MAX_SIZE;
    return true;
}*/
template <typename T>
bool stack<T>::EmptyStack()
{
    if(top == NULL)
    {
        return true;
    }
    else
    {
        return false;
    }
}
/*template <typename T>
int stack<T>::StackSize()
{
    return 0;
}*/
template <typename T>
bool stack<T>::push(T *node)
{
    stackNode<T> *s = new stackNode<T>;
    if(!s)
    {
        return false;
    }
    s->node = node;
    s->next = top;
    top = s;
    return true;
}
template <typename T>
bool stack<T>::pop(T *node)
{
    if(!top)
    {
        cout<<"栈为空!"<<endl;
        return false;
    }
    stackNode<T> *temp= top;
    node = top->node;
    top = top->next;
    delete temp;
    temp = NULL;
    return true;
}
