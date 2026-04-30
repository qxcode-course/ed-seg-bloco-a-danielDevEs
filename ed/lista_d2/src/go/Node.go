package main

class Node {
    Value int    // Valor é público
    next *Node   // o próximo nó da lista
    prev *Node   // o nó anterior

    root *Node   // aponta para o nó sentinela da lista da qual ele faz parte
    Next() *Node // retorna o próximo nó ou nulo, se o próximo é o root
    Prev() *Node // retorna o nó anterior ou nulo, se o anterior é o root
}