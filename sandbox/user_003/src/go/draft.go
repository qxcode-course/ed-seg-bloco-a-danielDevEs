package main
import "fmt"


type Node struct {
    data int
    next *Node
}

type Lista struct {
    head *Node
}


func __push_back (node *Node, value int) *Node {
    if node == nil {
        return &Node{data: value}
    }

    node.next = __push_back(node.next, value)
    return node
}

func (l *Lista) PushBack(value int) {
    l.head = __push_back(l.head, value)
}

func __pop_back (node *Node, value int) *Node {
    if node.next == nil {
        return nil
    }

    node.next = __pop_back(node.next)
    return node
}

func (l *Lista) PopBack(value int) {
    l.head = __push_back(l.head, value)
}

func main() {
    fmt.Println()
}