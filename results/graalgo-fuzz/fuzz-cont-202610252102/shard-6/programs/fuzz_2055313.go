package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 int16) int8 {
    fmt.Println(p1, p0)
    v7 := GenS1{F0: 32767, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v7
    return -128
}

func main() {
    var v1 float64 = float64(float32(1e10))
    var v2 uint32 = 0
    var v3 int64 = 1
    var v4 int8 = 79
    v5 := make(map[uint32]bool)
    var v6 int8 = 68
    v8 := []int{561}
    if (306 >= -34) {
        v9 := GenS2{F0: 1, F1: 470, F2: 780}
        _ = v9
    }
    var v10 int64 = v3
    fmt.Println(v2, v6, v1)
    v10 -= (v3 / (((v3 & 966)) | 1))
    var v11 bool = ((!true) || (42 >= 548))
    var v12 uint = 65535
    _ = len(v5)
    var v13 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
