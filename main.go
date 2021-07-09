package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var visited [10]int

var cost int

func main() {

	//Loop

	//1. Capture number of cases
	//2. Capture the size of the matrix
	//3. take matrix input line by line & create the matrix
	//4. Take the input(source and dest)
	//5. Execute algorithm and create result slice
	//6. Print result slice

	sc := bufio.NewScanner(os.Stdin)

	//sc.Scan()

	//log.Println("Enter number of cases")
	sc.Scan()
	i, _ := strconv.ParseInt(strings.Trim(sc.Text(), " "), 10, 64)
	log.Println("Number  of cases : ", i)

	var N int64

	results := []string{}
	for k := 0; k < int(i); k++ {

		//log.Println("Enter the size of the matrix")
		sc.Scan()
		N, _ = strconv.ParseInt(strings.Trim(sc.Text(), " "), 10, 64)
		log.Println("Matrix size : ", N)

		mat := readMatrix(int(N))

		//    log.Println("Enter the CityA ")
		sc.Scan()
		cities := strings.Split(strings.Trim(sc.Text(), " "), " ")

		log.Println("Cities :  ", cities)
		cityA, _ := strconv.Atoi(cities[0])

		//    log.Println("Enter the CityB ")

		cityB, _ := strconv.Atoi(cities[1])

		//    log.Println(" Matrix : ", mat)

		sp := ShortestPath(mat, cityA, cityB, int(N))

		if sp == 10000 {
			results = append(results, fmt.Sprint("no"))
		} else {
			results = append(results, fmt.Sprintf("yes %d", sp))
		}

	}

	for i, val := range results {
		fmt.Printf("\nCase #%d: %s", i+1, val)
	}

}

func findShortestPath(mat [][]int, N, cityA, cityB int) string {

	for i := 0; i < N; i++ {
		//fmt.Print("\n[")
		for j := 0; j < N; j++ {

			//fmt.Printf(" %d ", mat[i][j])
			if mat[i][j] == 1 {
				log.Printf("City : %d and City %d are directly connected ", i, j)
				continue
			}

		}
		//fmt.Print("]")
	}

	return ""
}

func DFS(mat [][]int, start int, target int, n int) {

	if visited[start] == 0 {
		log.Printf("Visiting %d ", start)

		visited[start] = 1
		for i := 1; i < n; i++ {
			if mat[start][i] == 1 && visited[i] == 0 {
				//visited[i] = 1
				//log.Printf("Start %d, Target : %d", start, target)
				if mat[start][target] == 1 {
					log.Println("Breaking the loop")
					break
				}
				//cost++
				DFS(mat, i, target, n)

			}

		}
	}
	//return cost

}

//func ShortestPath(mat [][]int, start int, target int, n int) int {
func ShortestPath(mat [10][10]int, start int, target int, n int) int {

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				mat[i][j] = int(math.Min(float64(mat[i][j]), float64(mat[i][k]+mat[k][j])))
			}
		}
	}

	return mat[start][target]

}

func readMatrix(N int) [10][10]int {
	var mat [10][10]int
	// sc := bufio.NewScanner(os.Stdin)
	sc := bufio.NewReader(os.Stdin)
	//    log.Println("Enter the matrix .....")
	for i := 0; i < N; i++ {

		// line := strings.TrimRight(sc.Text(), " ")
		line, _, _ := sc.ReadLine()
		log.Println("Line  :", line)

		values := strings.Split(string(line), " ")

		log.Println("Values : ", values)

		for j := 0; j < N; j++ {
			mat[i][j], _ = strconv.Atoi(values[j])

			if mat[i][j] == 0 {
				mat[i][j] = 10000
			}
		}

	}

	return mat

}
