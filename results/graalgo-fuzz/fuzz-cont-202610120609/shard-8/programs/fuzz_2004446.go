package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int16; F1 bool }
func genFn0() uint {
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 -= (v7 * v7)
        v7 -= v7
        _ = fi8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint16 = 115
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 255, F2: 42}
    _ = v11
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 609
    var v2 uint = 548
    var v3 int8 = int8(v1)
    var v4 uint = (v2 & v2)
    v5 := []uint32{4294967295, 375, 42, 1}
    var v6 int16 = 1
    var v12 uint16 = 65535
    v1 -= (v1 & 888)
    v1 = (v12 | (v1 | (v12 << ((974) & 63))))
    if true {
        v13 := GenS1{F0: 0, F1: 2, F2: 2147483647}
        _ = v13
    }
    v1 = v12
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v2, v2, v1)
        v15 := GenS2{F0: 1, F1: 9223372036854775807, F2: 128}
        _ = fi14
        _ = v15
    }
    v3 -= 96
    if false {
        v4 = (v4 + v2)
    }
    v4 -= v4
    if len(v5) > 0 {
        v5[0] = 214
    }
    v4 = v4
    var v16 float64 = (1.5 * (0.0 + 1.5))
    var v17 int32 = -2147483648
    var v18 uint8 = 191
    switch "hello" {
    case "foo bar":
        v12 -= v12
    case "":
        for fi19 := 0; fi19 < 8; fi19++ {
            if (!(!(!true))) {
                var v20 uint16 = (v12 - v1)
                v21 := GenS3{F0: 1, F1: false}
                var v22 int64 = 9223372036854775807
                _ = v20
                _ = v21
                _ = v22
            } else {
                v23 := GenS1{F0: 4294967295, F1: (v18 << (((v4 >> ((v2) & 63))) & 63)), F2: 2147483648}
                fmt.Println(v6, v3)
                v5 = append(v5, 448)
                _ = v23
            }
            _ = fi19
        }
    case "a":
        v17 += (v17 | (-(v17)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
