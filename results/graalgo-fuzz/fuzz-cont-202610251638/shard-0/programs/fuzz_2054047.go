package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 int, p1 int64) int8 {
    var v7 bool = false
    v7 = ((42 == 2147483647) && (4294967295 > 4294967295))
    _ = p0
    _ = p1
    _ = v7
    return -57
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 99
    v3 := []uint32{}
    var v4 int64 = 0
    var v5 string = ("hello" + "hello")
    v6 := []uint32{}
    fmt.Println(v4, v2, v4)
    v4 = v4
    fmt.Println(v5)
    v5 = v5
    _ = len(v1)
    v8 := GenS1{F0: 439, F1: 1, F2: v2}
    var v9 int = 204
    var v10 string = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
