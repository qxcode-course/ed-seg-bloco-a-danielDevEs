package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func trigo(pen *Pen, comprimento float64, x, y float64){
	if comprimento < 10{
		return
	}
	pen.Walk(comprimento)
				
	pen.Right(40)
	trigo(pen, comprimento*0.5, x, y)
		
	pen.Left(80)
	trigo(pen, comprimento*0.5, x, y)
		
	pen.Right(40)
	pen.Walk(-comprimento)
}

func main() {
	pen := NewPen(1500, 1500)
	pen.SetPosition(750, 1500)
	pen.SetHeading(90)
	
	trigo(pen, 400, 550, 550)
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}