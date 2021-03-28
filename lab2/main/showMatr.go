package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func ShowMatr(matr [][]float64) {
	for i := range matr {
		for j := range matr[i] {
			fmt.Print(matr[i][j], " ")
		}
		fmt.Println()
	}
}

func showMatrix2nd(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	return fmt.Sprintf("%s:\n%.2f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
}
