package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint32 }
func genFn0() int32 {
    var v4 int8 = -28
    v4 -= (v4 | v4)
    var v5 string = ""
    _ = v4
    _ = v5
    return 659
}

func main() {
    var v1 uint16 = 0
    v2 := []uint32{927}
    var v3 int8 = -1
    var v6 int16 = 32767
    v7 := GenS1{F0: 32767}
    v2 = append(v2, 2)
    var v8 float32 = 3.14159
    var v9 int8 = (v3 / ((0) | 1))
    if (781 != -576) {
        _ = len(v2)
        var v10 int64 = 9223372036854775807
        _ = v10
    } else {
        switch 127 {
        case -9223372036854775808:
            _ = len(v2)
        case 127:
            var v11 int32 = -554
            _ = v11
        }
        v12 := GenS1{F0: (^(v6))}
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int = (-(fi13))
        v14 = v14
        _ = fi13
        _ = v14
    }
    var v15 int32 = -857
    var v16 uint = 255
    v17 := GenS1{F0: v6}
    if true {
        switch 127 {
        case 0:
            v8 = (v8 - (v8 * v8))
        default:
            fmt.Println(v15)
        }
        fmt.Println(v3, v15, v9)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    switch 32767 {
    case 32767:
        switch 9223372036854775807 {
        case 128:
            if (v9 >= (v3 ^ int8(v6))) {
                _ = len(v2)
                var v18 bool = false
                _ = v18
            }
        case -226:
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v3)
            }
        case 32767:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        default:
            v9 = -128
        }
    case 9223372036854775807:
        v8 -= (v8 - v8)
    case -560:
        if (!((v16 * (v16 | 254)) != (v16 - 255))) {
            v19 := GenS2{F0: 148, F1: (v15 & (v15 * -2147483648)), F2: 756}
            fmt.Println(v15)
            for fi20 := 0; fi20 < 9; fi20++ {
                v3 -= (v9 >> ((2) & 63))
                v1 += uint16(v16)
                _ = fi20
            }
            _ = v19
        } else {
            v21 := GenS2{F0: 1, F1: -34, F2: 0}
            _ = v21
        }
    default:
        if false {
            var v22 float64 = float64(v8)
            v15 -= -1
            _ = v22
        } else {
            if true {
                fmt.Println(v9, v9)
                fmt.Println(v8)
            } else {
                fmt.Println(v16, v9)
                var v23 int32 = (v15 ^ (v15 - (v15 ^ genFn0())))
                if true {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v1)
                        var v25 int64 = 773
                        _ = fi24
                        _ = v25
                    }
                    v16 = (v16 & (v16 << (((v16 + (v16 ^ (v16 ^ (v16 / ((v16) | 1)))))) & 63)))
                }
                _ = v23
            }
            var v26 int = 42
            _ = v26
        }
    }
    v2 = append(v2, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
