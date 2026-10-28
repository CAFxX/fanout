package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 int32 = 42
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 20; fi5++ {
        if v2 {
            if v2 {
                fmt.Println(v4)
            }
            var v6 uint64 = 2
            _ = v6
        }
        var v7 int16 = 865
        _ = fi5
        _ = v7
    }
    v2 = v2
    v2 = (611 < 233)
    fmt.Println(v3, v2)
    v8 := []int{0}
    if (1 > 2) {
        fmt.Println(v2, v2)
        var v9 float64 = v4
        v4 = -2.25
        _ = v9
    } else {
        v4 = (v4 * 1e10)
    }
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v10 := GenS1{F0: (v2 || v2), F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
