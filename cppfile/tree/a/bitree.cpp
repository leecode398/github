#include<iostream>
using namespace std;
#include "bitree.h"



biNode* preCreat()
{
    biNode* tmpNode;
    char x;
    cin>>x;
    if(x == '#')
    {
        tmpNode = NULL;
    }
    else
    {
        tmpNode = new biNode;
        tmpNode->data = x;
        tmpNode->left = preCreat();
        tmpNode->right = preCreat();
    }
    return tmpNode;
}
void prePrint(biNode* root)
{
    if(root)
    {
        cout<<root->data<<" ";
        prePrint(root->left);
        prePrint(root->right);
    }
}
void inPrint(biNode* root)
{
    if(root)
    {
        inPrint(root->left);
        cout<<root->data<<" ";
        inPrint(root->right);
    }
}
void postPrint(biNode* root)
{
    if(root)
    {
        postPrint(root->left);
        postPrint(root->right);
        cout<<root->data<<" ";
    }
}
//先序遍历非递归
template<typename T>
void preOrder(Stack<T> *sroot, biNode *root)
{
    sroot->Push(*root);
    while(sroot)
    {
    	biNode tempNode;
    	sroot->getTop(tempNode);
        sroot->Pop();
        cout<<root->data<<" ";
        if(root->right)
        {
            sroot->Push(*(root->right));
        }
        if(root->left)
        {
            sroot->Push(*(root->left));
        }
    }

}
int main()
{
    biNode* root = NULL;
    cout<<"输入节点"<<endl;
    root = preCreat();
    cout<<"递归"<<endl;
    prePrint(root);
    cout<<endl;
    inPrint(root);
    cout<<endl;
    postPrint(root);
    cout<<endl;
    cout<<"非递归"<<endl;
    LinkedStack<biNode> *sroot = new LinkedStack<biNode>;
    preOrder(sroot,root);
    return 0;
}
//  abc##de###fg#h##i##

