package main

import "fmt"

type Square struct {
	Size     int
	Data     [][]int
	sum      int
	sumFound bool
}

func (sq Square) String() string {
	return fmt.Sprintf("%v", sq.Data)
}

func (sq *Square) IsMagic() bool {
	return sq.allUnique() && sq.rowsMatch() && sq.colsMatch() && sq.diagMatch()
}

func (sq *Square) allUnique() bool {
	seen := map[int]bool{}
	for _, row := range sq.Data {
		for _, n := range row {
			if seen[n] {
				return false
			}
			seen[n] = true
		}
	}
	return true
}

func (sq *Square) Sum() int {
	if !sq.sumFound {
		for _, x := range sq.Data[0] {
			sq.sum += x
		}
		sq.sumFound = true
	}
	return sq.sum
}

func (sq *Square) rowsMatch() bool {
	for _, row := range sq.Data {
		sum := 0
		for _, n := range row {
			sum += n
		}
		if sum != sq.Sum() {
			return false
		}
	}
	return true
}

func (sq *Square) colsMatch() bool {
	for i := 0; i < sq.Size; i++ {
		sum := 0
		for j := 0; j < sq.Size; j++ {
			sum += sq.Data[j][i]
		}
		if sum != sq.Sum() {
			return false
		}
	}
	return true
}

func (sq *Square) diagMatch() bool {
	sum1 := 0
	sum2 := 0
	for i := 0; i < sq.Size; i++ {
		sum1 += sq.Data[i][i]
		sum2 += sq.Data[i][sq.Size-1-i]
	}
	return sum1 == sq.Sum() && sum2 == sq.Sum()
}

func coord(i int, n int /* modulo */) (x, y int) {
	x = i % n
	y = i / n
	return
}

func (sq *Square) incrementModulo(n int) Square {
	result := Square{
		Size: sq.Size,
		Data: make([][]int, sq.Size),
	}
	for i := 0; i < sq.Size; i++ {
		result.Data[i] = make([]int, sq.Size)
	}

	totalLen := sq.Size * sq.Size
	carry := 1
	for i := 0; i < totalLen; i++ {
		x := i / sq.Size
		y := i % sq.Size
		result.Data[x][y] = sq.Data[x][y]
		result.Data[x][y] += carry
		carry = 0
		if result.Data[x][y] >= n {
			result.Data[x][y] %= n
			carry = 1
		}
	}
	return result
}

func genSquares(c chan *Square, closeChan chan interface{}, startIndex int) {
	n := 10
	currentSquare := &Square{
		Size: 3,
		Data: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, startIndex}},
	}

	for currentSquare.Data[2][2] == startIndex {
		select {
		case _, _ = <-closeChan:
			return
		default:
		}
		c <- currentSquare
		newSquare := currentSquare.incrementModulo(n)

		currentSquare = &newSquare
	}
}

func main() {
	k := make(chan *Square, 100)
	closeChan := make(chan interface{}, 0)
	for i := 0; i < 10; i++ {
		go genSquares(k, closeChan, i)
	}
	for i := 0; i < 10; i++ {
		go checkMagicSquares(i, k, closeChan)
	}
	fmt.Println("Working!")
	_, _ = <-closeChan
}

func checkMagicSquares(workerId int, c chan *Square, closeChan chan interface{}) {
	nMatches := 1
	k := 0
	for s := range c {
		if k%100000 == 0 {
			fmt.Printf("[%d], %d %v\n", workerId, k, s)
		}
		k++
		if s.IsMagic() {
			fmt.Printf("Magic: %v\n", s)
			nMatches--
			if nMatches == 0 {
				close(closeChan)
			}
		}
	}
}
