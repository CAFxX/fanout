package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8, p1 uint) uint8 {
    fmt.Println(p1, p0)
    var v5 int = -2147483648
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 uint64) bool {
    v6 := GenS1{F0: 42}
    fmt.Println(p0)
    p0 = 42
    v7 := GenS2{F0: -128, F1: 446}
    _ = p0
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (1e10 - (1e10 * 3.14159))
    var v3 int = 855
    v4 := make(map[string]int)
    v8 := GenS1{F0: -82}
    v3 = v3
    v3 += v3
    fmt.Println(v2)
    var v9 string = "\n\t"
    var v10 bool = (v9 != "")
    var v11 uint32 = 2147483648
    v12 := []int{v3, v3, (v3 ^ v3)}
    switch (v3 * 32767) {
    case -96:
        var v13 uint16 = 42
        _ = v13
    }
    v14 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
