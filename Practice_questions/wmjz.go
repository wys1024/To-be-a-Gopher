package main

import (
    "fmt"
)

func main() {
    //读 n
    var n int
    fmt.Scanf("%d", &n)

    //读矩阵
    matrix := make([][]int, n)
    for i := range matrix{
        matrix[i] = make([]int, n)
        for j := range matrix{
            fmt.Scanf("%d", &matrix[i][j])
        }
    }

    //创建前缀和
    prefixmax := make([][]int, n+1)
    for i := range prefixmax{
        matrix[i] = make([]int, n+1)
    }

    //创建存储完美矩阵值
    prefectmatrix := make([]int, n+1)


    //计算前缀和
    for i := 1; i < n; i++{
        for j := 1; j < n; j++{
            prefectmatrix[i][j] = prefectmatrix[i][j-1] + prefectmatrix[i-1][j] - prefectmatrix[i-1][j-1] + matrix[i][j]

        }


    }























}