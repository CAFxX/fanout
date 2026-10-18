package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 2
    var v3 bool = ((!(1e10 == 0.0)) || ((3.14159 < (0.0 + 0.1)) && true))
    v4 := make(map[string]int)
    var v5 float32 = float32(((3.14159 / (-2.25 / (1.5 / ((-2.25 * 1.5) - 0.0)))) + ((0.0 + 1e10) / 1e10)))
    v2 = (v2 % ((v2) | 1))
    if (32 < 2147483647) {
        v6 := GenS1{F0: 763, F1: 32768}
        var v7 int8 = 0
        v8 := GenS1{F0: 332, F1: -2147483648}
        _ = v6
        _ = v7
        _ = v8
    } else {
        v3 = (0.1 <= 1e10)
    }
    v9 := GenS2{F0: 1}
    var v10 int8 = 0
    fmt.Println(v10)
    var v11 uint64 = 42
    v12 := []int{85}
    if len(v12) > 0 {
        v12[0] = 0
    }
    delete(v1, "")
    v13 := GenS3{F0: -9223372036854775808, F1: (^(v11)), F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
