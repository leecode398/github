#include<stdio.h>
#include<stdlib.h>
typedef struct bitnode
{
 char data;
    struct bitnode *lchild,*rchild;
}bitnode,*bitree;//二叉树节点类型和节点指针类型

bitree create()//先序创建
{
 bitree root=NULL;

 char c;
 scanf("%c",&c);
 fflush(stdin);
 if(c=='#')return NULL;
 else
 {
  root=(bitnode*)malloc(sizeof(bitnode));
  root->data=c;
  root->lchild=create();
  root->rchild=create();
 }
 return root;
}

void preorder(bitree root)//先根遍历
{
 if(!root)return;
 else 
 {
	putchar(root->data);
    preorder(root->lchild);
    preorder(root->rchild);
 }
}

void inorder(bitree root)//中根遍历
{
 if(!root)return;
 else
 {
  inorder(root->lchild);
  putchar(root->data);
  inorder(root->rchild);
 }
}

void postorder(bitree root)//后根遍历
{
 if(!root)return;
 else
 {
  postorder(root->lchild);
  postorder(root->rchild);
  putchar(root->data);
 }
}

int leafcount(bitree root)//计算叶子节点
{
 if(!root)return 0;
 else
 {
  if(!root->lchild&&!root->rchild)return 1;
     else return leafcount(root->lchild)+leafcount(root->rchild);
 }
}

int depth(bitree root)//树的高度
{
 if(!root)return 0;
 else
 {
  int m=depth(root->lchild);
  int n=depth(root->rchild);
  return (m>n?m:n)+1;
 }
}

void Revolute(bitree root)// 交换左右子树
{
 bitree t;
 t=root->lchild;
 root->lchild=root->rchild;
 root->rchild=t;
 if(root->lchild)Revolute(root->lchild);
 if(root->rchild)Revolute(root->rchild);
}

int main()
{
 bitree root=NULL;
 printf("输入先序序列：\n");
 root=create();
 printf("前序遍历：\n");
 preorder(root);
 printf("\n");
 printf("中序遍历：\n");
 inorder(root);
 printf("\n");
    printf("后序遍历：\n");
 postorder(root);
    printf("\n");
 printf("二叉树的叶子结点个数为:%d\n",leafcount(root));
 printf("二叉树的高度为:%d\n",depth(root));
    printf("交换左右子树后\n");
 Revolute(root);
    
 printf("前序遍历：\n");
 preorder(root);
 printf("\n");
 printf("中序遍历：\n");
 inorder(root);
 printf("\n");
    printf("后序遍历：\n");
 postorder(root);
    printf("\n");
    return 0;
}