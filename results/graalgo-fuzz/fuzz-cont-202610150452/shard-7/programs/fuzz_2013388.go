package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }
func genFn0(p0 uint8, p1 int32, p2 int8) int16 {
    var v4 uint = 42
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint = 161
        _ = fi5
        _ = v6
    }
    if ((!(1 != -1)) || (462 >= 0)) {
        var v7 uint8 = 240
        var v8 int16 = -1
        _ = v7
        _ = v8
    } else {
        var v9 int8 = p2
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{fi10, (fi10 / ((-2147483648) | 1)), (fi10 ^ (^(fi10)))}
            v9 = p2
            v4 -= v4
            _ = fi10
            _ = v11
        }
        fmt.Println(p1, p1)
        _ = v9
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 351
}

func main() {
    v1 := []int{-190, 0, -37, 2147483647}
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v12 int8 = -1
    var v13 int64 = -9223372036854775808
    v1 = append(v1, -1)
    var v14 int = -9223372036854775808
    v15 := GenS1{F0: 2, F1: (!(!true)), F2: 959}
    v16 := GenS1{F0: 1, F1: true, F2: genFn0(10, -1, 1)}
    v12 -= (v12 / ((v12) | 1))
    v1 = append(v1, (v14 - (v14 >> ((18446744073709551615) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
