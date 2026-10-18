package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }
type GenS2 struct { F0 int16; F1 string }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    var v3 int32 = 2147483647
    v4 := []uint32{214, 65535}
    v5 := []int{127, 0, 32767, 128}
    v6 := make(map[uint32]uint32)
    fmt.Println(v3, v3)
    var v7 uint32 = 65535
    fmt.Println(v7, v1)
    if ((false || (!false)) || false) {
        v8 := GenS1{F0: 42, F1: (v3 + (-(v3))), F2: 255}
        _ = v8
    } else {
        var v9 uint8 = 1
        var v10 uint8 = v9
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 663, F1: "foo bar"}
    v12 := GenS1{F0: 915, F1: -1, F2: 0}
    delete(v6, (v7 / ((v7) | 1)))
    var v13 int16 = -687
    var v14 uint = 4294967295
    v15 := GenS1{F0: -481, F1: v3, F2: 1}
    v16 := GenS1{F0: -9223372036854775808, F1: v3, F2: uint16(v14)}
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(v13, v14, v7)
        _ = fi17
    }
    var v18 int64 = 159
    for fi19 := 0; fi19 < 16; fi19++ {
        _ = len(v6)
        fmt.Println(v14, v14, v13)
        _ = fi19
    }
    v1 = 32
    v13 += v13
    for fi20 := 0; fi20 < 11; fi20++ {
        switch (fi20 >> ((v14) & 63)) {
        case 32767:
            if true {
                var v21 bool = (1 == uint64(v14))
                v13 = 1000
                var v22 float64 = 0.0
                _ = v21
                _ = v22
            }
        case -9223372036854775808:
            var v23 uint8 = 128
            _ = v23
        case 168:
            switch (fi20 % ((1) | 1)) {
            case 61:
                var v24 int8 = 42
                _ = v24
            case 9223372036854775807:
                fmt.Println(v3, v18, v14)
            default:
                if ((3.14159 < float32(1e10)) || (false && true)) {
                    if false {
                        var v25 bool = (!((!true) && false))
                        _ = v25
                    }
                    switch -686 {
                    case 464:
                        for fi26 := 0; fi26 < 19; fi26++ {
                            _ = len(v2)
                            var v27 uint16 = 32767
                            v4 = append(v4, (v7 | (v7 | 4294967295)))
                            var v28 uint32 = v7
                            var v29 int32 = (v3 & (v3 ^ (v3 * v3)))
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    default:
                        _ = len(v6)
                    }
                }
            }
        }
        _ = len(v5)
        for fi30 := 0; fi30 < 19; fi30++ {
            v13 = v13
            v14 = 65535
            _ = fi30
        }
        _ = fi20
    }
    v31 := GenS2{F0: v13, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
