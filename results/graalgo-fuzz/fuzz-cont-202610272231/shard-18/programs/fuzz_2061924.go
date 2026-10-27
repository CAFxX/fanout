package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
func genFn0(p0 int8) uint32 {
    v6 := GenS1{F0: 2, F1: (true && false), F2: 127}
    p0 -= -29
    var v7 int32 = 2147483647
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -1
    v2 := []int{1, -272, 159, 1}
    var v3 uint8 = 255
    var v4 int = 42
    v5 := make(map[string]string)
    delete(v5, "a")
    v8 := GenS1{F0: 42, F1: (0 != 100), F2: -128}
    switch v4 {
    case 826:
        fmt.Println(v4)
    }
    fmt.Println(v4, v1)
    if true {
        var v9 uint32 = genFn0(8)
        _ = v9
    }
    var v10 int32 = -2147483648
    var v11 int32 = (^(v10))
    v10 += 852
    v12 := GenS1{F0: 0, F1: true, F2: 95}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
