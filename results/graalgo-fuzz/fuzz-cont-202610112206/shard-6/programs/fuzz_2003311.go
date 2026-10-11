package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0() uint32 {
    var v6 int64 = -294
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v6 += v6
            var v9 uint32 = 1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 float32 = 1e10
    _ = v10
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 uint64 = 868
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    switch 128 {
    case 128:
        for fi11 := 0; fi11 < 6; fi11++ {
            v5 = (v5 / (v5 + v5))
            v12 := []int{-1, (fi11 * fi11), fi11, 4294967295}
            _ = fi11
            _ = v12
        }
    case -9223372036854775808:
        v13 := GenS1{F0: 0, F1: 255}
        _ = v13
    case 42:
        var v14 float32 = float32((0.0 / v5))
        _ = v14
    default:
        v5 -= v5
    }
    var v15 int32 = -88
    v16 := []int{}
    switch "foo bar" {
    case "hello":
        switch ("hello" + "hello") {
        case "hello":
            delete(v4, 792)
        case "\n\t":
            v17 := GenS2{F0: 14, F1: -634}
            _ = v17
        case "foo bar":
            var v18 int = -9223372036854775808
            _ = v18
        }
    }
    var v19 uint8 = 0
    var v20 int32 = v15
    v20 = (v20 & v15)
    v19 = 0
    if false {
        if len(v16) > 0 {
            v16[0] = 127
        }
        fmt.Println(v3, v20, v3)
    }
    v15 = (^(v15))
    var v21 uint64 = (v3 - v3)
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(v21, v20)
        var v23 uint16 = 838
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
