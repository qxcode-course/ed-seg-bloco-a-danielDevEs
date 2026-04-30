package main
import "fmt"
func main() {


//func (n *Node)


func (list *List) String() string {
    fmt.Print("[ ")
    for x := l.head.next; x != l.head; x = x.next{
        fmt.Printf("%s ")
    }
    fmt.Println("]")
}

func pushFront(list *List, value int){
    inserir(list.head.next, value)
}

func pushFront(list *List, value int){
    inserir(list.head.prev, value)
}

func (list *List) find

func main() {
    list := newList()
    inserir(list.head, 7)
    inserir(list.head, 1)
    inserir(list.head, 3)
    inserir(list.head, 5)
}

type Node struct{
    data int
    next *Node
    prev *Node
}

type List struct{
    head *Node
}

func newList()*List{
    lista := List{}

    lista.head = &Node{}
    lista.head.next = lista.head
    lista.head.prev = lista.head

    return &lista
}

func inserir(node *Node, value int){
    novo := &Novo{
        data: value,
        next: node,
        prev: node.prev,
        node.prev.next: novo,
        node.prev: novo,
    }
}
}
