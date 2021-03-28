package main

import (
	"math"
)

func makeIkobi(matr [][]float64, ep float64) ([][]float64, [][]float64) {

	var N = len(matr)
	// шаг 1
	var T0 = makeEdMatr(N, N)
	// шаг 2
	var k = 0
	var bar float64
	bar = CalcFirstBar(matr)
	var a0 = bar
	for chekNotDioganal(matr, ep, a0) {
		//вычисляем новую преграду
		k++
		// шаг 3
		check, p, q := findMaxNotDMoreBar(matr, bar)
		if !check {
			// шаг 4
			var c, s = calcForPart4_1(matr, p, q)
			matr, T0 = calcForPart4_2(matr, T0, p, q, c, s)
		}
		// находим новую преграду
		bar = bar / math.Pow(float64(N), 2)
		//log.Println("Новая преграда: ", bar)
	}
	return matr, T0
}

// Вычисляем первую преграду
func CalcFirstBar(matr [][]float64) float64 {
	var N = len(matr)
	var res float64
	for j := 1; j < N; j++ {
		for i := 0; i < j-1; i++ {
			res += matr[i][j] * matr[i][j]
		}
	}
	return 1.0 / float64(N) * math.Sqrt(2.0*res)
}

func chekNotDioganal(matr [][]float64, ep float64, a0 float64) bool {
	var N = len(matr)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j {
				if math.Abs(matr[i][j]) > ep*a0 {
					//log.Println(matr[i][j])
					return true
				}
			}
		}
	}
	return false
}

func findMaxNotDMoreBar(matr [][]float64, bar float64) (bool, int, int) {
	var max = 0.0
	var p int
	var q int
	var N = len(matr)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if math.Abs(matr[i][j]) > max && math.Abs(matr[i][j]) > bar && i != j {
				max = math.Abs(matr[i][j])
				p = i
				q = j
			}
		}
	}
	//log.Println("Максимальный элемент: ", max)
	if max == -10000000.0 {
		return true, -1, -1
	}
	return false, p, q
}

// нахождение c и s
func calcForPart4_1(matr [][]float64, p int, q int) (c float64, s float64) {
	var y = (matr[p][p] - matr[q][q]) / 2
	var x float64
	if y == 0 {
		x = -1
	} else {
		x = -sign(y) * (matr[p][q] / (math.Sqrt(matr[p][q]*matr[p][q] + y*y)))
	}
	var sv = x / (math.Sqrt(2 * (1 + math.Sqrt(1-x*x))))
	var cv = math.Sqrt(1 - sv*sv)
	return cv, sv
}

// функция определения знака
func sign(a float64) float64 {
	if a > 0 {
		return 1
	}
	return -1
}

// преобразование исходной матрицы
func calcForPart4_2(matr [][]float64, T [][]float64, p int, q int, c float64, s float64) (matrv [][]float64, Tv [][]float64) {
	//log.Println("P = ", p)
	//log.Println("Q = ", q)
	//log.Println(showMatrix2nd("Исходная A: ",matr))
	//log.Println(showMatrix2nd("Исходная T: ",T))
	var n = len(matr)
	for i := 0; i < n; i++ {
		if i != p && i != q {
			var Z1 = matr[i][p]
			var Z2 = matr[i][q]
			matr[q][i] = Z1*s + Z2*c
			matr[i][q] = matr[q][i]
			matr[i][p] = Z1*c - Z2*s
			matr[p][i] = matr[i][p]
		}
		T = calcForPart4_2_1(T, i, p, q, s, c)
	}
	var Z5 = s * s
	var Z6 = c * c
	var Z7 = s * c
	var V1 = matr[p][p]
	var V2 = matr[p][q]
	var V3 = matr[q][q]
	matr[p][p] = V1*Z6 + V3*Z5 - 2*V2*Z7
	matr[q][q] = V1*Z5 + V3*Z6 + 2*V2*Z7
	matr[p][q] = (V1-V3)*Z7 + V2*(Z6-Z5)
	matr[q][p] = matr[p][q]
	//log.Println(showMatrix2nd("Преобразованная A: ",matr))
	//log.Println(showMatrix2nd("Преобразованная T: ",T))
	return matr, T
}

// преобразование матрицы T
func calcForPart4_2_1(T [][]float64, i int, p int, q int, s float64, c float64) [][]float64 {
	var Z3 = T[i][p]
	var Z4 = T[i][q]
	T[i][q] = Z3*s + Z4*c
	T[i][p] = Z3*c - Z4*s
	return T
}
