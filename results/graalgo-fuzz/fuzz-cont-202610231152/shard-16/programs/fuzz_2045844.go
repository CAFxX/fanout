package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }

func main() {
    v1 := []int{42, -9223372036854775808, -573, 1}
    var v2 int16 = -1
    v3 := make(map[string]bool)
    var v4 uint32 = 258
    var v5 int8 = 42
    v6 := GenS1{F0: 42, F1: v5, F2: 0}
    var v7 uint64 = 331
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v7, v7)
        var v9 uint = 18446744073709551615
        fmt.Println(v4)
        _ = fi8
        _ = v9
    }
    v4 -= 42
    fmt.Println(v5, v5, v4)
    if ((float32((0.0 + float64((((0.0 + (1.5 / float32(-2.25))) / 1.0) + 1e10)))) - (-2.25 - (-2.25 * 1.5))) == 3.14159) {
        var v10 uint = 18446744073709551615
        _ = v10
    }
    v7 -= (v7 + (^(v7)))
    v11 := GenS1{F0: -418, F1: (v5 >> ((42) & 63)), F2: -49}
    v12 := GenS1{F0: 2147483647, F1: -104, F2: 9223372036854775807}
    v2 = v2
    v13 := GenS2{F0: v2}
    if (1e10 >= (((1e10 + ((1e10 + (0.0 * -2.25)) * -2.25)) / 1e10) + 1e10)) {
        fmt.Println(v5)
    } else {
        v14 := GenS3{F0: "\n\t", F1: false, F2: -296}
        _ = v14
    }
    var v15 bool = (true || ((true && (true && false)) || (1 >= 16)))
    v5 = v5
    var v16 uint32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
