package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 626
    var v3 int8 = 127
    v4 := GenS1{F0: 18446744073709551615, F1: 42}
    var v5 uint8 = 193
    v5 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 bool = (false || (1 < v1))
        _ = fi6
        _ = v7
    }
    if (!(v3 < -128)) {
        for fi8 := 0; fi8 < 13; fi8++ {
            v5 = v5
            v9 := []int{fi8, (fi8 - (fi8 >> ((984) & 63))), (fi8 << ((0) & 63))}
            if (true || true) {
                if true {
                    if len(v9) > 0 {
                        v9[0] = fi8
                    }
                    var v10 bool = (v1 != (-(v1)))
                    var v11 float32 = 1.5
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS2{F0: (fi8 << ((0) & 63)), F1: 65535, F2: (^(v2))}
                    var v13 int64 = 42
                    _ = v12
                    _ = v13
                }
            } else {
                v2 += v2
                if true {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if (!false) {
                            var v15 int8 = v3
                            var v16 float64 = 0.0
                            var v17 int16 = int16(v5)
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                } else {
                    var v18 int32 = 0
                    _ = v18
                }
                fmt.Println(v1)
            }
            _ = fi8
            _ = v9
        }
    } else {
        v19 := GenS3{F0: ((!false) && false), F1: 1, F2: (v1 & (v1 * v1))}
        v1 = (v1 >> ((189) & 63))
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v5 += (v5 | (v5 ^ uint8(v2)))
        var v21 float64 = (1e10 * ((0.1 + 3.14159) * (1e10 - float64(1e10))))
        fmt.Println(fi20, v2, v1)
        _ = fi20
        _ = v21
    }
    v3 -= v3
    v2 += 228
    for fi22 := 0; fi22 < 1; fi22++ {
        v23 := GenS1{F0: 473, F1: 42}
        _ = fi22
        _ = v23
    }
    v24 := []int{}
    v1 -= (v1 << ((1) & 63))
    var v25 float32 = 1e10
    v24 = append(v24, 32767)
    fmt.Println(v25, v2, v1)
    var v26 uint = 146
    var v27 int8 = (v3 << (((v26 % ((v26) | 1))) & 63))
    switch "" {
    case "\n\t":
        fmt.Println(v1, v27, v26)
    default:
        _ = len(v24)
    }
    v1 += (v1 >> (((v26 + (v26 - v26))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
