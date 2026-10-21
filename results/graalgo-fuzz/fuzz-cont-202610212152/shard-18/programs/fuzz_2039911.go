package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
    }
    v8 := []int{-9223372036854775808, -947, 0}
    var v9 int16 = 611
    _ = v8
    _ = v9
    return 73
}
func genFn1(p0 uint64) int32 {
    var v10 uint = 42
    v11 := GenS1{F0: (p0 / ((p0) | 1)), F1: 128, F2: uint(p0)}
    fmt.Println(v10, v10)
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 float64 = (0.0 * (1.5 - (0.0 / 3.14159)))
    var v2 uint = 18446744073709551615
    var v3 int16 = -1
    var v4 int32 = 42
    v5 := make(map[int]bool)
    fmt.Println(v1, v4)
    v2 += 18446744073709551615
    if false {
        v12 := GenS1{F0: 2, F1: 36, F2: (^(v2))}
        v2 += (v2 >> ((2) & 63))
        _ = v12
    } else {
        var v13 int32 = v4
        _ = v13
    }
    var v14 uint64 = 0
    v15 := v5[4294967295]
    delete(v5, 1)
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 433
    var v19 uint = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
