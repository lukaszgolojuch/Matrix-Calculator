/*
Program name: Matrix Calculator
Programming language: Go
Programming environment: Visual Studio Code
Author: Łukasz Gołojuch
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

//Global variables declaration
var matrix_1 [100][100]float64 //First Matrix
var matrix_2 [100][100]float64 //Second Matrix
var matrix [100][100]float64   //Third Matrix after counting

func adding(x int, y int) {
	// function that add one matrix to other

	var line_str string
	var letter string

	for i := 0; i < x; i++ {
		line_str = ""
		for j := 0; j < y; j++ {
			matrix[i][j] = matrix_1[i][j] + matrix_2[i][j]
			letter = strconv.FormatFloat(matrix[i][j], 'f', 6, 64)
			line_str = line_str + letter + " "
		}
		fmt.Println(line_str)
	}

}

func multiplication(x1 int, y1 int, x2 int, y2 int) {
	// function that multiply one matrix with other

	var line_str string
	var letter string

	for i := 0; i < x1; i++ {
		line_str = ""
		for j := 0; j < y2; j++ {
			for k := 0; k < y1; k++ {
				matrix[i][j] += matrix_1[i][k] * matrix_2[k][j]
			}
			letter = strconv.FormatFloat(matrix[i][j], 'f', 6, 64)
			line_str = line_str + letter + " "
		}
		fmt.Println(line_str)
	}
}

func main() {
	//Local variable declaration
	var x1, x2, y1, y2 int //Size of matrices
	var answ int           //Answer for menu

	//Start of actual program
	fmt.Println("======================================")
	fmt.Println("Welcome in matrix calculator")
	fmt.Println("What is size of first matrix? (max 100x100)")

	// Input size of two matrix
	fmt.Println("X1: ") //First matrix x
	fmt.Scanln(&x1)
	fmt.Println("Y1: ") //First matrix y
	fmt.Scanln(&y1)
	fmt.Println("X2: ") //Second matrix x
	fmt.Scanln(&x2)
	fmt.Println("Y2: ") //Second matrix y
	fmt.Scanln(&y2)

	//Check input max matrix 100x100
	if x1 > 100 || y1 > 100 || x2 > 100 || y2 > 100 {
		//If wrong input try again
		fmt.Println("Wrong input...")
		fmt.Println("Please try again...")
		main()
	}

	fmt.Println("--------------------------------------")
	fmt.Println("First matrix: ")
	//Input for first matrix
	for i := 0; i < x1; i++ {
		for j := 0; j < y1; j++ {
			fmt.Println("Podaj pozycje x: ", i, " y: ", j)
			fmt.Scanln(&matrix_1[i][j])
		}
	}
	fmt.Println("--------------------------------------")
	fmt.Println("Second matrix: ")
	//Input for second matrix
	for i := 0; i < x2; i++ {
		for j := 0; j < y2; j++ {
			fmt.Println("Podaj pozycje x: ", i, " y: ", j)
			fmt.Scanln(&matrix_2[i][j])
		}
	}
	fmt.Println("--------------------------------------")

	//MENU
	fmt.Println("Menu:")
	fmt.Println("1. Adding a matrix")
	fmt.Println("2. Matrix multiplication")
	fmt.Println("3. Exit")
	fmt.Println("Your choice:")
	//End of menu and get answer
	fmt.Scanln(&answ)

	switch answ {
	case 1:
		//Add Matrix
		if x1 == x2 && y1 == y2 {
			// if matrices are same size we can add one to other
			adding(x1, y1)
		} else {
			//Wrong choice
			fmt.Println("You cannot add this matrices")
			main()
		}
	case 2:
		//Multiply matrix
		if y1 == x2 {
			// if y1 == x2 we can multiply both matricies
			multiplication(x1, y1, x2, y2)
		} else {
			//Wrong choice
			fmt.Println("Such matrices cannot be multiplied")
			main()
		}
	case 3:
		//Exit program
		fmt.Println("Thank you for using my program!")
		os.Exit(1)
	}
}
