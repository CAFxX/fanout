package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64(3.14159)
    v3 := []int{1}
    var v4 string = "\n\t"
    var v5 uint16 = 0
    var v6 int16 = 1000
    v7 := v1[v4]
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (v5 * v5), F1: fi8, F2: 128}
        delete(v1, v4)
        _ = fi8
        _ = v9
    }
    fmt.Println(v4)
    v4 = v4
    if ((((((0.1 * 1e10) <= 3.14159) || true) && true) && false) && true) {
        _ = len(v1)
        v6 = (v6 - v6)
    }
    var v10 float64 = float64(0.1)
    v4 = "foo bar"
    if ((((0.1 * 1e10) * v2) <= 3.14159) && ((0 >= 499) || false)) {
        var v11 int = 1
        v3 = append(v3, (v7 << ((42) & 63)))
        _ = v11
    }
    if false {
        v12 := GenS2{F0: (v7 & v7), F1: 42}
        _ = v12
    }
    fmt.Println(v2, v7)
    var v13 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
