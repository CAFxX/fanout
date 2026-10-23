package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16, p1 uint) int32 {
    fmt.Println(p0, p0)
    v4 := GenS1{F0: 1}
    p0 -= 32767
    v5 := GenS2{F0: 42, F1: 42}
    p1 = (p1 ^ (p1 | (p1 << (((p1 + 792)) & 63))))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 bool = (32768 <= -1)
    var v2 float64 = 0.0
    var v3 bool = (255 >= 1)
    var v6 uint64 = 801
    var v7 int = 127
    v8 := []int{v7}
    if false {
        fmt.Println(v1, v1, v3)
    }
    var v9 int16 = -1
    var v10 uint8 = 145
    v2 -= 0.1
    _ = len(v8)
    var v11 bool = true
    v6 = 401
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
