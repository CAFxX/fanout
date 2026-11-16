package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint64; F2 uint8 }

func main() {
    v1 := []uint32{42, 266, 712, 42}
    var v2 uint32 = 0
    var v3 uint32 = (v2 ^ v2)
    var v4 uint16 = 500
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int8 = 127
        v2 -= 42
        switch "" {
        case "hello":
            v6 = v6
        }
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        if (9223372036854775807 > -960) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v8 int64 = -9223372036854775808
            var v9 uint64 = 18446744073709551615
            _ = v8
            _ = v9
        } else {
            var v10 int8 = -128
            if (!true) {
                fmt.Println(v2)
                if (("" + "a") > "\n\t") {
                    if true {
                        v2 = (v2 % ((v2) | 1))
                        fmt.Println(v2, v4)
                    }
                    if (!false) {
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                    } else {
                        if false {
                            var v11 int64 = 9223372036854775807
                            _ = v11
                        }
                        var v12 int32 = 2147483647
                        _ = v12
                    }
                }
                switch (fi7 - fi7) {
                case -423:
                    v13 := GenS1{F0: fi7}
                    _ = v13
                }
            }
            _ = v10
        }
        if false {
            v2 -= v2
            v14 := GenS1{F0: fi7}
            v15 := GenS1{F0: (^(fi7))}
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    if false {
        var v16 uint64 = 4294967295
        for fi17 := 0; fi17 < 16; fi17++ {
            fmt.Println(v2, v2, v4)
            v18 := GenS1{F0: 4294967295}
            v16 = (v16 * (^(v16)))
            _ = fi17
            _ = v18
        }
        var v19 uint8 = 6
        _ = v16
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: 4294967295, F2: 107}
    if (0.0 <= ((3.14159 - (((float64(-2.25) / (((-2.25 * 3.14159) - 0.0) / float64(((1e10 - 3.14159) / (float32(((((1.5 + (-2.25 / 1e10)) - 0.1) * (((float64(3.14159) - 0.0) * -2.25) + -2.25)) * float64(0.1))) * -2.25))))) + (1e10 / 1.0)) + 0.0)) * (0.0 - 0.0))) {
        v3 -= (v3 | (v3 | 0))
        v1 = append(v1, (v2 | 854))
    } else {
        var v21 float32 = 0.1
        v4 += v4
        switch 42 {
        case -674:
            v22 := GenS2{F0: 373, F1: 4294967295, F2: 42}
            _ = v22
        case 9223372036854775807:
            fmt.Println(v4, v21)
        case 350:
            var v23 int32 = 2147483647
            _ = v23
        default:
            v4 += v4
        }
        _ = v21
    }
    if (17 <= 178) {
        v2 += v3
    }
    v3 -= (^(v2))
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 string = ("" + "a")
        var v26 uint8 = 243
        v4 = (v4 - 255)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
