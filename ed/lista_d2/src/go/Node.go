package main

type Node struct{
	value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

// class Node {
//     Value int    // Valor é público
//     next *Node   // o próximo nó da lista
//     prev *Node   // o nó anterior

//     root *Node   // aponta para o nó sentinela da lista da qual ele faz parte
//     Next() *Node // retorna o próximo nó ou nulo, se o próximo é o root
//     Prev() *Node // retorna o nó anterior ou nulo, se o anterior é o root
// }