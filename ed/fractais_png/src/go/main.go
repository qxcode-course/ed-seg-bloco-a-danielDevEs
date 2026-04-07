package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func arvore(pen *Pen, dis float64, larguralinha float64){
	ang := 30.0
	if(dis < 10){
		return
	}
	pen.SetLineWidth(larguralinha)
	pen.SetRGB(float64(randInt(100, 200)), float64(randInt(100, 250)), float64(randInt(100, 200)))
	pen.Walk(dis)
	pen.Left(ang)
	arvore(pen, dis - 6, larguralinha - 1)
	pen.Right(2*ang)
	arvore(pen, dis - 6, larguralinha - 1)
	pen.Left(ang)
	pen.Walk(-dis)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetPosition(250, 500)
	pen.SetHeading(90)
	arvore(pen, 60.0, 10.0)
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
