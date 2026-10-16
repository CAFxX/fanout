package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 32767
    var v3 int32 = 1
    var v4 bool = ((v2 <= (v2 - v2)) || (false || false))
    var v5 uint16 = (v2 * (v2 + (v2 ^ v2)))
    v6 := []int{42, 299, -74}
    var v7 int16 = -1
    if v4 {
        for fi8 := 0; fi8 < 5; fi8++ {
            v5 -= v5
            var v9 uint16 = (v5 - v2)
            _ = fi8
            _ = v9
        }
        if true {
            v2 = v2
            var v10 uint32 = 1
            _ = v10
        } else {
            var v11 float64 = 0.0
            if v4 {
                v5 += (^(v2))
            } else {
                v3 = v3
            }
            _ = v11
        }
    } else {
        v12 := GenS1{F0: (v3 % (((v3 + v3)) | 1)), F1: uint16(v1)}
        var v13 uint8 = 243
        _ = v12
        _ = v13
    }
    switch "hello" {
    case "hello":
        var v14 int = -9223372036854775808
        _ = v14
    default:
        if len(v6) > 0 {
            v6[0] = 32768
        }
    }
    fmt.Println(v5, v5)
    v4 = (0 < 0)
    for fi15 := 0; fi15 < 6; fi15++ {
        v2 += (v2 & 32767)
        v3 -= (v3 + (v3 & (v3 % ((v3) | 1))))
        v6 = append(v6, int(v1))
        _ = fi15
    }
    if (0.1 <= float32(0.1)) {
        v1 = (v1 + (^(v1)))
    } else {
        for fi16 := 0; fi16 < 1; fi16++ {
            v17 := GenS1{F0: int32(v2), F1: 192}
            _ = fi16
            _ = v17
        }
    }
    var v18 float32 = -2.25
    if v4 {
        v2 += v2
        switch -9223372036854775808 {
        case 4294967295:
            v19 := GenS2{F0: (v3 & v3), F1: v2}
            _ = v19
        default:
            fmt.Println(v3, v18, v1)
        }
        var v20 int64 = 42
        _ = v20
    }
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
