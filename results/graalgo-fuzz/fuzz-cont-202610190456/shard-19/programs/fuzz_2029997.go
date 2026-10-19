package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8; F1 bool; F2 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
func genFn0(p0 bool) bool {
    if p0 {
        var v7 float64 = 0.1
        switch -1 {
        case -526:
            v8 := GenS1{F0: 125, F1: 127}
            _ = v8
        case -751:
            var v9 string = "\n\t"
            _ = v9
        default:
            fmt.Println(v7, p0)
        }
        _ = v7
    } else {
        v10 := GenS1{F0: 709, F1: -124}
        _ = v10
    }
    var v11 int8 = 84
    for fi12 := 0; fi12 < 5; fi12++ {
        v11 = v11
        var v13 float32 = 0.1
        var v14 uint = 2
        _ = fi12
        _ = v13
        _ = v14
    }
    if p0 {
        v15 := GenS2{F0: v11, F1: (1 < 9223372036854775807), F2: 1}
        _ = v15
    } else {
        var v16 uint16 = 255
        v16 -= (v16 / ((65535) | 1))
        _ = v16
    }
    if ((0.0 / 3.14159) >= -2.25) {
        v17 := GenS3{F0: 42, F1: 385, F2: 783}
        _ = v17
    }
    _ = p0
    _ = v11
    return p0
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + "")}
    v2 := []string{}
    v3 := []uint32{4294967295, 2147483647, 1}
    v4 := []string{"foo bar", (("\n\t" + (("" + ("\n\t" + "hello")) + ("foo bar" + ("a" + "hello")))) + "a"), ("" + "hello"), ""}
    v5 := []uint32{132, 2147483648}
    v6 := make(map[uint32]uint32)
    if (-351 <= 42) {
        var v18 uint8 = 127
        switch -801 {
        case 0:
            fmt.Println(v18)
        case 1:
            v18 += (^(v18))
        case 527:
            if (0 < -450) {
                switch "\n\t" {
                case "hello":
                    v18 -= v18
                }
                var v19 uint8 = v18
                switch -1 {
                case -338:
                    _ = len(v2)
                case 127:
                    var v20 int32 = -1
                    _ = v20
                }
                _ = v19
            }
        }
        _ = v18
    } else {
        var v21 int16 = -1
        _ = v21
    }
    if (419 != 0) {
        var v22 bool = ((0.0 - 3.14159) < ((1.5 + ((-2.25 * 0.0) * 1e10)) - (((1e10 + ((0.0 + float64((0.0 / 3.14159))) + float64(1e10))) + 1.5) - ((1.5 + 3.14159) + 0.1))))
        _ = v22
    } else {
        var v23 int64 = -1
        _ = v23
    }
    delete(v6, 863)
    var v24 uint64 = 4294967295
    v24 -= 2
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v24, v24)
        _ = fi25
    }
    v26 := GenS3{F0: -2147483648, F1: v24, F2: 32767}
    v24 -= 228
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
