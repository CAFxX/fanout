package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
func genFn0() uint8 {
    v4 := GenS1{F0: 649, F1: 2147483647, F2: (!((0 > 1) && false))}
    var v5 uint16 = 843
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 bool = (0 == 42)
    var v2 uint32 = 2147483648
    v3 := []uint32{533, v2}
    switch -2147483648 {
    case 1:
        v6 := GenS1{F0: 42, F1: (v2 * 355), F2: v1}
        _ = v6
    default:
        _ = len(v3)
    }
    _ = len(v3)
    v2 = (v2 - (v2 / ((2) | 1)))
    var v7 string = ""
    var v8 int = -1
    switch ("a" + "") {
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS2{F0: v1}
            var v11 float64 = -2.25
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        var v12 int16 = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS3{F0: 18446744073709551615, F1: 1, F2: 0}
        fmt.Println(fi13)
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int8 = -45
            var v17 int32 = 2147483647
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    v2 = 2147483647
    for fi18 := 0; fi18 < 4; fi18++ {
        if false {
            v19 := GenS4{F0: 1}
            _ = v19
        }
        var v20 string = v7
        _ = fi18
        _ = v20
    }
    fmt.Println(v7, v2, v1)
    v1 = (!(0 == 9223372036854775807))
    fmt.Println(v7, v7, v2)
    v21 := GenS5{F0: v8, F1: 801}
    v22 := GenS1{F0: 2147483648, F1: v2, F2: (103 == genFn0())}
    if (!false) {
        var v23 int = 220
        v24 := GenS6{F0: -11}
        _ = v23
        _ = v24
    }
    var v25 float32 = (1.5 - 1e10)
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS7{F0: (v2 | v2)}
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
