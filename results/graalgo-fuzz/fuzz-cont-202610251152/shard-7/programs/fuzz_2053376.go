package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }

func main() {
    var v1 int16 = -379
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    var v5 uint32 = 985
    v6 := make(map[string]int)
    v4 = v4
    v7 := GenS1{F0: 42, F1: v2}
    v8 := GenS1{F0: 1, F1: v2}
    for fi9 := 0; fi9 < 12; fi9++ {
        v3[(v5 >> ((0) & 63))] = (v5 * v5)
        var v10 int32 = 111
        _ = fi9
        _ = v10
    }
    v11 := []int{1}
    if len(v11) > 0 {
        v11[0] = 316
    }
    if ((v1 - 206) == (v1 << ((0) & 63))) {
        var v12 int16 = -32768
        var v13 uint = 26
        fmt.Println(v5, v1, v1)
        _ = v12
        _ = v13
    } else {
        fmt.Println(v1, v2)
    }
    _ = len(v11)
    switch -690 {
    case 128:
        v5 = v5
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS1{F0: 542, F1: v2}
        var v16 int32 = -2147483648
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 float64 = 1.5
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if false {
        var v18 uint16 = 42
        var v19 int = 32767
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS2{F0: 2147483647, F1: 2, F2: (^(v4))}
            _ = len(v3)
            v22 := GenS1{F0: 42, F1: (2 >= v4)}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v18
        _ = v19
    }
    _ = len(v11)
    fmt.Println(v4, v4, v4)
    v11 = append(v11, 554)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
