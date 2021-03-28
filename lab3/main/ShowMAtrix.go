package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strings"
)

func showMatrix2nd(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	var s = fmt.Sprintf("%s:\n%.2f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
	s = deleteChars(s)
	return s
}
func showMatrix0nd(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	var s = fmt.Sprintf("%s:\n%.1f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
	s = deleteChars(s)
	return s
}
func showMatrix(name string, matr [][]float64) string {
	A := mat.NewDense(len(matr), len(matr[0]), nil)
	for i := range matr {
		A.SetRow(i, matr[i])
	}

	var s = fmt.Sprintf("%s:\n%f\n\n", name, mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
	s = deleteChars(s)
	return s
}
func deleteChars(str string) string {
	str = strings.ReplaceAll(str, "⎡", "")
	str = strings.ReplaceAll(str, "⎢", "")
	str = strings.ReplaceAll(str, "⎥", "")
	str = strings.ReplaceAll(str, "⎣", "")
	str = strings.ReplaceAll(str, "⎤", "")
	str = strings.ReplaceAll(str, "⎦", "")

	return str
}
