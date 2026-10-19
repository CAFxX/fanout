package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 string }

func main() {
    var v1 string = ("foo bar" + (("" + "\n\t") + ""))
    var v2 int32 = 981
    var v3 float64 = 0.0
    var v4 uint8 = 2
    fmt.Println(v1, v4, v1)
    v5 := []int{-29, -744}
    fmt.Println(v3, v3, v3)
    v5 = append(v5, 0)
    if (v1 >= v1) {
        if true {
            v5 = append(v5, 1)
        }
    } else {
        var v6 int = 9223372036854775807
        _ = v6
    }
    var v7 string = "foo bar"
    v8 := GenS1{F0: 42, F1: v2, F2: v1}
    if true {
        fmt.Println(v2, v7)
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = 0.1
        var v11 int = (fi9 / (((fi9 ^ (fi9 * fi9))) | 1))
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
