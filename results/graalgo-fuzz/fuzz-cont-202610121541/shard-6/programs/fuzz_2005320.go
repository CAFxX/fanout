package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int32 }
type GenS4 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 uint64) uint8 {
    var v4 int64 = -617
    fmt.Println(v4, v4)
    _ = p0
    _ = v4
    return 189
}

func main() {
    var v1 uint64 = 293
    var v2 float64 = (-2.25 + (0.1 + 1.5))
    v3 := []uint32{4294967295, 65535}
    var v5 int16 = -413
    var v6 float32 = (-2.25 * -2.25)
    var v7 float64 = v2
    if ((^(v5)) != v5) {
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    if true {
        for fi8 := 0; fi8 < 9; fi8++ {
            if true {
                if (true && (true || (236 >= 42))) {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    fmt.Println(v6)
                }
                v9 := GenS1{F0: 2, F1: 2}
                var v10 bool = (127 > 154)
                _ = v10
                _ = v9
            }
            fmt.Println(v5)
            if (false && (!false)) {
                fmt.Println(v1, fi8)
                var v11 uint32 = 928
                v12 := GenS2{F0: uint8(v11)}
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 float32 = v6
                    v6 = (0.1 + (v6 - v6))
                    v2 -= (1e10 - ((0.1 + v2) - ((0.0 * (v7 + float64(v6))) - 1e10)))
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v7, fi8)
                fmt.Println(v6, v1)
            }
            _ = fi8
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 = ((v6 - 3.14159) - float32(float64((v6 * v6))))
            _ = fi15
        }
        _ = len(v3)
    } else {
        var v16 uint16 = 369
        _ = v16
    }
    v6 = (0.1 / ((((v6 + float32(v7)) - (v6 + (v6 + 1e10))) - (v6 + -2.25)) + 1e10))
    v17 := GenS1{F0: genFn0(18446744073709551615), F1: 53}
    if (!((1 <= 4294967295) && ((v5 | (v5 | v5)) < 401))) {
        var v18 uint32 = 2147483648
        v3 = append(v3, (v18 * (^(v18))))
        _ = v18
    }
    if true {
        var v19 int16 = v5
        fmt.Println(v2, v6, v6)
        v6 = v6
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v21 := GenS1{F0: 0, F1: 65535}
        _ = fi20
        _ = v21
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v22 := GenS3{F0: ("foo bar" + "\n\t"), F1: (v5 & (v5 * (v5 ^ v5))), F2: 0}
    v23 := GenS4{F0: 4294967295, F1: 1, F2: false}
    v24 := GenS3{F0: "foo bar", F1: (-(v5)), F2: 0}
    fmt.Println(v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
