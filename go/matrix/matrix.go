package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	rows [][]int
	n    int
	m    int
}

func New(s string) (*Matrix, error) {
	res := make([][]int, 0)
	rows := strings.Split(s, "\n")
	var n int
	var cnt int
	for _, row := range rows {
		nums := strings.Split(row, " ")

		tmp := make([]int, 0)
		for _, num := range nums {
			if num == "" {
				continue
			}
			x, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, x)
		}
		if n == 0 && cnt == 0 {
			n = len(tmp)
			cnt++
		} else if n != len(tmp) {
			return nil, errors.New("The number of colums not equal")
		}
		res = append(res, tmp)
	}
	return &Matrix{rows: res, n: len(res), m: len(res[0])}, nil
}

func (matrix Matrix) Rows() [][]int {
	res := make([][]int, len(matrix.rows))
	for i, row := range matrix.rows {
		res[i] = make([]int, len(row))
		copy(res[i], row)
	}
	// copy(res, matrix.rows)
	return res
}

func (matrix Matrix) Cols() [][]int {
	r := matrix.Rows()
	res := make([][]int, 0)
	n, m := len(r), len(r[0])
	for j := 0; j < m; j++ {
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, r[i][j])
		}
		res = append(res, tmp)
	}
	return res
}

func (matrix Matrix) Set(r, c, val int) bool {
	if r >= matrix.n || r < 0 || c >= matrix.m || c < 0 {
		return false
	}
	matrix.rows[r][c] = val
	return true
}
