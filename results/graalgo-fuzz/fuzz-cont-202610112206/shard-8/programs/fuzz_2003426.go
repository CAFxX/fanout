package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 bool) string {
    var v5 int8 = -128
    fmt.Println(p0, p0)
    var v6 int64 = -25
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = (true || (float64(1.5) < -2.25))
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]bool)
    v3 := []uint32{}
    var v4 uint = 4294967295
    v8 := GenS1{F0: 542, F1: (v1 % ((124) | 1))}
    v4 = (v4 ^ 18446744073709551615)
    v4 -= 2
    v9 := GenS1{F0: -118, F1: (v1 / ((0) | 1))}
    switch 32767 {
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 2147483647:
        v10 := GenS1{F0: -420, F1: v1}
        _ = v10
    case -163:
        fmt.Println(v1, v1)
    }
    v1 = (v1 | (v1 * 42))
    if ((-2.25 >= 3.14159) && (!false)) {
        fmt.Println(v1, v1)
        var v11 float32 = 0.1
        _ = v11
    } else {
        var v12 string = ("" + ("foo bar" + ""))
        if len(v3) > 0 {
            v3[0] = v1
        }
        _ = v12
    }
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 13; fi13++ {
        v1 += v1
        if false {
            var v14 uint8 = 128
            _ = len(v3)
            _ = v14
        } else {
            v4 = v4
            var v15 float64 = ((0.0 + (1.5 * (3.14159 * 1e10))) - ((1.5 + -2.25) - (float64(float32(0.1)) * 1.5)))
            if (2 >= 32767) {
                var v16 uint64 = 42
                _ = v16
            }
            _ = v15
        }
        _ = fi13
    }
    v1 += 650
    if false {
        switch "foo bar" {
        case "":
            v1 = v1
        default:
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 int = -1
                _ = fi17
                _ = v18
            }
        }
    } else {
        var v19 uint32 = 2147483647
        for fi20 := 0; fi20 < 4; fi20++ {
            if (1 > 935) {
                v1 = v19
            }
            v19 += v1
            _ = fi20
        }
        _ = v19
    }
    switch ("a" + ("" + "\n\t")) {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v1, v1)
        case "a":
            v1 = 2147483647
        default:
            var v21 uint = (v4 >> ((300) & 63))
            _ = v21
        }
    case "a":
        fmt.Println(v1, v1)
    }
    v22 := GenS1{F0: 42, F1: v1}
    _ = len(v3)
    v4 = (v4 >> (((v4 << ((v4) & 63))) & 63))
    v23 := GenS1{F0: 762, F1: 1}
    if true {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint32 = (v1 / ((uint32(v4)) | 1))
            _ = fi24
            _ = v25
        }
        delete(v2, 9223372036854775807)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
