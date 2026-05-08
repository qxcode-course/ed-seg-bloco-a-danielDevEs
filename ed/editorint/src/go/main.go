package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	linhaSalva := NewList[rune]()

	if e.cursor != e.line.Value.End(){
		posCursor := e.line.Value.IndexOf(e.cursor)
		posFinal := e.line.Value.IndexOf(e.line.Value.root)
		for i := posCursor; i < posFinal; i++{
			linhaSalva.PushBack(e.cursor.Value)
			e.line.Value.Erase(e.cursor)
			e.cursor = e.cursor.Next()
		}
	}

	novaLinha := NewList[rune]()

	if linhaSalva.Size() > 0 {
		valorSalvo := linhaSalva.Front()
		for range linhaSalva.Size(){
			novaLinha.PushBack(valorSalvo.Value)
			valorSalvo = valorSalvo.Next()
		}
	}
	e.lines.Insert(e.line.Next(), novaLinha) // Insere uma nova linha na proxima posição
	e.line = e.line.Next() // Move para a linha sucessora
	e.cursor = e.line.Value.Front() // Move o cursor para o inicio da linha
}


func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.root{ // Se o cursor não está no final da linha
		e.cursor = e.cursor.Next() // Move o cursor para a direita
		return
	}

	if e.line != e.lines.Back() { // Se não está na primeira linha
		e.line = e.line.Next()        // Move para a linha anterior
		e.cursor = e.line.Value.Front() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyUp() {
	if e.line == e.lines.Front() { // Se não está na primeira linha
		return
	}

	pos := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Prev()
	it := e.line.Value.Front()
	size := e.line.Value.Size()
	
	for i:= 0; i < pos; i++{
		it = it.Next()
	}
	
	e.cursor = it
	if size < e.line.Next().Value.Size(){
		e.cursor = e.line.Value.Back().root
	}
}

func (e *Editor) KeyDown() {
	if e.line == e.lines.Back() { // Se não está na ultima linha
		return
	}

	pos := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Next()
	it := e.line.Value.Front()
	size := e.line.Value.Size()
	
	for i:= 0; i < pos; i++{
		it = it.Next()
	}
	
	e.cursor = it
	if size < e.line.Prev().Value.Size(){
		e.cursor = e.line.Value.Back().root
	}
}

func (e *Editor) KeyBackspace() {

	if e.cursor == e.line.Value.Front() && e.line != e.lines.Front(){
		e.lines.Erase(e.line)
		e.line = e.line.Prev()
		e.cursor = e.line.Value.root
		return
	}
	e.line.Value.Erase(e.cursor.prev)
}

func (e *Editor) KeyDelete() {
	if e.cursor.Next() == e.line.Value.root && e.line != e.lines.Back(){
		pos := e.line.Next().Value.Size()
		ValorInicial := e.line.Next().Value.Front()
		for range pos{
			e.line.Value.PushBack(ValorInicial.Value)
			ValorInicial = ValorInicial.Next()
		}
		e.lines.Erase(e.line.Next())
		return
	}
	e.line.Value.Erase(e.cursor.Next())
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorGhostWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
