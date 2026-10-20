package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 int64 = -1
    var v2 uint = 2
    var v3 int64 = (v1 % ((v1) | 1))
    var v4 bool = true
    v2 += (v2 % ((v2) | 1))
    if false {
        v2 = v2
        v2 -= v2
        v5 := []int{2147483647}
        _ = v5
    } else {
        v6 := GenS1{F0: -2147483648, F1: 127}
        v7 := []int{-2147483648, 913, 32768}
        fmt.Println(v1)
        _ = v6
        _ = v7
    }
    var v8 uint64 = 4294967295
    v9 := []int{0}
    var v10 uint8 = 0
    var v11 bool = ((v8 / (((^(v8))) | 1)) == (v8 / ((0) | 1)))
    v1 = 379
    switch "" {
    case "a":
        var v12 string = ("a" + ("\n\t" + "a"))
        _ = v12
    }
    if (false || v11) {
        v9 = append(v9, 4294967295)
    } else {
        var v13 int8 = 121
        _ = v13
    }
    var v14 int16 = 521
    var v15 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
