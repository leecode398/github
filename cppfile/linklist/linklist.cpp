#include<iostream>
using namespace std;

class node
{
public:
    int data;
    node* next;
    node* last;
};
node* linklistinit(node *root)
{
    root->next = NULL;
    root->last = root;
    return root;
}
node* preinsert(node *root, node *tmpNode)
{
    tmpNode->next = root->next;
    root->next = tmpNode;
    return root;
}
node* ininsert(node *root, node *tmpNode)
{
    while(root->next)
    {
        root = root->next;
    }
    root->last = root;
    root->last->next = tmpNode;
    tmpNode->next = NULL;
    root->last = tmpNode;
    return root;
}
void print(node *root)
{
    node* tmpNode = root;
    while(tmpNode->next)
    {
        cout<<tmpNode->next->data<<" ";
        tmpNode = tmpNode->next;
    }
}
node* inverse(node *root)
{
    node* tmpNode = root->next;
    root->next = NULL;
    while(tmpNode)
    {
        node* inNode = tmpNode;
        node* nNode = tmpNode->next;
        preinsert(root, inNode);
        tmpNode = nNode;
    }
    return root;
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
        preinsert(root, tmpNode);
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
    inverse(root);
    cout<<"反转后"<<endl;
    print(root);
    cout<<endl;
    return 0;
}
