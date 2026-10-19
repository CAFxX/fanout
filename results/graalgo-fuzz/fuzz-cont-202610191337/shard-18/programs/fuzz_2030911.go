package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0() uint16 {
    if true {
        var v7 uint64 = 699
        _ = v7
    } else {
        var v8 int = -657
        fmt.Println(v8)
        v9 := []int{128, 42, v8, 2147483647}
        _ = v8
        _ = v9
    }
    var v10 uint8 = 2
    _ = v10
    return 1
}

func main() {
    var v1 uint64 = 163
    v2 := []int{2147483647, -156, -401}
    var v3 int16 = 32767
    var v4 uint32 = 1
    var v5 uint16 = 32767
    v6 := make(map[string]int)
    if ((v4 - v4) >= 2) {
        var v11 uint = 514
        var v12 uint16 = 919
        _ = v11
        _ = v12
    }
    if false {
        v4 -= v4
        var v13 int32 = 42
        var v14 float32 = 1.5
        _ = v13
        _ = v14
    } else {
        var v15 float64 = (0.1 * 1e10)
        _ = v15
    }
    v3 = v3
    _ = len(v6)
    switch -28 {
    case -2147483648:
        _ = len(v2)
    case 2147483647:
        v16 := GenS1{F0: (!false), F1: v1}
        _ = v16
    case -9223372036854775808:
        v3 -= int16(v5)
    }
    v17 := GenS1{F0: false, F1: v1}
    for fi18 := 0; fi18 < 16; fi18++ {
        v1 += (^(v1))
        for fi19 := 0; fi19 < 13; fi19++ {
            delete(v6, ("a" + ("\n\t" + "\n\t")))
            _ = fi19
        }
        _ = len(v2)
        _ = fi18
    }
    v20 := GenS1{F0: (22 >= 74), F1: uint64(v5)}
    switch 207 {
    case 42:
        v4 -= (v4 / (((v4 | (v4 >> ((16) & 63)))) | 1))
    case 9223372036854775807:
        var v21 bool = false
        _ = v21
    case 4294967295:
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 uint8 = 127
            _ = fi22
            _ = v23
        }
    default:
        fmt.Println(v4)
    }
    v4 -= (^(v4))
    v24 := GenS2{F0: 0, F1: -732}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
