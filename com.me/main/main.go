package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	log.Println("Process Started")
	var result string
	var matr = fileReader("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\com.me\\main\\data.txt")
	result += showMatrix2nd("Средние", [][]float64{findS2ForAll(matr)})
	fmt.Println()
	var N = len(matr)
	result += "n = " + strconv.Itoa(len(matr)) + "\n"
	result += showMatrix2nd("Исходная", matr)

	var res = getKovarcAll(matr, N)
	result += showMatrix2nd("Ковариационная матрица", res)
	var matrX = getXMatr(matr)
	result += showMatrix2nd("Стандартизированная матрица", matrX)
	var matrR = getRMatrix(matrX, N)
	result += showMatrix2nd("Корреляционная матрица", matrR)
	var matrT = getTMatr(matrR, N)
	result += showMatrix2nd("Матрица статистики - по диагонали none", matrT)
	var matrH = getHMatrix(matrT, 1.994)
	result += showMatrix0nd("Матрица H1/H0 по диагонали -1", matrH)
	result += writeInfo()
	result += "a = 0.05\n"
	result += "t = 1.994\n"
	writeTofile("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\com.me\\main\\out.txt", result)
	log.Println(result)
}

func findS2(matr [][]float64, column int) float64 {
	var res float64
	var z = getZColumn(matr, column)
	for i := range matr {
		res += (matr[i][column] - z) * (matr[i][column] - z)
	}
	return res / float64(len(matr))
}

func findS2ForAll(matr [][]float64) []float64 {
	var res []float64
	for j := range matr[0] {
		res = append(res, findS2(matr, j))
	}
	return res
}

func getKovarc(matr [][]float64, line int, column int, N int) float64 {
	var res float64
	for k := range matr {
		res += (matr[k][line] - getZColumn(matr, line)) * (matr[k][column] - getZColumn(matr, column))
	}
	return res / float64(N)
}

func getKovarcAll(matr [][]float64, N int) [][]float64 {
	var res [][]float64
	for i := range matr[0] {
		var newLine []float64
		for j := range matr[i] {
			newLine = append(newLine, getKovarc(matr, i, j, N))
		}
		res = append(res, newLine)
	}
	return res
}

func getXMatr(matr [][]float64) [][]float64 {
	var res [][]float64
	for i := range matr {
		var newLine []float64
		for j := range matr[i] {
			var res = (matr[i][j] - getZColumn(matr, j)) / math.Sqrt(findS2(matr, j))
			newLine = append(newLine, res)
		}
		res = append(res, newLine)
	}
	return res
}

// на вход X матр
func getRMatrix(matr [][]float64, N int) [][]float64 {
	var res [][]float64

	for i := range matr[0] {
		var newLine []float64
		for j := range matr[i] {
			var sum float64
			for k := range matr {
				sum += matr[k][i] * matr[k][j]
			}
			sum /= float64(N)
			if sum >= 1 && i != j {
				log.Fatal("Error")
			}
			newLine = append(newLine, sum)
		}
		res = append(res, newLine)
	}
	return res
}

// на вход R матрица
func getTMatr(matr [][]float64, N int) [][]float64 {
	var res [][]float64
	for i := range matr[0] {
		var newLine []float64
		for j := range matr[i] {
			var zn = matr[i][j]
			if zn == 1 {
				fmt.Println(zn)
			}
			var cnt = zn * math.Sqrt(float64(N)) / (math.Sqrt(1 - zn*zn))
			if i == j {
				cnt = 0
			}
			newLine = append(newLine, cnt)
		}
		res = append(res, newLine)
	}
	return res
}

func getHMatrix(matr [][]float64, koeff float64) [][]float64 {
	for i := range matr {
		for j := range matr[i] {
			if math.Abs(matr[i][j]) >= koeff {
				matr[i][j] = 1
			} else {
				matr[i][j] = 0
			}
			if i == j {
				matr[i][j] = -1
			}
		}
	}
	return matr
}
