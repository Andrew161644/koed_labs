package main

import (
	"golang.org/x/exp/errors/fmt"
	"log"
	"strconv"
)

func main() {
	testExample()
	mainF()
}

func getMainDiagonalAsArray(matr [][]float64) []float64 {
	var res []float64
	var N = len(matr)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				res = append(res, matr[i][j])
			}
		}
	}
	return res
}

func SortMatrParallel(arr []float64, matr [][]float64) {
	var N = len(arr)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if arr[i] < arr[j] {
				var temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
				var tempMatrColumn = getColumnAsMatrix(matr, j)
				SetColumnAtMatrix(getColumnAsMatrix(matr, i), matr, j)
				SetColumnAtMatrix(tempMatrColumn, matr, i)
			}
		}
	}
}

func testExample() {
	var matr = fileReader("lab3\\main\\test.txt")
	var matrRes, tRes = makeIkobi(matr, 0.00000001)
	var arr = getMainDiagonalAsArray(matrRes)
	SortMatrParallel(arr, tRes)
	var result string
	result += fmt.Sprintln("Сортированные собственные числа: ", arr)
	result += fmt.Sprintln()
	result += showMatrix2nd("Отсортированные собственные векторы", tRes)
	log.Println(result)

}

func mainF() {
	log.Println("Process Started")
	var result string
	var matr = fileReader("lab3\\main\\data.txt")
	fmt.Println()
	var N = len(matr)
	result += "n = " + strconv.Itoa(len(matr)) + "\n"
	result += showMatrix2nd("Исходная", matr)

	var matrX = getXMatr(matr)
	result += showMatrix2nd("Стандартизированная матрица", matrX)
	var matrR = getRMatrix(matrX, N)
	result += showMatrix2nd("Корреляционная матрица", matrR)
	writeTofile("lab3\\main\\out.txt", result)
	var matrRes, tRes = makeIkobi(matrR, 0.00000001)
	result += showMatrix2nd("Результат преобразований из единичной - собственные векторы", tRes)
	result += showMatrix2nd("Ковариационная матрица главных компонент - собственные векторы", matrRes)
	var arr = getMainDiagonalAsArray(matrRes)
	SortMatrParallel(arr, tRes)
	result += fmt.Sprintln("Сортированные собственные числа: ", arr)
	result += fmt.Sprintln()
	result += showMatrix2nd("Отсортированные собственные векторы", tRes)
	result += showMatrix2nd("Проекции объектов на главные компоненты", calcProctions(matrX, tRes))
	var res, col = getIp(arr)
	result += fmt.Sprint("I(p`) = ", res, "\n")
	result += fmt.Sprint("Col = ", col, "\n")
	var resBool = calcD(matrR, N)
	if resBool {
		result += fmt.Sprint("Применение метода главных компонент\n" +
			"нецелесообразно\n")
	} else {
		result += fmt.Sprint("Применение метода главных компонент\n" +
			"целесообразно\n")
	}
	result += fmt.Sprintln()
	result += WriteInfo()
	writeTofile("lab3\\main\\out.txt", result)
	log.Println("Process finished")
}

func calcProctions(xmatr [][]float64, cmatr [][]float64) [][]float64 {
	var p = len(xmatr[0])
	res := make([][]float64, len(xmatr))
	for i := range res {
		res[i] = make([]float64, p)
	}

	for j := 0; j < p; j++ {
		var resultColumn = make([][]float64, len(xmatr))
		for i := range resultColumn {
			resultColumn[i] = []float64{0}
		}
		for k := 0; k < p; k++ {
			var xk = mnogCh(cmatr[j][k], getColumnAsMatrix(xmatr, k))
			resultColumn = sumMatrixSimpl(xk, resultColumn)
		}
		SetColumnAtMatrix(resultColumn, res, j)
	}
	return res
}

func sumMatrixSimpl(a [][]float64, b [][]float64) [][]float64 {
	var m = len(a)
	var n = len(a[0])

	res := make([][]float64, m)
	for i := range res {
		res[i] = make([]float64, n)
	}

	for i := range a {
		for j := range a[i] {
			res[i][j] = a[i][j] + b[i][j]
		}
	}
	return res
}

func getIp(arr []float64) (float64, int) {
	var Hp = 0.0
	var sumH = 0.0
	var Ip = 0.0
	var col int
	for i := range arr {
		sumH += arr[i]
	}

	for i := range arr {
		if Ip < 0.95 {
			Hp += arr[i]
			Ip = Hp / sumH
			col++
		} else {
			break
		}
	}
	return Ip, col
}
func WriteInfo() string {
	return "Если c 1k большой по модулю, то зависимость первой главной компоненты от\n" +
		"k x сильная, если коэффициент положительный, то зависимость\n" +
		"возрастающая, аналогично остальные случаи."
}
func test() {
	var xmatr = [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 1, 0},
	}
	var k = 2
	var cmatr = [][]float64{
		{7, 2, 8},
		{1, 3, 6},
		{6, 1, 0},
	}
	var j = 1
	var xk = getColumnAsMatrix(xmatr, k)
	var resultColumn = mnogCh(cmatr[j][k], xk)
	var cl = getColumnAsMatrix(xmatr, 0)
	log.Println(showMatrix2nd("", resultColumn))
	log.Println(showMatrix2nd("Test sum", sumMatrixSimpl(cl, resultColumn)))
}

// на вход r матрица
func calcD(matr [][]float64, N int) bool {
	var sum float64
	for i := range matr {
		for j := range matr[i] {
			if i != j {
				sum += matr[i][j]
			}
		}
	}
	return float64(N)*sum <= 69.9568
}
