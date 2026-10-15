package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }

func main() {
    v1 := []int{9223372036854775807}
    v2 := []int{}
    var v3 int8 = 76
    var v4 int8 = v3
    v4 -= v4
    var v5 int = -1
    if (839 > 691) {
        _ = len(v2)
        var v6 float64 = ((1e10 * ((1e10 + 3.14159) + 3.14159)) * 1.5)
        _ = v6
    }
    v5 -= -614
    var v7 int32 = 0
    var v8 int64 = 42
    switch (v5 + v5) {
    case 244:
        fmt.Println(v7)
    }
    v9 := GenS1{F0: ("" + "a"), F1: 523, F2: (v5 / (((v5 - (v5 * 128))) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
