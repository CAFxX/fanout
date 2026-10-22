package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0() int16 {
    var v5 float64 = (-2.25 / 1e10)
    v6 := GenS1{F0: 32768, F1: false, F2: -976}
    var v7 float64 = v5
    _ = v5
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []int{128, 2147483647, 239, 4294967295}
    var v3 uint32 = 1
    v4 := make(map[string]int)
    var v8 bool = (485 != 2)
    v9 := GenS2{F0: -9223372036854775808, F1: 100}
    fmt.Println(v8, v8)
    var v10 string = "hello"
    v3 += v3
    fmt.Println(v8)
    if false {
        for fi11 := 0; fi11 < 8; fi11++ {
            v8 = (419 <= 873)
            _ = fi11
        }
    }
    v12 := GenS1{F0: -101, F1: ((42 >= 946) || v8), F2: -985}
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := v4[(v10 + "\n\t")]
        _ = fi13
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
