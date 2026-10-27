package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 int }
func genFn0(p0 uint8) uint32 {
    var v5 uint64 = 42
    var v6 uint32 = uint32(v5)
    if true {
        if true {
            var v7 uint64 = 0
            _ = v7
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return (v6 | (v6 << ((65535) & 63)))
}

func main() {
    v1 := []uint32{431, 1, 469}
    var v2 uint8 = 82
    var v3 int16 = 42
    var v4 int32 = 351
    v4 = 0
    v8 := GenS1{F0: ((true || false) && false), F1: 4294967295}
    v3 = v3
    var v9 int8 = 100
    v10 := GenS1{F0: false, F1: 42}
    var v11 uint = 1
    if true {
        if (65535 < 282) {
            var v12 uint8 = (^(v2))
            _ = v12
        } else {
            if false {
                fmt.Println(v3)
                v13 := GenS1{F0: true, F1: uint64(v11)}
                _ = v13
            } else {
                v1 = append(v1, 1)
            }
            fmt.Println(v3)
            if ((2147483647 != genFn0(v2)) && true) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    v15 := GenS2{F0: (v3 >> ((v11) & 63))}
                    _ = fi14
                    _ = v15
                }
                v2 = 0
                var v16 int32 = v4
                _ = v16
            }
        }
    } else {
        v17 := GenS3{F0: (0 > 12), F1: int16(v11), F2: 128}
        _ = v17
    }
    v4 = (-(v4))
    for fi18 := 0; fi18 < 9; fi18++ {
        fmt.Println(v3)
        var v19 bool = ((v2 & (v2 << ((v11) & 63))) == 1)
        v3 = (v3 * (v3 - v3))
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: v3}
    var v21 uint8 = (v2 ^ v2)
    var v22 int32 = (v4 * 0)
    v1 = append(v1, genFn0((v21 | 205)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
