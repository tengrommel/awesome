/*
Finding the length of connected cells of 1s (regions) in a matrix of 0s and 1s: Given a matrix, each
of which may be 1 or 0. The filled cells that are connected form a region. Two cells are said to be connected if
they are adjacent to each other horizontally, vertically or diagonally. There may be several regions in the matrix.
How do you find the largest region (in terms of number of cells) in the matrix?

Sample Input: 	11000						Sample Output: 5
				01100
				00101
				10001
				01011

Solution: The simplest idea is: for each location traverse in all 8 directions and in each of those directions keep
track of maximum region found.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findConnects(matrix [][]int, M, N, r, c int) int {
	answer := 0
	if r < 0 || c < 0 || r >= M || c >= N {
		answer = 0
	} else if matrix[r][c] == 1 {
		matrix[r][c] = 0
		answer = 1 +
			findConnects(matrix, M, N, r-1, c) +
			findConnects(matrix, M, N, r+1, c) +
			findConnects(matrix, M, N, r, c-1) +
			findConnects(matrix, M, N, r, c+1) +
			findConnects(matrix, M, N, r-1, c-1) +
			findConnects(matrix, M, N, r-1, c+1) +
			findConnects(matrix, M, N, r+1, c-1) +
			findConnects(matrix, M, N, r+1, c+1)
	}
	return answer
}

func findMaxConnects(matrix [][]int, M, N int) int {
	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	return maxConnects
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, M)
	for i := 0; i < M; i++ {
		matrix[i] = make([]int, N)
		for j := 0; j < N; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			matrix[i][j] = v
		}
	}
	maxConnects := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			maxConnects = max(maxConnects, findConnects(matrix, M, N, r, c))
		}
	}
	fmt.Println(findMaxConnects(matrix, M, N))
}
