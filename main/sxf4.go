package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, error) {
	line = strings.Trim(line, "[] ")
	parts := strings.Split(line, ",")
	row := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		num, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		row = append(row, num)
	}
	return row, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var grid [][]int
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		row, err := parseLine(text)
		if err != nil {
			fmt.Println("Invalid input.")
			return
		}
		grid = append(grid, row)
	}

	if len(grid) == 0 || len(grid[0]) == 0 {
		fmt.Println(0)
		return
	}

	m, n := len(grid), len(grid[0])
	maxSum := make([][2]int, m*n)
	for i := range maxSum {
		maxSum[i][0] = -1
		maxSum[i][1] = -1
	}

	queueX := make([]int, 0, m*n*2)
	queueY := make([]int, 0, m*n*2)
	queueUsed := make([]bool, 0, m*n*2)
	queueCoins := make([]int, 0, m*n*2)

	startCoins := 0
	if grid[0][0] == -1 {
		queueX = append(queueX, 0)
		queueY = append(queueY, 0)
		queueUsed = append(queueUsed, true)
		queueCoins = append(queueCoins, 0)
		maxSum[0][1] = 0
	} else {
		if grid[0][0] > 0 {
			startCoins += grid[0][0]
		}
		queueX = append(queueX, 0)
		queueY = append(queueY, 0)
		queueUsed = append(queueUsed, false)
		queueCoins = append(queueCoins, startCoins)
		maxSum[0][0] = startCoins
	}

	maxCoins := startCoins
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	head := 0
	for head < len(queueX) {
		x, y, used, coins := queueX[head], queueY[head], queueUsed[head], queueCoins[head]
		head++
		for _, d := range directions {
			ni, nj := x+d.dx, y+d.dy
			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}
			val := grid[ni][nj]
			newCoins := coins
			newUsed := used
			if val == -1 {
				if !used {
					newUsed = true
					val = 0
				} else {
					continue
				}
			}
			if val > 0 {
				newCoins += val
			}
			idx := ni*n + nj
			state := 0
			if newUsed {
				state = 1
			}
			if newCoins > maxSum[idx][state] {
				maxSum[idx][state] = newCoins
				if newCoins > maxCoins {
					maxCoins = newCoins
				}
				queueX = append(queueX, ni)
				queueY = append(queueY, nj)
				queueUsed = append(queueUsed, newUsed)
				queueCoins = append(queueCoins, newCoins)
			}
		}
	}

	fmt.Println(maxCoins)
}
