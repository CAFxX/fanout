package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int = -669
    v2 := []int{(-(v1))}
    v3 := []uint32{2147483647, 238}
    var v4 float32 = (-2.25 / ((-2.25 - (0.1 * 1.5)) / 3.14159))
    fmt.Println(v4, v4)
    var v5 int = (^(v1))
    var v6 float64 = 1.5
    v6 += (((v6 / v6) * v6) * float64((0.1 * float32((v6 / (0.0 - v6))))))
    v3 = append(v3, 776)
    if false {
        _ = len(v2)
        switch v5 {
        case 127:
            if (false && true) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    var v8 bool = (!true)
                    _ = fi7
                    _ = v8
                }
            } else {
                var v9 uint32 = 2147483647
                v10 := GenS1{F0: 544, F1: 4294967295}
                fmt.Println(v5)
                _ = v10
                _ = v9
            }
        case -708:
            v1 -= (v5 & v5)
        case 128:
            switch v1 {
            case 55:
                if (!(4294967295 < 4294967295)) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        fmt.Println(v4, v1)
                        var v12 uint8 = 115
                        v3 = append(v3, 972)
                        _ = fi11
                        _ = v12
                    }
                    v1 = (v1 - (v1 ^ (-(v1))))
                    var v13 uint = 1
                    _ = v13
                }
            case -1:
                v2 = append(v2, (v5 << ((4294967295) & 63)))
            case 486:
                if ((v1 << ((42) & 63)) != 9223372036854775807) {
                    var v14 uint8 = 0
                    fmt.Println(v4, v1)
                    switch v1 {
                    case 4294967295:
                        for fi15 := 0; fi15 < 2; fi15++ {
                            for fi16 := 0; fi16 < 16; fi16++ {
                                v17 := GenS1{F0: 2147483647, F1: 65535}
                                v18 := GenS1{F0: 1, F1: 647}
                                v6 -= v6
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            if len(v2) > 0 {
                                v2[0] = 127
                            }
                            _ = fi15
                        }
                    case 2147483647:
                        var v19 bool = (0 == -306)
                        _ = v19
                    default:
                        v2 = append(v2, (v5 << ((65535) & 63)))
                    }
                    _ = v14
                } else {
                    v6 = (((((v6 - (v6 - (v6 + (v6 - float64(v4))))) + ((1e10 - 0.1) * v6)) + v6) + (3.14159 + float64(((float32((((0.0 + ((v6 - v6) + v6)) * v6) * v6)) + v4) + (1e10 / v4))))) / 1e10)
                    fmt.Println(v6, v6, v4)
                    v5 += 1
                }
            default:
                var v20 uint64 = 4294967295
                _ = v20
            }
        default:
            v5 += 1
        }
    } else {
        var v21 uint64 = 0
        var v22 uint32 = 1
        _ = v21
        _ = v22
    }
    var v23 uint16 = 0
    if (128 != 2) {
        v24 := GenS2{F0: -773, F1: 18446744073709551615}
        v4 -= v4
        _ = v24
    }
    fmt.Println(v23)
    v25 := GenS2{F0: 17, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
