//package main 
package linkedlist 

import "fmt"


type Node struct {
  Data int;
  Next *Node;
}

func NewNode(data int) (*Node){

  var node = Node{
    Data: data,
    Next: nil,
  }
  
  return &node;
}

func (n *Node)LinkNode(node *Node){

  var curr *Node = n;

  if curr == nil {
    fmt.Println("Nil Head From Link Node !")
    n.Data = node.Data;
    return;
  }

  for curr.Next != nil {
    curr = curr.Next; 
  }

  curr.Next = node; 
}

func (node *Node)PrintNodes(){

  var curr *Node = node;

  for curr != nil {
    fmt.Printf("[%v]=>", curr.Data);
    curr = curr.Next;
  }
}


func main(){

  var head *Node;

  for i := 0; i < 10; i++ {
    if head == nil {
      fmt.Println("nil head !")
    }
    node := NewNode(i + 1);
    head.LinkNode(node)
  }

  head.PrintNodes()

}
