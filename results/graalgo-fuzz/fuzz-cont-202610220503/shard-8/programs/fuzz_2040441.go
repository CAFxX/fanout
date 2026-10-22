package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
type GenS5 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 uint = 711
    var v3 uint64 = 195
    var v4 uint8 = 0
    var v5 int8 = 127
    v1 = append(v1, 65535)
    v5 = v5
    v3 = 4294967295
    v4 = v4
    if (!true) {
        for fi6 := 0; fi6 < 17; fi6++ {
            v3 += (^(v3))
            switch int(v3) {
            case 127:
                var v7 uint16 = 255
                _ = v7
            case 812:
                v8 := GenS1{F0: v5, F1: (v4 & 20)}
                _ = v8
            }
            _ = fi6
        }
        v9 := GenS2{F0: -350, F1: (("hello" + ("foo bar" + "")) + "\n\t")}
        _ = v9
    } else {
        var v10 uint32 = 0
        v11 := GenS1{F0: -88, F1: v4}
        var v12 float32 = 0.1
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println(v5, v2)
    for fi13 := 0; fi13 < 5; fi13++ {
        if true {
            var v14 float64 = (1e10 + 1.5)
            _ = v14
        } else {
            fmt.Println(v5, v5, v5)
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 float32 = float32(((0.0 * 3.14159) / float64(-2.25)))
            var v17 int16 = 25
            switch fi13 {
            case 128:
                var v18 int8 = v5
                _ = v18
            case 636:
                if ((32767 >= 0) && false) {
                    var v19 uint = (^(v2))
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    v16 += 1e10
                    _ = v19
                } else {
                    v20 := GenS3{F0: -1}
                    fmt.Println(fi15, v4, fi13)
                    switch fi13 {
                    case 833:
                        var v21 float64 = 0.1
                        _ = v21
                    }
                    fmt.Println(v16, v4)
                    _ = len(v1)
                    var v22 int32 = 0
                    _ = v20
                    _ = v22
                }
            case 510:
                fmt.Println(fi13, v4, v16)
            default:
                var v23 uint64 = v3
                _ = v23
            }
            v16 = v16
            fmt.Println(v5, v5)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v24 := GenS4{F0: v4, F1: 32767, F2: uint64(v4)}
    v25 := GenS4{F0: 44, F1: 32768, F2: (v3 % (((v3 % (((v3 & (v3 & uint64(v2)))) | 1))) | 1))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v4, v4)
    v26 := GenS5{F0: -2, F1: 1, F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
