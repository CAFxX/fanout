package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }

func main() {
    var v1 int32 = 42
    var v2 bool = false
    var v3 bool = (!v2)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 13; fi6++ {
        v2 = (16 >= 42)
        for fi7 := 0; fi7 < 3; fi7++ {
            v1 -= v1
            fmt.Println(v1, fi6)
            _ = fi7
        }
        _ = fi6
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int8 = 0
        var v10 uint8 = 127
        fmt.Println(v2)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v2, F1: 135, F2: 100}
    switch 9223372036854775807 {
    case 639:
        v1 -= -1
    default:
        v12 := GenS1{F0: (3.14159 != 0.1), F1: 829, F2: 109}
        _ = v12
    }
    fmt.Println(v1, v3, v3)
    v13 := GenS1{F0: v3, F1: 1, F2: 0}
    switch -330 {
    case -667:
        var v14 int = 44
        _ = v14
    case -14:
        fmt.Println(v1, v1)
    case 430:
        var v15 uint = 599
        _ = v15
    }
    v16 := []int{}
    v3 = v3
    var v17 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
