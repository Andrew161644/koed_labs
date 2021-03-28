package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func getZ(arr []float64) float64 {
	var sum float64
	for _, i := range arr {
		sum += i
	}
	return sum / (float64(len(arr)))
}

func getZColumn(matr [][]float64, column int) float64 {
	return getZ(getColumn(matr, column))
}

func getColumn(matr [][]float64, column int) []float64 {
	var res []float64
	for _, v := range matr {
		res = append(res, v[column])
	}
	return res
}

func showMatrix2nd(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	return fmt.Sprintf("%s:\n%.2f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
}

func showMatrix0nd(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	return fmt.Sprintf("%s:\n%f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
}
