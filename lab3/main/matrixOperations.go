package main

import (
	"log"
	"math"
)

// умножение двух матриц
func mnog(matr1 [][]float64, matr2 [][]float64) [][]float64 {
	var m = len(matr1)
	var n = len(matr1[0])
	var q = len(matr2[0])

	res := make([][]float64, m)
	for i := range res {
		res[i] = make([]float64, q)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			for k := 0; k < n; k++ {
				res[i][j] += matr1[i][k] * matr2[k][j]
			}
		}
	}

	return res
}

// транспонирование матрицы
func makeT(matr [][]float64) [][]float64 {
	var m = len(matr)
	var n = len(matr[0])

	res := make([][]float64, n)
	for i := range res {
		res[i] = make([]float64, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matr[i][j]
		}
	}

	return res
}

// детерминант матрицы
func det(a [][]float64) float64 {
	if len(a) == 1 {
		return a[0][0]
	}
	sign, d := 1, float64(0)
	for i, x := range a[0] {
		var v = det(excludeColumn(a[1:], i))
		d += float64(sign) * x * v
		sign *= -1
	}
	return d
}

// исключение колонки
func excludeColumn(a [][]float64, i int) [][]float64 {
	b := make([][]float64, len(a))
	n := len(a[0]) - 1
	for j, row := range a {
		r := make([]float64, n)
		copy(r[:i], row[:i])
		copy(r[i:], row[i+1:])
		b[j] = r
	}
	return b
}

// исключение ряда
func excludeRow(a [][]float64, r int) [][]float64 {
	var b [][]float64
	for i, row := range a {
		if i == r {
			continue
		}
		b = append(b, row)
	}
	return b
}

// обратная матрицы
func makeOp(matr [][]float64) [][]float64 {
	var m = len(matr)
	var n = len(matr[0])

	res := make([][]float64, m)
	for i := range res {
		res[i] = make([]float64, n)
	}
	var det = det(matr)
	tr := makeT(matr)
	//fmt.Println(showMatrix2nd("Транспонированная",tr))
	matr = makeAlgAdd(tr)
	//fmt.Println(det)
	//fmt.Println(showMatrix2nd("Алг доп",matr))
	var ch = 1.0 / det
	res = mnogCh(ch, matr)
	//fmt.Println(showMatrix2nd("Обратная матрица",res))
	return res
}

// умножение матрицы на число
func mnogCh(ch float64, matr [][]float64) [][]float64 {
	var m = len(matr)
	var n = len(matr[0])

	res := make([][]float64, m)
	for i := range res {
		res[i] = make([]float64, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = matr[i][j] * ch
		}
	}
	return res
}

// расчет алг дополнений
func makeAlgAdd(matr [][]float64) [][]float64 {
	var m = len(matr)
	var n = len(matr[0])

	res := make([][]float64, m)
	for i := range res {
		res[i] = make([]float64, n)
	}

	for i := range matr {
		for j := range matr[i] {
			var temp = excludeColumn(excludeRow(matr, i), j)
			res[i][j] = math.Pow(-1, float64(i+j)) * det(temp)
		}
	}
	return res
}

// умножение на вектор
func mnogVector(matr [][]float64, vector []float64) [][]float64 {
	var column [][]float64
	for i := range vector {
		column = append(column, []float64{vector[i]})
	}
	return mnog(matr, column)
}

// получение стобца
func getColumnAsMatrix(matr [][]float64, column int) [][]float64 {
	var res [][]float64
	for _, val := range matr {
		res = append(res, []float64{val[column]})
	}
	return res
}

func SetColumnAtMatrix(column [][]float64, matr [][]float64, position int) {
	for i := range matr {
		for j := range matr[i] {
			if j == position {
				matr[i][j] = column[i][0]
			}
		}
	}
}

// получение среднего по колонне
func getColumnZn(matr [][]float64, column int) float64 {
	var res = 0.0
	for i := range matr {
		res += matr[i][column]
	}
	return res / float64(len(matr))
}

func findS2(matr [][]float64, column int) float64 {
	var res float64
	var z = getColumnZn(matr, column)
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
		res += (matr[k][line] - getColumnZn(matr, line)) * (matr[k][column] - getColumnZn(matr, column))
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
			var res = (matr[i][j] - getColumnZn(matr, j)) / math.Sqrt(findS2(matr, j))
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

func makeEdMatr(N int, M int) [][]float64 {
	var T0 [][]float64
	for i := 0; i < N; i++ {
		var newLine []float64
		for j := 0; j < M; j++ {
			if i == j {
				newLine = append(newLine, 1)
				continue
			}
			newLine = append(newLine, 0)
		}
		T0 = append(T0, newLine)
	}
	return T0
}
