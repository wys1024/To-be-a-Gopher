package main

import "fmt"

func main() {
    var S string
    fmt.Scanf("%s", &S)

    x, y, direction := 0, 0, 0 // x, y 坐标, direction 指示当前方向 (0 - 北, 1 - 东, 2 - 南, 3 - 西)

    for _, c := range S {
        switch c {
        case 'W':
            switch direction {
            case 0:
                y++
            case 1:
                x++
            case 2:
                y--
            case 3:
                x--
            }
        case 'S':
            // 原地不动
        case 'A':
            direction = (direction + 3) % 4 // 逆时针旋转 90 度
        case 'D':
            direction = (direction + 1) % 4 // 顺时针旋转 90 度
        }
    }

    fmt.Printf("%d %d\n", x, y)
}