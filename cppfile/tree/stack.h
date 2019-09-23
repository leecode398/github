#ifndef STACK_H
#define STACK_H
#include "bitree.h"
#include<iostream>
using namespace std;
template<typename T>
class stackNode
{
public:
	T *node;
	stackNode<T>* next;
};
template <typename T>
class stack
{
public:
	int size;
	stackNode<T>* top;
	stack();
	/*bool InitStack();*/
	bool push(T* node);
	bool pop(T* node);
	bool EmptyStack();
	/*bool GetTop(T *node);
	int StackSize();*/
};


template <typename T>
stack<T>::stack()
{
	top = NULL;
	size = 0;
}
/*template <typename T>
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
	if (size == 0)
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
bool stack<T>::push(T* node)
{
	if (!node)
	{
		return false;
	}
	stackNode<T>* temp = new stackNode<T>;
	temp->node = node;
	temp->next = top;
	top = temp;
	size++;
	return true;
}
template <typename T>
bool stack<T>::pop(T* node)
{
	if (size == 0)
	{
		cout << "栈为空!" << endl;
		return false;
	}
	stackNode<T>* temp = top;
	memcpy(node, top->node, sizeof(struct biNode));
	top = top->next;
	size--;
	return true;
}
#endif
