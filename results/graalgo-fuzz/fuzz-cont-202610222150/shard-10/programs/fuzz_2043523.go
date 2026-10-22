package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{4294967295, 9223372036854775807}
    var v3 int64 = -217
    v4 := make(map[uint32]int)
    v3 += v3
    delete(v4, 1)
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v5 int64 = v3
    v5 -= 42
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        if false {
            if (true && false) {
                var v7 float64 = ((1.5 + 0.0) / (3.14159 * (((0.0 / ((((0.0 / 1.0) - -2.25) / float64(-2.25)) - 0.1)) - 3.14159) * -2.25)))
                v5 += (v5 | (v3 & (v5 & v5)))
                _ = v7
            }
            var v8 float64 = 3.14159
            if len(v2) > 0 {
                v2[0] = fi6
            }
            _ = v8
        }
        _ = len(v2)
        _ = fi6
    }
    v3 += v5
    fmt.Println(v5, v3)
    v9 := GenS1{F0: -1, F1: 4294967295}
    v5 += -908
    if true {
        var v10 uint8 = 128
        var v11 int = 689
        _ = v10
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
