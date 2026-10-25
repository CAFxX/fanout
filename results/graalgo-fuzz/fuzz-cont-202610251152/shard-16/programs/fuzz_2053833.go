package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    var v1 string = "\n\t"
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    if false {
        fmt.Println(v1)
    }
    v4 := GenS1{F0: int16(v2), F1: 1}
    if true {
        if (310 == -2147483648) {
            var v5 uint32 = 2147483647
            fmt.Println(v5, v2, v2)
            _ = v5
        } else {
            v6 := v3[752]
            if (true || true) {
                var v7 uint64 = 4294967295
                _ = v7
            } else {
                v1 = v1
                for fi8 := 0; fi8 < 14; fi8++ {
                    switch ("" + ("\n\t" + (v1 + (v1 + ("" + v1))))) {
                    case "foo bar":
                        v9 := GenS2{F0: -9223372036854775808, F1: -407}
                        _ = v9
                    case "":
                        var v10 int = fi8
                        _ = v10
                    case "\n\t":
                        v1 = v1
                    }
                    v6 = 715
                    _ = fi8
                }
                v6 -= v6
            }
            _ = v6
        }
        v3[0] = 0
    }
    v11 := GenS2{F0: 0, F1: 42}
    v12 := GenS1{F0: 0, F1: 0}
    if true {
        v1 = (v1 + v1)
        if (100 < -5) {
            var v13 uint16 = 255
            v14 := GenS1{F0: int16(v2), F1: 18446744073709551615}
            _ = v13
            _ = v14
        } else {
            var v15 int8 = 0
            _ = v15
        }
        if (925 == -1) {
            var v16 float64 = (((((3.14159 * (((float64(1e10) + float64((float32(3.14159) + float32(3.14159)))) + (0.0 - float64((3.14159 + 3.14159)))) - 0.1)) / (0.0 - 1e10)) / (1e10 + 1.5)) * float64(-2.25)) * 3.14159)
            v2 = 0
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS2{F0: 0, F1: -746}
            var v19 int32 = -2147483648
            var v20 int8 = 108
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        v2 = (^(v2))
        switch "\n\t" {
        case "":
            v21 := GenS1{F0: -1, F1: 42}
            _ = v21
        case "a":
            var v22 int = 128
            _ = v22
        default:
            for fi23 := 0; fi23 < 6; fi23++ {
                if true {
                    _ = len(v3)
                } else {
                    v2 = v2
                    v2 = v2
                }
                _ = fi23
            }
        }
    }
    fmt.Println(v2)
    v24 := v3[2147483647]
    var v25 float64 = (3.14159 + 3.14159)
    var v26 int64 = 42
    var v27 float32 = (0.1 * 0.0)
    var v28 string = "foo bar"
    v29 := []int{2147483647, 32768, 9223372036854775807, 9223372036854775807}
    v30 := GenS2{F0: v26, F1: 809}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
