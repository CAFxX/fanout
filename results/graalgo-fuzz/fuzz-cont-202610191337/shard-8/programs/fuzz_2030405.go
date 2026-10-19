package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0() uint16 {
    var v5 float64 = 1e10
    v5 += 3.14159
    _ = v5
    return 65535
}

func main() {
    var v1 uint = 389
    v2 := []int{1, -9223372036854775808, 4294967295}
    v3 := []int{}
    var v4 uint8 = 128
    _ = len(v2)
    v1 += v1
    v1 += v1
    var v6 uint64 = 482
    v4 = (v4 - 255)
    var v7 uint16 = 65535
    var v8 int16 = 1000
    fmt.Println(v8, v7, v7)
    var v9 bool = (!(!(uint8(v1) < (v4 - uint8(v6)))))
    var v10 int8 = -1
    v11 := GenS1{F0: 967, F1: int(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
