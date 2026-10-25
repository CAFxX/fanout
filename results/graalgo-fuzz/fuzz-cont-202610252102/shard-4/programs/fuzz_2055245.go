package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 int32 = -125
    var v2 string = "a"
    var v3 uint16 = 350
    v4 := []int{}
    v3 -= v3
    v3 -= (v3 << ((42) & 63))
    if ((0.1 != float64(3.14159)) && true) {
        v3 = v3
        v5 := GenS1{F0: 42, F1: 904}
        _ = v5
    } else {
        var v6 int = 32768
        _ = v6
    }
    if (!false) {
        v7 := GenS1{F0: 42, F1: 42}
        _ = v7
    }
    v3 = v3
    var v8 bool = ((false || false) || (42 != 0))
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint8 = 1
        _ = fi9
        _ = v10
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
