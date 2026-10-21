package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint = 316
    var v2 string = ("hello" + "\n\t")
    var v3 uint = 1
    var v4 float64 = 1e10
    switch -565 {
    case -389:
        v5 := GenS1{F0: "\n\t"}
        _ = v5
    }
    var v6 string = (v2 + "a")
    v7 := GenS2{F0: 42, F1: -1}
    v3 = v3
    var v8 int32 = -416
    for fi9 := 0; fi9 < 8; fi9++ {
        v8 = (v8 / (((-(v8))) | 1))
        _ = fi9
    }
    v10 := []int{128, -924, 1, 556}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
