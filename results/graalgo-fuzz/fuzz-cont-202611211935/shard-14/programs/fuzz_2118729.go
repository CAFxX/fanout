package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 65535
    var v3 int8 = 0
    v4 := []string{}
    v2 = 42
    if (42 < 65535) {
        for fi5 := 0; fi5 < 8; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                v3 = int8(fi6)
                fmt.Println(v3, v3, v3)
                v3 += (-(v3))
                _ = fi6
            }
            _ = fi5
        }
        _ = len(v4)
    }
    v7 := GenS1{F0: 2147483647}
    switch 127 {
    case 32767:
        var v8 uint16 = (^(v2))
        _ = v8
    default:
        delete(v1, 2147483647)
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v3 = (v3 & 18)
        v10 := GenS1{F0: uint32(v2)}
        _ = fi9
        _ = v10
    }
    var v11 int8 = (v3 & -1)
    fmt.Println(v11, v3)
    var v12 float64 = 3.14159
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int8 = v11
            _ = fi14
            _ = v15
        }
        var v16 uint = 4294967295
        if false {
            fmt.Println(fi13, v3, v2)
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                switch -9223372036854775808 {
                case -1:
                    if false {
                        v18 := GenS1{F0: 2147483648}
                        _ = v18
                    }
                case 127:
                    var v19 int64 = 9223372036854775807
                    _ = v19
                default:
                    v12 += v12
                }
                _ = fi17
            }
            v2 += (^(v2))
            fmt.Println(v16)
        }
        _ = fi13
        _ = v16
    }
    switch -9223372036854775808 {
    case 1:
        v11 = int8(v2)
    case 127:
        var v20 int = -572
        _ = v20
    case -9223372036854775808:
        switch 4294967295 {
        case -318:
            v21 := GenS1{F0: 852}
            _ = v21
        }
    }
    var v22 float32 = 0.0
    fmt.Println(v2)
    var v23 int = -2147483648
    v24 := GenS2{F0: 818, F1: -2147483648, F2: (v11 * v11)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
