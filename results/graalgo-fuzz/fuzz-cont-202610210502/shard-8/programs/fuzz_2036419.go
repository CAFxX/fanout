package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64; F1 int16; F2 uint }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int8 = 42
    var v2 string = ""
    v3 := make(map[int]bool)
    v4 := []string{"", "\n\t"}
    var v5 bool = false
    var v6 uint8 = 0
    v1 = 127
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi8
        }
        _ = fi7
    }
    v1 = -70
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = "foo bar"
        var v11 float32 = (0.1 - ((float32(((((float64(-2.25) / 1.0) - 0.1) * 1e10) / 1e10)) - 0.1) / ((float32(1.5) - 1.5) - 1.5)))
        var v12 uint8 = (v6 & (v6 * (v6 ^ 127)))
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    if (((^(v6)) > (v6 + v6)) || false) {
        var v13 uint = 4294967295
        if (323 <= 1) {
            switch -503 {
            case 4294967295:
                switch 9223372036854775807 {
                case -9223372036854775808:
                    fmt.Println(v13)
                case 9223372036854775807:
                    switch -667 {
                    case -310:
                        v14 := GenS1{F0: 9223372036854775807, F1: -840}
                        _ = v14
                    case -901:
                        var v15 float32 = 1e10
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v13 -= v13
                            _ = fi16
                        }
                    }
                default:
                    delete(v3, -411)
                }
            case 2147483647:
                fmt.Println(v2)
            case 0:
                var v17 int32 = 0
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 float64 = ((3.14159 + 1.5) * 3.14159)
                    var v20 uint = (^(v13))
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 int32 = -2147483648
                        _ = fi21
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        } else {
            fmt.Println(v13)
            _ = len(v4)
        }
        v13 += 240
        _ = v13
    } else {
        v6 -= v6
        v23 := v3[128]
        fmt.Println(v2, v6, v23)
        _ = v23
    }
    var v24 int8 = -20
    v24 = v24
    delete(v3, 128)
    v6 -= (v6 + (v6 * 199))
    v1 = v1
    switch (v2 + v2) {
    case "\n\t":
        v25 := GenS2{F0: 1000, F1: 543}
        _ = v25
    default:
        delete(v3, 354)
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v5 = ((1.5 / 0.1) > float32((((0.1 / 1.0) - 1.5) - 0.1)))
        var v27 uint = 0
        _ = fi26
        _ = v27
    }
    if v5 {
        v5 = ((v1 / ((0) | 1)) < v24)
        v2 = v2
    } else {
        if ((32767 >= 644) || (-2.25 <= (0.1 - (float64((1.5 * 0.1)) / ((float64((0.1 - 0.0)) + (1e10 + 3.14159)) - 1e10))))) {
            if (0.1 <= (1.5 / 1e10)) {
                v28 := GenS3{F0: 9223372036854775807, F1: -1, F2: 1}
                v6 = v6
                _ = v28
            }
            v24 = v1
            delete(v3, 393)
        }
    }
    v29 := GenS4{F0: v5, F1: 1}
    if false {
        v6 -= 255
        if ((v24 ^ v24) == (v1 - (v1 >> ((2) & 63)))) {
            v30 := GenS5{F0: 1}
            switch 54 {
            case 4294967295:
                v1 = v1
            default:
                var v31 float64 = 3.14159
                _ = v31
            }
            _ = v30
        }
        fmt.Println(v5, v24)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
