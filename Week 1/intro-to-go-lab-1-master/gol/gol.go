package main

 //import "fmt"

func getStateOfNeighbour(column int, row int, world[][] byte, width int, height int)bool{
	if column < 0{
		return getStateOfNeighbour(width -1, row, world, width, height)
	}
	if column >= width{
		return getStateOfNeighbour(0, row, world, width, height)
	}
	if row < 0{
		return getStateOfNeighbour(column, height -1, world, width, height)
	}
	if row >= height{
		return getStateOfNeighbour(column, 0, world, width, height)
	}
	if world[row][column] == 0{
		return false
	}
	// fmt.Println("Found a cell that is alive")
	// fmt.Println(int(world[row][column]))
	return true
}

func getNumOfNeighbours(column int, row int, world[][] byte, p golParams) byte{
	var value byte
	value = 0

	xs := []int{ -1, 0, 1, -1, 1, -1, 0, 1}
	ys := []int{ -1, -1, -1, 0, 0, 1, 1, 1}

	for i:=0; i<8; i++{
		if getStateOfNeighbour(column + xs[i], row + ys[i], world, p.imageWidth, p.imageHeight){
			value += 1
		}
	}

	return value
}

func calculateNextState(p golParams, world [][]byte) [][]byte {

	nextState := [][]byte{}

	for rowNum:= 0; rowNum < p.imageWidth; rowNum++{
		row := []byte{}
		for colNum:=0; colNum < p.imageHeight; colNum++{
			value := getNumOfNeighbours(colNum, rowNum, world, p)
			if world[rowNum][colNum] == 255{
				if value == 2 || value == 3{
					row = append(row, 255)
				}else{
					row = append(row, 0)
				}
			}else{
				if value == 3{
					row = append(row, 255)
				}else{
					row = append(row, 0)
				}
			}
		}
		nextState = append(nextState, row)
	}
	
	//VisualiseMatrix(nextState, p.imageWidth, p.imageHeight)
	//fmt.Println(" ")

	return nextState
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	toReturn := []cell{}

	for rowNum, row := range world{
		for colNum, column := range row{
			if column == 255{
				newCell := cell{x: colNum, y: rowNum}
				toReturn = append(toReturn, newCell)
			}
		}
	}

	return toReturn
}
