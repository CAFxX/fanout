package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 int16, p1 int16) uint32 {
    v5 := GenS1{F0: 36, F1: 0, F2: -1}
    var v6 uint32 = 263
    if true {
        fmt.Println(p0, v6, p1)
        fmt.Println(v6, v6, p1)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v6
}
func genFn1() uint64 {
    var v7 int32 = -938
    var v8 string = (((((("\n\t" + "foo bar") + "a") + (("hello" + "\n\t") + "")) + "hello") + ("" + "\n\t")) + ("" + "\n\t"))
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]int)
    var v3 uint16 = 42
    v4 := []int{4294967295}
    v2["hello"] = 128
    v3 = v3
    v9 := v2[(v1 + v1)]
    v3 = v3
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint32 = 276
        delete(v2, v1)
        _ = fi10
        _ = v11
    }
    v3 = 2
    var v12 int8 = -74
    v13 := GenS1{F0: v12, F1: v3, F2: -2147483648}
    switch (v1 + (((v1 + v1) + v1) + v1)) {
    case "":
        v3 -= (^(v3))
    }
    var v14 int8 = 100
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch (fi15 | v9) {
            case 128:
                v14 = (v12 / (((v14 & v12)) | 1))
            case 4294967295:
                var v17 float32 = (1.5 + 3.14159)
                _ = v17
            }
            _ = len(v4)
            v14 = (v14 ^ 1)
            _ = fi16
        }
        v18 := GenS1{F0: (-(v12)), F1: (^(v3)), F2: 920}
        var v19 int16 = -32768
        _ = fi15
        _ = v18
        _ = v19
    }
    switch (v9 >> ((255) & 63)) {
    case -932:
        switch v9 {
        case 128:
            var v20 uint8 = 2
            _ = v20
        case 127:
            v14 = -1
        case 32768:
            v21 := GenS2{F0: -1, F1: v3, F2: (v12 / ((-79) | 1))}
            _ = v21
        }
    case 425:
        v22 := GenS1{F0: v12, F1: (v3 | (v3 ^ v3)), F2: -2147483648}
        _ = v22
    case -9223372036854775808:
        var v23 uint8 = 128
        _ = v23
    default:
        switch 127 {
        case 128:
            switch v9 {
            case 2147483647:
                for fi24 := 0; fi24 < 18; fi24++ {
                    v9 = (v9 / (((-(fi24))) | 1))
                    v3 -= 239
                    v3 -= v3
                    _ = fi24
                }
            default:
                fmt.Println(v9, v9)
            }
        case -1:
            v3 = (v3 | (v3 * 2))
        case -611:
            var v25 int32 = 42
            _ = v25
        default:
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float64 = float64(-2.25)
                v3 -= v3
                var v28 uint8 = 97
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    }
    var v29 float64 = 0.0
    if (65535 != 1) {
        v30 := GenS1{F0: 1, F1: v3, F2: -1}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
