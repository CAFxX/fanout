package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int64 = 42
    var v2 float64 = (3.14159 + 3.14159)
    v3 := []uint32{1}
    var v4 int16 = 616
    v5 := make(map[int]string)
    v6 := []int{}
    for fi7 := 0; fi7 < 14; fi7++ {
        if (-1 <= 42) {
            v8 := GenS1{F0: 1, F1: false, F2: (fi7 >> ((18446744073709551615) & 63))}
            var v9 uint64 = 484
            v4 = int16(v1)
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: 0, F1: false, F2: -506}
    v11 := GenS1{F0: 32767, F1: (true && false), F2: -739}
    fmt.Println(v2)
    v12 := GenS1{F0: 0, F1: (!((v1 * v1) > 232)), F2: -219}
    switch "a" {
    case "foo bar":
        var v13 uint32 = 4294967295
        _ = v13
    case "hello":
        if (!(!false)) {
            var v14 float64 = float64(1e10)
            switch 128 {
            case 683:
                var v15 int16 = (v4 << ((42) & 63))
                _ = v15
            }
            var v16 uint64 = 1
            _ = v14
            _ = v16
        } else {
            fmt.Println(v4)
            v4 += (v4 | v4)
        }
    case "\n\t":
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS2{F0: 0, F1: 42}
            _ = fi17
            _ = v18
        }
    default:
        var v19 int8 = -128
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        switch "a" {
        case "foo bar":
            v2 += ((1.5 + (((float64(3.14159) + v2) * (v2 * 0.1)) / v2)) * (v2 * 3.14159))
        case "\n\t":
            for fi21 := 0; fi21 < 11; fi21++ {
                if len(v3) > 0 {
                    v3[0] = 42
                }
                v22 := GenS1{F0: 575, F1: true, F2: 32767}
                _ = fi21
                _ = v22
            }
        default:
            var v23 uint32 = 375
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v4 -= v4
        v2 -= ((v2 + ((v2 / -2.25) * (-2.25 * 1e10))) + 3.14159)
        v1 -= (v1 % (((v1 & 9223372036854775807)) | 1))
        _ = fi24
    }
    v25 := GenS3{F0: 255}
    var v26 uint8 = 233
    _ = len(v5)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
