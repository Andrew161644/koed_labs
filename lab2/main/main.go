package main

import (
	"fmt"
	"log"
	"math"
)

//1 y
//2 3 4 5 6 7 8 x
func main() {
	ex1()
	ex2()
}
func ex1() {
	var str = ""
	var matr1 = fileReader("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\data1.txt")
	str += showMatrix2nd("Исходная", matr1)
	var xmatr = excludeColumn(matr1, 0)
	var y = getColumnAsMatrix(matr1, 0)
	for i, _ := range xmatr {
		xmatr[i] = append(xmatr[i], 1)
	}
	str += showMatrix2nd("Y", y)

	str += showMatrix2nd("Значения X", xmatr)

	var tr = makeT(xmatr)
	var a = mnog(mnog(makeOp(mnog(tr, xmatr)), tr), y)
	str += showMatrix2nd("Вектор alpha", a)
	var yR = mnog(xmatr, a)
	str += showMatrix2nd("Вектор yˆ", yR)
	var zn = getColumnZn(yR, 0)
	str += fmt.Sprint("Среднее по yˆ ", zn, "\n")
	str += fmt.Sprint("Среднее по y ", getColumnZn(y, 0), "\n")
	str += fmt.Sprint("R= ", createR(y, yR, zn), "\n")
	str += "Коэффициент детерминации изменяется в пределах от 0 до 1. Он показывает,\n" +
		"как велика доля объясненной дисперсии в общей дисперсии, какая часть\n" +
		"общей дисперсии может быть объяснена зависимостью переменной y от\n" +
		"переменных x1, x2,…, xm."
	writeTofile("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\out.txt", str)
	log.Println("Completed")
}
func ex2() {
	var str = ""
	var matr1 = fileReader("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\data2.txt")
	str += showMatrix2nd("Исходная", matr1)
	var xmatr = excludeColumn(matr1, 0)
	var y = getColumnAsMatrix(matr1, 0)
	for i, _ := range xmatr {
		xmatr[i] = append(xmatr[i], 1)
	}
	str += showMatrix2nd("Y", y)
	str += showMatrix2nd("Значения X", xmatr)
	var tr = makeT(xmatr)
	var a = mnog(mnog(makeOp(mnog(tr, xmatr)), tr), y)
	str += showMatrix2nd("Вектор alpha", a)
	var yR = mnog(xmatr, a)
	str += showMatrix2nd("Вектор yˆ", yR)
	var zn = getColumnZn(yR, 0)
	str += fmt.Sprint("Среднее по yˆ ", zn, "\n")
	str += fmt.Sprint("Среднее по y ", getColumnZn(y, 0), "\n")
	str += fmt.Sprint("R= ", createR(y, yR, zn), "\n")
	str += "Коэффициент детерминации изменяется в пределах от 0 до 1. Он показывает,\n" +
		"как велика доля объясненной дисперсии в общей дисперсии, какая часть\n" +
		"общей дисперсии может быть объяснена зависимостью переменной y от\n" +
		"переменных x1, x2,…, xm."
	writeTofile("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\out2.txt", str)
	log.Println("Completed")
}

func createR(y [][]float64, yR [][]float64, yZ float64) float64 {
	res := 0.0
	sumE := 0.0
	sumY := 0.0
	for i := range y {
		sumE += math.Pow(y[i][0]-yR[i][0], 2)
		sumY += math.Pow(y[i][0]-yZ, 2)
	}

	res = 1 - sumE/sumY
	return res
}

func TestOb() {
	var matr2 = fileReader("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\data2.txt")
	var obr = makeOp(matr2)
	fmt.Println(showMatrix2nd("Обратная", obr))
}

func TestMnogY() {
	var matr2 = fileReader("C:\\Users\\Андрей\\go\\src\\koed_labs\\laba1\\lab2\\main\\data2.txt")
	fmt.Println(showMatrix2nd("", mnogVector(matr2, []float64{2, 3})))
}
