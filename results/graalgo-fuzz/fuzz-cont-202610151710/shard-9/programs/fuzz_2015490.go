package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }

func main() {
    var v1 uint32 = 65535
    var v2 int8 = 26
    v3 := make(map[int]bool)
    v4 := v3[238]
    if v4 {
        var v5 int = 32767
        _ = v5
    } else {
        v4 = v4
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(fi6)
            v2 += (^(v2))
            _ = fi6
        }
    }
    v4 = (!false)
    v7 := []int{-204, 42, -931, 128}
    var v8 int8 = (^(v2))
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v4 = (((^(fi10)) >= 1) || (1 != -1))
            var v11 uint = 65535
            var v12 uint = (^(v11))
            v13 := GenS1{F0: -523, F1: (fi9 / ((fi10) | 1)), F2: (v2 / ((-1) | 1))}
            v1 = v1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int = (^(fi14))
            v16 := v3[v15]
            _ = len(v3)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    v17 := v3[-305]
    _ = len(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
