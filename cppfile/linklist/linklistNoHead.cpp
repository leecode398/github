#include<iostream>
using namespace std;

class node
{
public:
    int data;
    node* next;
};
node* linklistinit(node *root)
{
    cout<<"输入第一个节点data"<<endl;
    int x;
    cin>>x;
    root->next = NULL;
    root->data = x;
    return root;
}
node* preinsert(node **root, node *tmpNode)
{
    tmpNode->next = *root;
    *root = tmpNode;
    return *root;
}
node* ininsert(node *root, node *tmpNode)
{
    node* pNode = root;
    while(pNode->next)
    {
        pNode = pNode->next;
    }
    tmpNode->next = NULL;
    pNode->next = tmpNode;
    return root;
}
void print(node *root)
{
    while(root)
    {
        cout<<root->data<<" ";
        root = root->next;
    }
}
node* inverse(node **root)
{
    node* pNode = (*root)->next;
    node* qNode = pNode->next;
    (*root)->next = NULL;
    pNode->next = *root;
    while(qNode->next)
    {
        node* mNode = qNode->next;
        qNode->next = pNode;
        pNode = qNode;
        qNode = mNode;
    }
    qNode->next = pNode;
    *root = qNode;
    return *root;
}
int main()
{
    node* root = new node;
    root->data = 0;
    linklistinit(root);
    for(int i = 0; i < 5; i++)
    {
        int x;
        cout<<"输入第"<<i+1<<"个数:"<<endl;
        cin>>x;
        node* tmpNode = new node;
        tmpNode->data = x;
        tmpNode->next = NULL;
        preinsert(&root, tmpNode);
    }
    print(root);
    cout<<endl;
    for(int i = 0; i < 5; i++)
    {
        int x;
        cout<<"输入第"<<i+1<<"个数:"<<endl;
        cin>>x;
        node* tmpNode = new node;
        tmpNode->data = x;
        tmpNode->next = NULL;
        ininsert(root, tmpNode);
    }
    print(root);
    cout<<endl;
    cout<<"反转后"<<endl;
    inverse(&root);
    print(root);
    cout<<endl;
    return 0;
}
