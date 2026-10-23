package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }

func main() {
    v1 := []string{"", (("foo bar" + "\n\t") + "hello")}
    var v2 uint8 = 241
    v3 := []uint32{388, 4294967295, 364}
    v2 = v2
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int32 = 42
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        if (1 <= 2147483647) {
            v2 += 221
            v7 := GenS1{F0: -2147483648, F1: 121, F2: 1}
            var v8 uint32 = 2147483647
            _ = v7
            _ = v8
        }
        if len(v3) > 0 {
            v3[0] = 65535
        }
        _ = fi6
    }
    if (!true) {
        v2 -= v2
        for fi9 := 0; fi9 < 13; fi9++ {
            v2 = (v2 / ((96) | 1))
            var v10 int8 = -19
            switch "hello" {
            case "":
                fmt.Println(v10, v10, v2)
            case "a":
                var v11 int64 = 297
                _ = v11
            case "\n\t":
                for fi12 := 0; fi12 < 19; fi12++ {
                    var v13 float32 = 1.5
                    v10 = v10
                    _ = fi12
                    _ = v13
                }
            default:
                var v14 float64 = float64(((-2.25 + ((3.14159 - 0.0) * ((0.1 * float32((3.14159 * 0.0))) * 1e10))) + -2.25))
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        _ = len(v1)
    } else {
        var v15 uint32 = 389
        _ = v15
    }
    v2 = 1
    v2 = (v2 & v2)
    v16 := GenS2{F0: 127}
    switch 42 {
    case 128:
        var v17 int64 = -126
        _ = v17
    case 9223372036854775807:
        v18 := GenS2{F0: (v2 | (^(v2)))}
        _ = v18
    case 105:
        var v19 uint8 = 34
        _ = v19
    default:
        switch 0 {
        case 32768:
            v2 += 0
        case 127:
            if ((false && true) || true) {
                var v20 int16 = 32767
                v20 += (-(v20))
                _ = v20
            }
        }
    }
    v21 := GenS3{F0: 530, F1: (v2 + v2), F2: 207}
    v22 := GenS1{F0: -173, F1: 18446744073709551615, F2: 255}
    v2 = v2
    v2 = v2
    if len(v3) > 0 {
        v3[0] = 1
    }
    v2 = (^(v2))
    var v23 uint16 = 2
    v24 := GenS1{F0: 0, F1: 4294967295, F2: 0}
    for fi25 := 0; fi25 < 9; fi25++ {
        if len(v3) > 0 {
            v3[0] = 65535
        }
        _ = fi25
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
