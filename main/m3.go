package main

import (
    "fmt"
)

func minOperations(b []int) int {
    operations := 0
    n := len(b)
    current := make([]int, n)

    // Step 1: Increment all elements by 1
    operations++
    for i := 0; i < n; i++ {
        current[i]++
    }

    for {
        // Check if we have reached the target
        allEqual := true
        for i := 0; i < n; i++ {
            if current[i] != b[i] {
                allEqual = false
                break
            }
        }
        if allEqual {
            break
        }

        // Find the first index where current and b differ
        start := 0
        for start < n && current[start] == b[start] {
            start++
        }

        // Find the end index of the segment to operate on
        end := start
        for end < n && current[end] != b[end] {
            end++
        }

        // Process the segment
        for i := start; i < end; i++ {
            diff := b[i] - current[i]
            // If the difference is greater than the current value
            if diff > 0 {
                current[i] *= 2
                if current[i] > b[i] {
                    current[i] = b[i] // Prevent overshooting
                }
            } else {
                current[i]++
            }
        }
        operations++
    }

    return operations
}

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        b := make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&b[j])
        }
        result := minOperations(b)
        fmt.Println(result)
    }
}