package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0() int16 {
    var v5 int32 = 958
    var v6 uint8 = 0
    var v7 uint32 = 82
    var v8 bool = false
    fmt.Println(v7, v8)
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -448
}

func main() {
    var v1 int32 = 42
    var v2 bool = (true && ((18446744073709551615 >= 952) || true))
    var v3 float64 = 3.14159
    var v4 int = 128
    if (!false) {
        switch "" {
        case "":
            v4 = (^(v4))
        case "\n\t":
            v4 = (v4 & v4)
        case "a":
            v9 := GenS1{F0: 42, F1: 472, F2: 255}
            _ = v9
        }
        fmt.Println(v3, v4, v3)
    }
    var v10 int = (v4 | 42)
    var v11 int8 = -85
    var v12 uint16 = 255
    var v13 uint64 = 0
    v14 := GenS1{F0: 9223372036854775807, F1: -697, F2: 234}
    v1 += (v1 | (-(v1)))
    var v15 int16 = 42
    if (2 >= 255) {
        for fi16 := 0; fi16 < 11; fi16++ {
            v10 += 1
            var v17 uint = 4294967295
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: 1}
        v19 := GenS2{F0: 803}
        _ = v18
        _ = v19
    }
    if (!false) {
        v10 += int(v1)
        var v20 int32 = (v1 / ((v1) | 1))
        for fi21 := 0; fi21 < 18; fi21++ {
            if v2 {
                var v22 float64 = float64(((((((float32(0.1) + (1e10 * 1.5)) - (0.0 + ((0.0 * -2.25) / (((float32(1e10) - (3.14159 - 0.0)) * 0.1) + 1e10)))) * ((-2.25 * -2.25) - 0.0)) + 3.14159) * 1.5) - 3.14159))
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint64 = (v13 << ((2) & 63))
                    switch "" {
                    case "foo bar":
                        var v25 int8 = v11
                        _ = v25
                    case "\n\t":
                        switch -1 {
                        case 75:
                            var v26 bool = (v24 >= v24)
                            _ = v26
                        }
                    }
                    var v27 int8 = -128
                    _ = fi23
                    _ = v24
                    _ = v27
                }
                _ = v22
            }
            var v28 bool = (900 == 229)
            _ = fi21
            _ = v28
        }
        _ = v20
    }
    v29 := []int{v10, (v10 ^ (v10 - (v10 ^ 127)))}
    if (true || (13 != 100)) {
        v1 -= -2147483648
        var v30 int8 = (v11 & (v11 ^ (v11 * v11)))
        switch v4 {
        case -1:
            var v31 int16 = 0
            _ = v31
        default:
            var v32 float32 = 1.5
            _ = v32
        }
        _ = v30
    } else {
        v4 -= -662
        v4 = v4
        switch v10 {
        case -505:
            var v33 uint8 = 127
            _ = v33
        case -1:
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 uint64 = v13
                _ = fi34
                _ = v35
            }
        case 355:
            fmt.Println(v3, v15, v13)
        default:
            for fi36 := 0; fi36 < 9; fi36++ {
                v10 -= v4
                _ = fi36
            }
        }
    }
    v1 = v1
    v37 := GenS1{F0: -1, F1: -356, F2: 2}
    var v38 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
