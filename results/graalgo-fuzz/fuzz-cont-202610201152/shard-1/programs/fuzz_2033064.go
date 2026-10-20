package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 bool; F1 int16; F2 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{2, 649, 496, 0}
    var v2 uint8 = 255
    var v3 int64 = 42
    var v4 uint64 = 1
    v2 = (v2 % ((128) | 1))
    var v5 int16 = 32767
    var v6 int32 = -2147483648
    v7 := GenS1{F0: -1, F1: 4294967295, F2: 65535}
    var v8 int64 = 118
    v1 = append(v1, 958)
    v9 := GenS2{F0: v2, F1: ""}
    v2 = 38
    v10 := GenS1{F0: 0, F1: 1000, F2: 935}
    if false {
        v3 = 792
        var v11 int8 = 82
        var v12 float32 = 0.0
        _ = v11
        _ = v12
    }
    switch "\n\t" {
    case "hello":
        for fi13 := 0; fi13 < 5; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 += (^(v2))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi13
        }
    case "\n\t":
        fmt.Println(v6)
    default:
        v14 := GenS3{F0: false, F1: v5, F2: 65535}
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int16 = (v5 << ((18446744073709551615) & 63))
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 float64 = (-2.25 * -2.25)
            var v19 uint32 = 4294967295
            v20 := GenS4{F0: 127, F1: (v2 % (((^(v2))) | 1)), F2: 2147483647}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    }
    switch 4294967295 {
    case -325:
        v8 = v8
    case 831:
        v1 = append(v1, 0)
    case -9223372036854775808:
        switch "a" {
        case "foo bar":
            var v21 uint16 = 32767
            _ = v21
        case "hello":
            v22 := GenS3{F0: (int8(v2) == int8(v6)), F1: 1, F2: 4294967295}
            _ = v22
        case "":
            v8 += (v3 / ((v8) | 1))
        }
    default:
        v1 = append(v1, 1)
    }
    v2 = v2
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 float32 = float32(-2.25)
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
