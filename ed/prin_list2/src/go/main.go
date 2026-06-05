package main

import (
	"container/list"
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	str := ""
	str += "[ "
	for node := l.Front(); node != nil; node = node.Next(){
		if node.Value == 0 {
			continue
		}
		if sword.Value.(int) > 0 && sword == node {
			str += fmt.Sprintf("%v> ", node.Value)
			continue
		}
		if sword.Value.(int) < 0 && sword == node {
			str += fmt.Sprintf("<%v ", node.Value)
			continue
		}
		str += fmt.Sprintf("%v ", node.Value)
	}
	str += "]"
	return str
}

// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
	if it.Next() == nil {
		return l.Front()
	}
	return it.Next()
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
	if it.Prev() == nil {
		return l.Back()
	}
	return it.Prev()
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}
