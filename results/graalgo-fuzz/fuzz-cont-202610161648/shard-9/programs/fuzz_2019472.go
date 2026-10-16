package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 297
    v2 := []int{}
    var v3 int = 0
    v1 += (v1 % (((v1 - 51)) | 1))
    fmt.Println(v1)
    if (1e10 == 1.5) {
        var v4 bool = false
        v1 = v1
        var v5 float32 = 3.14159
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v1, fi6)
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(fi6)
                v3 = v3
                _ = fi7
            }
            _ = fi6
        }
        if (!((0.0 <= -2.25) || false)) {
            v3 = (-(v3))
            v3 += -2147483648
        }
    }
    v1 -= 4294967295
    switch v3 {
    case 691:
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 42, F1: -1}
            _ = fi8
            _ = v9
        }
    case 9223372036854775807:
        var v10 uint8 = 2
        _ = v10
    case 128:
        fmt.Println(v1)
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    v1 -= v1
    v3 -= v3
    v3 -= (-(v3))
    v12 := GenS2{F0: 477, F1: 441}
    v3 = v3
    v1 = v1
    switch (v3 | -2147483648) {
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch ("" + "a") {
    case "\n\t":
        v1 = 18446744073709551615
    case "":
        if true {
            var v13 uint = v1
            _ = v13
        } else {
            if false {
                for fi14 := 0; fi14 < 4; fi14++ {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 uint8 = 130
                        fmt.Println(fi14, fi14, v3)
                        _ = fi15
                        _ = v16
                    }
                    v1 -= 223
                    switch 127 {
                    case -2147483648:
                        v3 = (v3 + 732)
                    case 4294967295:
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                    }
                    _ = fi14
                }
                v17 := GenS2{F0: (^(v1)), F1: 0}
                if len(v2) > 0 {
                    v2[0] = v3
                }
                _ = v17
            }
            v18 := GenS1{F0: 901, F1: -128}
            _ = v18
        }
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v19 int16 = 0
        _ = v19
    }
    fmt.Println(v1, v1)
    v1 = (v1 / (((^(v1))) | 1))
    v20 := GenS1{F0: 2, F1: 127}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
}
