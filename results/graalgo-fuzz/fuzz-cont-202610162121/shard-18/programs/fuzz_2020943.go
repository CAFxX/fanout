package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
func genFn0() uint16 {
    var v4 float64 = 1e10
    var v5 int64 = 841
    v6 := GenS1{F0: true, F1: 32768, F2: -265}
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{1, 65535, 2147483648, 42}
    var v2 uint8 = 2
    var v3 int8 = -128
    v3 += (v3 ^ 13)
    if (false || true) {
        var v7 string = (("\n\t" + "hello") + "")
        _ = v7
    }
    v1 = append(v1, 1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS2{F0: 2147483648, F1: 752}
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "foo bar" {
        case "hello":
            for fi10 := 0; fi10 < 14; fi10++ {
                switch fi9 {
                case -900:
                    v11 := GenS2{F0: 923, F1: 1}
                    _ = v11
                case -357:
                    v1 = append(v1, 2147483647)
                case -1:
                    v12 := GenS3{F0: 32767}
                    _ = v12
                }
                var v13 int8 = v3
                _ = fi10
                _ = v13
            }
        default:
            v2 = (v2 * (^(v2)))
        }
        fmt.Println(v2)
        v2 = 197
        _ = fi9
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch (fi14 * fi14) {
        case 599:
            if false {
                var v15 uint = 18446744073709551615
                if true {
                    var v16 uint = v15
                    v17 := GenS1{F0: false, F1: 2147483647, F2: 1}
                    _ = v16
                    _ = v17
                } else {
                    var v18 int32 = -272
                    _ = v18
                }
                var v19 uint16 = 255
                _ = v15
                _ = v19
            }
        case 9223372036854775807:
            var v20 int64 = -9223372036854775808
            _ = v20
        }
        _ = fi14
    }
    v21 := GenS1{F0: true, F1: 32768, F2: -720}
    v22 := GenS4{F0: 0, F1: v3, F2: 56}
    for fi23 := 0; fi23 < 5; fi23++ {
        v2 = v2
        v2 += v2
        v2 = 156
        _ = fi23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 int16 = 1
            var v27 uint8 = (v2 % ((v2) | 1))
            _ = fi25
            _ = v26
            _ = v27
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch fi24 {
        case 42:
            fmt.Println(v3)
        }
        _ = fi24
    }
    switch 0 {
    case 127:
        v28 := GenS2{F0: 0, F1: 65535}
        _ = v28
    case -155:
        var v29 bool = ("" != "")
        _ = v29
    case -433:
        v2 += 127
    default:
        v30 := GenS1{F0: ((425 > 0) && true), F1: -9223372036854775808, F2: 1}
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 float32 = float32((-2.25 / 1e10))
        if (0 < (fi31 % ((fi31) | 1))) {
            var v33 float32 = float32(((((3.14159 - 3.14159) / -2.25) - 1.5) / 0.1))
            var v34 uint32 = 2147483648
            _ = v33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(v3, v2)
    v2 += (v2 % ((v2) | 1))
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v8
}
