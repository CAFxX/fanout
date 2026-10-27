package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = (1.5 - 0.0)
    var v2 string = (("a" + "a") + "foo bar")
    v3 := make(map[uint32]uint32)
    var v4 float64 = -2.25
    var v5 float64 = -2.25
    v6 := GenS1{F0: 42, F1: -193}
    if (((true || false) && true) || (0 == 42)) {
        v5 = 0.1
        switch 4294967295 {
        case -2147483648:
            v7 := GenS2{F0: 4294967295, F1: 1}
            _ = v7
        }
        var v8 int8 = 0
        _ = v8
    } else {
        v9 := []int{-539}
        var v10 bool = (0 != 374)
        _ = v10
        _ = v9
    }
    switch "a" {
    case "a":
        fmt.Println(v2)
    case "\n\t":
        if false {
            fmt.Println(v1, v5, v4)
        }
    case "hello":
        v2 = ("" + (("hello" + ("" + ("" + v2))) + ""))
    default:
        v4 = (0.1 + float64(v1))
    }
    fmt.Println(v1, v4)
    switch v2 {
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        var v11 bool = true
        _ = v11
    }
    v5 = (1.5 / (v4 - 1.5))
    v4 = (float64(3.14159) + -2.25)
    v2 = v2
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
