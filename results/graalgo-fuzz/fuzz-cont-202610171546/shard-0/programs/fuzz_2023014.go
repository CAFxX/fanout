package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32 }
func genFn0() int16 {
    var v4 int64 = 7
    v5 := []int{}
    var v6 int64 = v4
    _ = v4
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 uint64, p1 int64) uint8 {
    var v7 uint64 = p0
    v8 := GenS1{F0: false, F1: false}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 45
}

func main() {
    var v1 int8 = -107
    var v2 int32 = 0
    v3 := []int{32767}
    v2 = 42
    v9 := GenS2{F0: v2}
    var v10 int64 = 368
    fmt.Println(v2, v2)
    fmt.Println(v10, v10)
    var v11 uint8 = genFn1(2, (v10 & 1))
    var v12 uint8 = (v11 ^ (v11 | 68))
    var v13 int16 = 685
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
