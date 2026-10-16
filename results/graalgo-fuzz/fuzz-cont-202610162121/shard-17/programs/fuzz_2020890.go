package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 int16 = 0
    var v5 bool = (false || false)
    if v5 {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 float32 = -2.25
            v8 := []int{fi6}
            v4 -= v4
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(v4, v5)
    }
    var v9 float32 = (float32(float64(0.1)) * 0.0)
    v10 := GenS1{F0: 96}
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{(fi11 / ((-2147483648) | 1)), fi11, -2147483648, fi11}
        _ = len(v12)
        switch fi11 {
        case -142:
            v5 = (v9 < 1.5)
        case 0:
            var v13 int8 = 127
            _ = v13
        case -9223372036854775808:
            v14 := GenS1{F0: 680}
            _ = v14
        default:
            v15 := GenS2{F0: 303, F1: 1}
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v4 = (v4 | (v4 - (v4 ^ v4)))
    var v16 uint = 1
    v5 = (!(!(722 < 67)))
    var v17 uint = v16
    switch 128 {
    case 42:
        var v18 uint16 = 133
        _ = v18
    case 32768:
        v16 = v17
    case 32767:
        var v19 int64 = 42
        _ = v19
    default:
        v20 := GenS1{F0: 1}
        _ = v20
    }
    var v21 int16 = v4
    var v22 float64 = (0.0 / 0.1)
    switch ("\n\t" + "foo bar") {
    case "a":
        v17 = (v16 * (v17 & v17))
    }
    var v23 int = 4294967295
    v24 := []int{v23, 32768, -984}
    v21 = v4
    if len(v24) > 0 {
        v24[0] = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
