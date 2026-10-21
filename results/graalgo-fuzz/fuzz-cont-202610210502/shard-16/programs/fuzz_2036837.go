package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = 809
    var v2 int64 = -977
    v3 := []string{}
    v4 := make(map[int]bool)
    v5 := make(map[int]int)
    v2 -= -903
    v2 -= v2
    v1 = (v1 + (v1 * v1))
    v6 := GenS1{F0: -1}
    v2 += -1
    delete(v4, 32767)
    if true {
        var v7 float64 = 1e10
        _ = v7
    } else {
        if true {
            v3 = append(v3, "\n\t")
            v8 := GenS1{F0: 127}
            _ = v8
        }
        fmt.Println(v2)
    }
    v1 -= (v1 | (v1 / ((v1) | 1)))
    if (127 <= 0) {
        v1 = (^(v1))
        v9 := GenS2{F0: 761}
        _ = v9
    }
    v3 = append(v3, "")
    var v10 uint = 2
    var v11 uint = (v10 | v10)
    v11 = v11
    v12 := GenS2{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
