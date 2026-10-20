package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint8 }

func main() {
    var v1 bool = true
    var v2 int = -579
    var v3 int64 = 0
    var v4 int8 = 127
    var v5 float64 = 0.1
    v1 = (0 >= 227)
    v6 := GenS1{F0: 0}
    var v7 int64 = v3
    v8 := []int{(v2 >> ((340) & 63)), (v2 ^ (-(v2))), (v2 / ((v2) | 1)), v2}
    v4 = v4
    v9 := GenS1{F0: -951}
    var v10 uint32 = 398
    fmt.Println(v3, v7, v3)
    v11 := GenS2{F0: "", F1: 304, F2: 47}
    v12 := GenS3{F0: (!(!(379 >= uint64(v10)))), F1: 2}
    if (!(!(((v2 - 4294967295) != v2) || v1))) {
        v10 += 758
        v4 = (v4 & 87)
        var v13 uint64 = 18446744073709551615
        _ = v13
    }
    if (18446744073709551615 != 2) {
        var v14 int16 = 42
        var v15 uint16 = 42
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
