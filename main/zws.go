package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array1 float浮点型一维数组 第一个有序数组
 * @param array2 float浮点型一维数组 第二个有序数组
 * @return float浮点型
 */
func find_median( array1 []float32 ,  array2 []float32 ) float32 {
    // 合并两个有序数组
    merged := make([]float32, 0, len(array1)+len(array2))
    i, j := 0, 0
    for i < len(array1) && j < len(array2) {
        if array1[i] < array2[j] {
            merged = append(merged, array1[i])
            i++
        } else {
            merged = append(merged, array2[j])
            j++
        }
    }
    

    // 计算中位数
    n := len(merged)
    if n%2 == 0 {
        return (merged[n/2-1] + merged[n/2]) / 2
    } else {
        return merged[n/2]
    }
    
}