package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func floco(pen *Pen, raio float64, x, y float64){
    if raio < 5 {
        return
    }

    ang := 0.0

    for range 6 {
        theta := ang * math.Pi / 180.0

        newX := x + raio * math.Cos(theta)
        newY := y + raio * math.Sin(theta)

        pen.Up()
        pen.SetPosition(newX, newY)
        pen.Down()

        pen.DrawCircle(raio * 0.34)

        floco(pen, raio*0.34, newX, newY)

        ang += 60.0
    }
}
	
func main() {
	pen := NewPen(1500, 1500)
	pen.SetHeading(0)
	pen.SetPosition(750, 750)
	pen.DrawCircle(300)
	floco(pen, 300, 750, 750)
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}