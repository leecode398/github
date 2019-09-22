#ifndef Stack_h
#define Stack_h
 
template <class T>
class Stack {
public:
    Stack(){}; // 构造函数
    void Push(const T x); // 新元素进栈
    bool Pop(); // 栈顶元素出栈
    virtual bool getTop(T &x) const = 0; // 读取栈顶元素，由x返回
    virtual bool isEmpty() const = 0; // 判断栈空否
    // virtual bool isFull() const = 0; // 判断栈满否,因为链式栈不存在不满的情况
    virtual int getSize() const = 0; // 计算栈中元素个数
};
 
 
#endif /* Stack_h */
