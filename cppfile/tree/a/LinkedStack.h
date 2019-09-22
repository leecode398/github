#ifndef LinkedStack_h
#define LinkedStack_h
#include <iostream>
#include "Stack.h"
using namespace std;
 
template <class T>
struct LinkNode {
    T data;
    LinkNode<T> *link;
};
 
//类的前置声明
template <class T>
class LinkedStack;
 
//友元函数的声明
template <class T>
ostream& operator<<(ostream& out, LinkedStack<T>& s);
 
 
template <class T>
class LinkedStack: public Stack<T> {
public:
    LinkedStack(); // 构造函数
    ~LinkedStack();// 析构函数
    void Push(const T x); // 进栈
    bool Pop(); // 出栈
    bool getTop(T &x) const; // 读取栈顶元素
    bool isEmpty()const; // 判断栈是否为空
    int getSize()const; // 求栈的元素个数
    void makeEmpty(); // 清空栈的内容
    friend ostream& operator << <T>(ostream& out, LinkedStack<T>& s); // 重载输出函数
private:
    LinkNode<T> *top; // 栈顶指针，即链头指针
};
 
template <class T>
LinkedStack<T>::LinkedStack() {
    // 构造函数，置空栈
    top = new LinkNode<T>(); // 引入头指针：不存放数据
    top->link = NULL;
}
 
template <class T>
LinkedStack<T>::~LinkedStack() {
    // 析构函数,释放内存空间
    makeEmpty();
}
 
template <class T>
void LinkedStack<T>::Push(const T x) {
    // 进栈：将元素值x插入到链式栈的栈顶，即链头
    LinkNode<T> *newNode = new LinkNode<T>(); // 创建包含x的新结点
    if(newNode == NULL) {
        cerr << "内存空间分配失败！" << endl;
        exit(1);
    }
    newNode->data = x;
    newNode->link = top->link; // 指向头指针的下一个结点：即栈中第一个存放有效数据的结点
    top->link = newNode; // 头指针往前移
}
 
template <class T>
bool LinkedStack<T>::Pop()  {
    // 出栈：删除栈顶结点
    if(isEmpty())
        return false; // 栈空，不出栈
    LinkNode<T> *p = top->link; // 暂存栈顶元素
    top->link = p->link; // 栈顶指针退到新的栈顶位置
    delete p;
    p = NULL;
    return true;
}
 
template <class T>
bool LinkedStack<T>::getTop(T &x) const {
    // 读取栈顶元素
    if(isEmpty())
        return false;
    x = top->link->data; // 栈不空，返回栈顶元素的值。这里top为头指针，所以栈顶元素为：top->link
    return true;
}
 
template <class T>
bool LinkedStack<T>::isEmpty()const {
    // 判断栈是否为空
    if(top->link == NULL) // 栈为空
        return true;
    return false;
}
 
template <class T>
int LinkedStack<T>::getSize()const {
    // 求栈的元素个数
    int len = 0;
    
    LinkNode<T> *current = top->link;
    while(current != NULL) {
        len++;
        current = current->link;
    }
    return len;
}
 
template <class T>
void LinkedStack<T>::makeEmpty() {
    // 清空栈的内容
    LinkNode<T> *current = top->link;
    while(current != NULL) {
        top->link = current->link; // 保存链式栈准备要删除的结点的下一个结点，防止丢失
        delete current; // 释放
        current = NULL; // 先指向空
        current = top->link; // 再指向剩下链表的首结点
    }
}
 
template <class T>
ostream& operator<<(ostream& out, LinkedStack<T>& s) {
    // 重载输出函数
    LinkNode<T> *current = s.top->link;
    while(current != NULL) {
        out << current->data << " ";
        current = current->link;
    }
    return out;
}
 
#endif /* LinkedStack_h */
