package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    var v6 uint32 = 831
    _ = p0
    _ = v6
    return 100
}
func genFn1(p0 int8, p1 int64) string {
    p1 += (p1 / (((p1 | p1)) | 1))
    var v7 uint = 600
    p1 = (p1 & 0)
    v7 -= v7
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 700
    v3 := []int{int(v2)}
    var v4 int = 128
    v5 := make(map[string]bool)
    var v8 int32 = 543
    var v9 bool = (true && true)
    var v10 uint32 = 305
    var v11 int16 = 0
    var v12 int16 = v11
    var v13 uint8 = 255
    _ = len(v1)
    v14 := GenS1{F0: (^(v4)), F1: v9}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
    _ = v8
    _ = v9
}
