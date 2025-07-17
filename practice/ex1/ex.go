package ex1

import (
	"fmt"
	"strings"
)

func Run() {
	//realization()

}

// As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.
//
// Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:
//
// z -= (z*z - x) / (2*z)
// Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.
//
// Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
//
// Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:
//
// z := 1.0
// z := float64(1)
// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the math.Sqrt in the standard library?
//
// (Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method. It works well for many functions but especially well for square root.)
func findSquare(x float64) float64 {
	z := float64(1)
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("Итерация %d: z = %.10f\n", i+1, z)
	}
	return z
}

func realization() {

	for i := 2.0; i <= 5.0; i++ {
		countSqrt := findSquare(i)
		fmt.Printf("Counting sqrt(%v): %v\n", i, countSqrt)
	}
}

// крестики нолики
func makeGame() {
	//Здесь создаётся двумерный срез board — игровое поле размером 3×3.
	//Каждая строка — это отдельный срез из 3 элементов, содержащих строки "_", которые обозначают пустую клетку.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X" //X делает ход в верхний левый угол ([0][0])
	board[2][2] = "O" //O — в нижний правый ([2][2])
	board[1][2] = "X" //X — в центр правой колонки ([1][2])
	board[1][0] = "O" //O — в центр левой колонки ([1][0])
	board[0][2] = "X" //X — в верхний правый угол ([0][2])

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
