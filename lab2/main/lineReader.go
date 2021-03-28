package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseString(content string) []float64 {

	if strings.Contains(content, ",") {
		content = strings.ReplaceAll(content, ",", ".")
	}

	var values []float64
	var s = ""

	for i, v := range content {
		val := string(v)

		if val != " " && !strings.Contains(val, "\t") {
			s += val
		}
		if val == " " || strings.Contains(val, "\t") {
			if s != "" {
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					log.Fatal(err)
				}
				values = append(values, f)
			}
			s = ""
		}
		if i == len(content)-1 && len(s) > 0 {
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, f)
		}
	}

	return values
}

func readLines(path string) ([]string, error) {
	s := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var str = scanner.Text()
		if str == "" || str == " " {
			continue
		}
		s = append(s, str)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s, err
}

func fileReader(path string) [][]float64 {
	lines, err := readLines(path)
	if err != nil {
		log.Fatal(err)
	}
	var values [][]float64
	for _, v := range lines {
		values = append(values, parseString(v))
	}
	return values
}

func writeTofile(path string, info string) {
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(info)
}

func writeInfo() string {
	var res = "Отличие от нуля коэффициента корреляции обусловлено случайными величинами \n" +
		"В стандартизованной ковариационная и корреляционной: в каждом столбце мат ожидание равно 0, оценка дисперсии равняется единице \n" +
		"У ковариационной матрицы по диагонали стоят дисперсии, у корреляционной матрице по диагонали стоят единицы \n" +
		"Поэтому в статистической матрице по диагонали деление на ноль \n"

	return res
}
