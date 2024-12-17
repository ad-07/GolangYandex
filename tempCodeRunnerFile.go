package main

import "fmt"

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, fmt.Errorf("negative number/s")
	}
	cells := make([][]bool, height)
    for i := range cells {
        cells[i] = make([]bool, width)
    }
	return &World{Height: height, Width: width, Cells: cells}, nil
}

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func main(){
	print(NewWorld(0, 50))
}