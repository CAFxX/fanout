package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0(p0 int64, p1 int8) int32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 237, F1: ((0.1 - 1.5) < 0.0)}
        _ = fi6
        _ = v7
    }
    var v8 float64 = 1.5
    v9 := GenS1{F0: -9223372036854775808, F1: false}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -385
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{343, 0, 2147483647, 128}
    var v3 bool = true
    var v4 uint8 = 127
    var v5 float32 = float32(float64(3.14159))
    var v10 float32 = 0.1
    v10 += float32(((3.14159 / 3.14159) + -2.25))
    var v11 string = "foo bar"
    var v12 uint64 = 1
    v13 := GenS2{F0: (v12 + (v12 | 18446744073709551615)), F1: int32(v12), F2: genFn0(42, 0)}
    fmt.Println(v10, v10)
    var v14 int32 = genFn0(42, -39)
    v2 = append(v2, int(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
