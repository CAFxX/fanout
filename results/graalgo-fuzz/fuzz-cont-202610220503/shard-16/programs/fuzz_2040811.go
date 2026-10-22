package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    if true {
        v4 := GenS1{F0: -210, F1: 1}
        v5 := GenS1{F0: 296, F1: 1}
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v1 = v1
        if (1.5 == float64(0.0)) {
            v1 = v1
        }
        _ = fi6
    }
    var v7 uint16 = 368
    v8 := GenS1{F0: -9223372036854775808, F1: 0}
    v1 = float32(1e10)
    var v9 int8 = -1
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int16 = 0
        fmt.Println(v9)
        _ = fi10
        _ = v11
    }
    var v12 uint16 = (v7 - 0)
    var v13 uint16 = v7
    var v14 uint = 0
    fmt.Println(v7, v12)
    var v15 int8 = -119
    if ((true || (!false)) && (false || true)) {
        v16 := GenS1{F0: 42, F1: 0}
        _ = v16
    }
    switch -1 {
    case 561:
        var v17 uint = v14
        _ = v17
    }
    v1 = v1
    var v18 uint64 = 734
    v19 := []int{int(v12)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
