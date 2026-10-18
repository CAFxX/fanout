package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint }
func genFn0() int16 {
    v5 := []int{2147483647, -238}
    var v6 int8 = 42
    var v7 int16 = -1
    v6 -= (v6 + v6)
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []uint32{1}
    var v2 bool = ((0 <= 1000) && true)
    var v3 uint = 838
    var v4 float32 = 0.1
    v4 += v4
    if v2 {
        fmt.Println(v2, v4)
    }
    fmt.Println(v3)
    var v8 string = ("a" + "a")
    fmt.Println(v4, v8)
    v9 := GenS1{F0: (v3 % ((v3) | 1)), F1: 0}
    var v10 int8 = 127
    v11 := GenS2{F0: v10, F1: (^(v3)), F2: 18446744073709551615}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
