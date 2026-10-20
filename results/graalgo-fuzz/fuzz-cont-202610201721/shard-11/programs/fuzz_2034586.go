package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    v3 := []uint32{5}
    v4 := []int{-572, 875}
    var v5 int = 42
    v6 := []uint32{42, 65535, 76}
    if false {
        if ((((("hello" + "a") + ("" + (("foo bar" + "\n\t") + ("foo bar" + "hello")))) + ("" + "a")) + "a") == "") {
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 int64 = -910
                switch "hello" {
                case "a":
                    v9 := GenS1{F0: 127, F1: (1 == 722), F2: -60}
                    _ = v9
                case "\n\t":
                    var v10 uint16 = 32767
                    _ = v10
                default:
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
                v11 := GenS2{F0: v8, F1: 1}
                _ = fi7
                _ = v11
                _ = v8
            }
        } else {
            v5 = v5
        }
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := v2[830]
        var v14 uint16 = 227
        fmt.Println(v14, v13, v5)
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        _ = len(v2)
        var v16 bool = (true || ((v1 && (-128 <= -71)) && (3.14159 >= 1e10)))
        _ = fi15
        _ = v16
    }
    v1 = v1
    switch 127 {
    case -9223372036854775808:
        for fi17 := 0; fi17 < 5; fi17++ {
            v5 = -245
            _ = fi17
        }
    case -13:
        if v1 {
            fmt.Println(v5)
            v5 = -9223372036854775808
            v18 := GenS1{F0: 1, F1: v1, F2: 100}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int16 = 1000
                v1 = true
                _ = fi19
                _ = v20
            }
        }
    case -488:
        var v21 float32 = 1e10
        _ = v21
    default:
        var v22 uint8 = 1
        _ = v22
    }
    v23 := GenS3{F0: 372}
    if (v1 && (!(true && (true || (0 == 902))))) {
        v5 = (v5 - (v5 / (((v5 & v5)) | 1)))
        var v24 int = (-(v5))
        v5 = (v5 << ((2) & 63))
        _ = v24
    }
    v1 = (!((v5 << ((211) & 63)) < v5))
    v1 = v1
    var v25 bool = (!(0 > 59))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
