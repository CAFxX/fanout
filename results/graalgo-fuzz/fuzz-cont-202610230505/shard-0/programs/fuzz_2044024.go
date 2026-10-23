package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    v1 := []string{"", "", "hello"}
    var v2 int = 880
    v3 := []int{-595, (^(v2)), v2, v2}
    var v4 uint64 = 0
    var v5 uint64 = (v4 | v4)
    var v6 uint64 = v5
    v7 := GenS1{F0: v2, F1: -2147483648, F2: 613}
    if (4294967295 <= 0) {
        v5 += (v5 * 18446744073709551615)
        var v8 int64 = 9223372036854775807
        _ = v8
    } else {
        v9 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: -1, F2: (v5 / ((v5) | 1))}
        v10 := GenS2{F0: "a", F1: 945}
        _ = v10
        _ = v9
    }
    var v11 int16 = -1
    v12 := GenS2{F0: "", F1: 255}
    var v13 int32 = 1
    v4 -= 0
    var v14 float32 = (float32(float64((1e10 / float32((1e10 * float64(1.5)))))) + (0.0 + 0.1))
    var v15 int8 = -40
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
