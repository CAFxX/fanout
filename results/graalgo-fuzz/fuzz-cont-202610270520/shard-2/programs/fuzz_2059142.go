package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }

func main() {
    var v1 uint = 255
    var v2 float64 = 0.0
    var v3 int16 = 32767
    var v4 int32 = 2147483647
    var v5 float64 = -2.25
    var v6 int8 = 1
    var v7 uint16 = 559
    v7 = v7
    v8 := []int{-713, 370}
    if (!(((0.0 + 1e10) < 1.5) && (42 <= 127))) {
        var v9 uint8 = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            v11 := GenS1{F0: "hello", F1: -182, F2: (v4 + (-(v4)))}
            var v12 float64 = v2
            _ = v11
            _ = v12
        }
        v7 = (v7 - v7)
        _ = fi10
    }
    var v13 uint64 = 685
    var v14 int32 = (v4 ^ (-(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
