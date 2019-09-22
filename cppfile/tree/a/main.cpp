#include "LinkedStack.h"
using namespace std;
 
int main(int argc, const char * argv[]) {
    int n, x, choice, len; // val存储值，choose存储用户的选择
    bool finished = false;
    LinkedStack<int> L; // 对象
    
    while(!finished) {
        cout << "1:建栈：" << endl;
        cout << "2:进栈" << endl;
        cout << "3:出栈：" << endl;
        cout << "4:读取栈顶元素：" << endl;
        cout << "5:栈是否为空：" << endl;
        cout << "6:栈中的元素个数：" << endl;
        cout << "7:清空栈的内容：" << endl;
        cout << "8:输出栈中元素的值：" << endl;
        cout << "9:退出" << endl;
        cout << "请输入你的选择[1-9]：" << endl;
        cin >> choice;
        switch(choice) {
            case 1:
                cout << "请输入要进栈的数的个数:";
                cin >> n;
                cout << "请输入要进栈的数(以空格隔开)：" << endl;
                for(int i=0; i < n; i++) {
                    cin >> x;
                    L.Push(x);
                }
                break;
            case 2:
                cout << "请输入要进栈的数:";
                cin >> x;
                L.Push(x);
                break;
            case 3:
                if(L.Pop())
                    cout << "出栈成功!" << endl;
                else
                    cout << "栈为空!" << endl;
                break;
            case 4:
                if(L.getTop(x))
                    cout << "栈顶元素的值为：" << x << endl;
                else
                    cout << "栈为空!" << endl;
                break;
            case 5:
                if(L.isEmpty())
                    cout << "栈为空！" << endl;
                else
                    cout << "栈不为空!" << endl;
                break;
            case 6:
                len = L.getSize();
                cout << "栈中的元素个数为：" << len << endl;
                break;
            case 7:
                L.makeEmpty(); // 清空栈
                break;
            case 8:
                if(L.isEmpty())
                    cout << "栈为空！" << endl;
                else
                    cout << L << endl;
                break;
            case 9:
                finished = true;
                break;
            default:
                cout << "输入错误，请重新输入!" << endl;
        } // switch
    } // while
    return 0;
}
