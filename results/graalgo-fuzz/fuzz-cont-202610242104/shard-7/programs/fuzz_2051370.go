package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 float32 = (((float32(3.14159) / (-2.25 / 1.0)) + 1.5) / (3.14159 / (0.1 * 1e10)))
    var v2 int8 = -13
    v3 := make(map[uint32]bool)
    v2 = (^(v2))
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int = (fi4 / (((fi4 * (-(fi4)))) | 1))
        _ = fi4
        _ = v5
    }
    fmt.Println(v1, v1)
    v2 -= 42
    v6 := []int{-57, 127, 0, 1}
    v1 = (v1 - 1e10)
    var v7 bool = (true || (-2.25 != -2.25))
    switch "foo bar" {
    case "foo bar":
        if true {
            v2 = v2
            fmt.Println(v1, v7, v1)
        } else {
            v6 = append(v6, 356)
        }
    default:
        delete(v3, 996)
    }
    if (-128 >= (v2 + v2)) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int = (fi8 / (((fi8 << ((4294967295) & 63))) | 1))
            delete(v3, 2147483647)
            fmt.Println(v1, v2, v1)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: v2}
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
