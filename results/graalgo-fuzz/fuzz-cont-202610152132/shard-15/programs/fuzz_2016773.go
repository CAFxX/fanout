package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8, p1 int8) int16 {
    var v4 bool = (!true)
    var v5 float32 = 0.0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}
func genFn1() int16 {
    var v6 float32 = 1e10
    var v7 uint = 1
    v8 := GenS1{F0: 2, F1: 255}
    _ = v6
    _ = v7
    _ = v8
    return -671
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 725
    var v3 uint32 = (^(v2))
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, v3, v2)
        var v10 int64 = int64(v3)
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 4294967295
    v2 += (v2 ^ v2)
    fmt.Println(v3)
    var v12 uint8 = 127
    v3 = (v3 - v3)
    v13 := GenS1{F0: uint16(v11), F1: 354}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint8 = (^(v12))
        _ = fi14
        _ = v15
    }
    var v16 bool = false
    v1 = append(v1, (v3 >> ((255) & 63)))
    var v17 uint32 = (v3 ^ 65535)
    var v18 uint32 = (v3 << ((255) & 63))
    var v19 bool = v16
    switch -9223372036854775808 {
    case -1:
        if (!((!false) || v19)) {
            v1 = append(v1, 2147483647)
            var v20 bool = true
            _ = v20
        }
    case 420:
        var v21 float64 = (1.5 + 1.5)
        _ = v21
    case -9223372036854775808:
        for fi22 := 0; fi22 < 4; fi22++ {
            fmt.Println(v18, v16)
            var v23 float32 = -2.25
            if len(v1) > 0 {
                v1[0] = (v18 >> ((745) & 63))
            }
            _ = fi22
            _ = v23
        }
    }
    switch 2147483647 {
    case 747:
        fmt.Println(v18, v19, v19)
    case -878:
        if v19 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 128:
        v2 -= (v2 % ((v18) | 1))
    default:
        var v24 float32 = 1e10
        _ = v24
    }
    var v25 int64 = 320
    v2 = 4294967295
    v16 = (-2.25 > -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
