package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 string, p1 uint, p2 int) uint8 {
    fmt.Println(p1, p0)
    v4 := []int{(-(p2)), 32767, 333}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 82
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = (v1 / ((v1) | 1))
    v3 := []uint32{2147483647, 0, 219, 994}
    var v5 bool = false
    switch "a" {
    case "foo bar":
        v2 = v1
    case "hello":
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    case "\n\t":
        fmt.Println(v1, v2, v5)
    default:
        v6 := GenS1{F0: 65535, F1: -113}
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (("\n\t" + "a") + "foo bar") {
        case "a":
            _ = len(v3)
        case "hello":
            v1 = v1
        }
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if (2 >= 42) {
        _ = len(v3)
        var v9 int16 = -105
        fmt.Println(v1, v2, v1)
        _ = v9
    }
    fmt.Println(v1, v1)
    var v10 int8 = -109
    fmt.Println(v5, v5)
    v10 = 100
    var v11 int16 = 1
    var v12 float64 = (1e10 / 0.1)
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
}
