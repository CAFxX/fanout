package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    v1 := []uint32{2, 323, 42}
    v2 := make(map[uint32]uint32)
    var v3 float64 = ((float64(float32(3.14159)) * (3.14159 - 3.14159)) / (1e10 + 1.5))
    delete(v2, 211)
    var v4 int16 = -528
    if (100 >= 100) {
        if (0 >= -1) {
            v4 -= (v4 ^ (-(v4)))
            switch 0 {
            case 1:
                for fi5 := 0; fi5 < 15; fi5++ {
                    if false {
                        for fi6 := 0; fi6 < 11; fi6++ {
                            v7 := GenS1{F0: (fi6 + fi6), F1: 42, F2: 255}
                            _ = fi6
                            _ = v7
                        }
                    } else {
                        for fi8 := 0; fi8 < 3; fi8++ {
                            v4 += v4
                            var v9 uint16 = 255
                            _ = fi8
                            _ = v9
                        }
                    }
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        delete(v2, 2)
                    case "a":
                        v10 := GenS1{F0: (fi5 & fi5), F1: -564, F2: 255}
                        _ = v10
                    }
                    _ = len(v2)
                    _ = fi5
                }
            default:
                if (65535 == 2147483647) {
                    v4 = (v4 / ((978) | 1))
                    var v11 int16 = v4
                    v11 -= -32768
                    _ = v11
                }
            }
            if (false || (-1 != 1)) {
                if (!false) {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v3 = float64(0.1)
                        if true {
                            fmt.Println(fi12, v3, fi12)
                            var v13 int64 = 857
                            v13 += 0
                            _ = v13
                        }
                        _ = fi12
                    }
                }
                v4 -= v4
            }
        }
        fmt.Println(v4, v3)
        v3 -= float64(3.14159)
    }
    var v14 int = -9223372036854775808
    fmt.Println(v3, v14)
    if ((true || true) || (!false)) {
        switch "foo bar" {
        case "hello":
            if true {
                if false {
                    if (true || (1 != 286)) {
                        delete(v2, 1)
                        v2[1] = 0
                        for fi15 := 0; fi15 < 9; fi15++ {
                            var v16 int16 = (v4 & -132)
                            _ = fi15
                            _ = v16
                        }
                    }
                    v17 := GenS1{F0: v14, F1: -528, F2: 784}
                    _ = v17
                }
            } else {
                if true {
                    var v18 float64 = (0.0 * ((v3 / 1.5) / v3))
                    var v19 int64 = 753
                    _ = v18
                    _ = v19
                }
                v4 = (v4 & v4)
                v14 += v14
            }
        case "foo bar":
            v20 := GenS1{F0: 9223372036854775807, F1: -628, F2: 32767}
            _ = v20
        default:
            v21 := GenS2{F0: 694, F1: v4, F2: 18446744073709551615}
            _ = v21
        }
        v22 := v2[2147483647]
        _ = v22
    }
    v3 += (float64((((1e10 + 0.1) - 1e10) / -2.25)) - v3)
    v23 := v2[65535]
    v2[(v23 % (((v23 / (((v23 << ((211) & 63))) | 1))) | 1))] = (v23 & v23)
    v4 -= 416
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
