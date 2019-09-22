#ifndef STACK_H
#define STACK_H
#include "bitree.h"
template<typename T>
class stackNode
{
    T node;
    stackNode<T> *next;
};
template <typename T>
class stack
{
private:
    int size;
    stackNode<T> *top;
public:
    /*stack();
    bool InitStack();*/
    bool push(T *node);
    bool pop(T *node);
    bool EmptyStack();
    /*bool GetTop(T *node);
    int StackSize();*/
};
#endif
