package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }

func main() {
    var v1 float32 = 3.14159
    v2 := []uint32{}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 65535
    var v5 int8 = -128
    if (1e10 != float64(v1)) {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint32 = 1
            var v8 uint = 0
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(v3, v5)
            var v10 int = fi9
            var v11 bool = (42 == 53)
            _ = fi9
            _ = v10
            _ = v11
        }
        v5 = 0
    } else {
        if false {
            fmt.Println(v4, v1)
            var v12 uint8 = 2
            _ = v12
        }
        var v13 int8 = v5
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi14
    }
    var v15 string = ("" + "foo bar")
    v16 := GenS1{F0: 455, F1: 756, F2: 65535}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint16 = v4
        if true {
            v19 := GenS2{F0: ((fi17 * (fi17 - (fi17 - fi17))) == (fi17 * fi17)), F1: uint(v4), F2: v5}
            var v20 float32 = 0.0
            _ = v19
            _ = v20
        } else {
            switch "foo bar" {
            case "":
                if false {
                    v2 = append(v2, 712)
                    v18 = 1
                    if (true || false) {
                        v1 = 3.14159
                    }
                }
            case "hello":
                fmt.Println(v4)
            }
            v21 := GenS2{F0: ((0.1 + 1.5) < float64((-2.25 - -2.25))), F1: 965, F2: (v5 / (((v5 >> ((65535) & 63))) | 1))}
            _ = v21
        }
        v18 += (v4 / ((uint16(v3)) | 1))
        _ = fi17
        _ = v18
    }
    if (-9223372036854775808 >= -297) {
        v1 += float32((0.0 / 3.14159))
        if len(v2) > 0 {
            v2[0] = uint32(v4)
        }
    }
    switch v15 {
    case "\n\t":
        if false {
            v22 := GenS2{F0: false, F1: 501, F2: v5}
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS2{F0: (1 != 0), F1: uint(v3), F2: int8(fi23)}
                v2 = append(v2, 250)
                v5 = (-(v5))
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            var v25 int64 = -9223372036854775808
            _ = len(v2)
            var v26 float64 = (float64(v1) - 1e10)
            _ = v25
            _ = v26
        }
    case "":
        v5 += v5
    case "a":
        fmt.Println(v1, v4)
    default:
        v15 = "hello"
    }
    v27 := GenS1{F0: uint16(v3), F1: 1, F2: 18446744073709551615}
    v15 = "foo bar"
    v4 = v4
    if false {
        v28 := GenS2{F0: false, F1: 42, F2: (v5 >> ((204) & 63))}
        var v29 uint32 = 2
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 uint16 = v4
        _ = fi30
        _ = v31
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
