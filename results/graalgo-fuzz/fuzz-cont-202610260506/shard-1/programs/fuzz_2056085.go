package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []int{}
    var v2 int64 = 9223372036854775807
    var v3 float64 = 1e10
    var v4 int8 = 127
    if false {
        v5 := GenS1{F0: 65535, F1: 42, F2: v2}
        _ = v5
    } else {
        var v6 uint8 = 67
        v7 := GenS2{F0: (v2 / ((203) | 1)), F1: "\n\t"}
        _ = v6
        _ = v7
    }
    var v8 uint32 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 += -42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    switch 225 {
    case 0:
        v9 := GenS3{F0: 18446744073709551615, F1: (true || (true || (!false))), F2: v8}
        _ = v9
    default:
        switch -562 {
        case 0:
            _ = len(v1)
        case 127:
            v3 = 0.0
        }
    }
    switch -1 {
    case -1:
        for fi10 := 0; fi10 < 2; fi10++ {
            v8 -= 2147483648
            v8 = 538
            v3 -= float64(1.5)
            _ = fi10
        }
    case 2147483647:
        v4 = int8(v8)
    case -312:
        var v11 int16 = 0
        _ = v11
    default:
        switch "foo bar" {
        case "":
            v4 = v4
        }
    }
    var v12 int16 = 325
    _ = len(v1)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS4{F0: -599, F1: 0}
            v8 -= 2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi14
            _ = v15
        }
        v2 -= (^(v2))
        var v16 uint16 = 32767
        _ = fi13
        _ = v16
    }
    var v17 uint8 = 42
    v18 := GenS5{F0: uint16(v8), F1: 2147483647}
    var v19 uint8 = 42
    for fi20 := 0; fi20 < 10; fi20++ {
        v19 -= v19
        v8 -= v8
        var v21 float32 = 3.14159
        _ = fi20
        _ = v21
    }
    fmt.Println(v12, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
