package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    var v1 uint16 = 65535
    v2 := []int{655, 0, -856, -9223372036854775808}
    var v3 uint32 = 42
    var v4 int16 = 1
    var v5 int64 = 0
    _ = len(v2)
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint32 = (^(v3))
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 string = "a"
            v1 += (v1 << ((18446744073709551615) & 63))
            var v10 int = fi8
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS1{F0: v5, F1: (v5 * (v5 * (v5 / ((v5) | 1)))), F2: 127}
    var v12 int32 = 0
    var v13 float32 = (3.14159 - float32(1e10))
    if false {
        var v14 uint32 = 122
        _ = v14
    } else {
        if (!true) {
            if (740 > 18446744073709551615) {
                v3 -= (v3 | (^(v3)))
                if len(v2) > 0 {
                    v2[0] = -2147483648
                }
                v15 := GenS2{F0: (v12 % ((v12) | 1)), F1: 2, F2: v5}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 7; fi16++ {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v13 = v13
                        var v18 float64 = (3.14159 + (float64(((((float32(1e10) * ((v13 - ((float32(float64((float32(0.0) * (v13 + ((((3.14159 + (v13 - (0.0 / 1.5))) - v13) * float32(0.1)) - v13))))) / v13) / 1.5)) + v13)) * v13) - float32((1e10 * (1.5 / (1.5 - (float64(-2.25) - -2.25)))))) - 1e10)) / ((float64(float32(1.5)) * -2.25) - -2.25)))
                        v3 = (v3 / ((v3) | 1))
                        _ = fi17
                        _ = v18
                    }
                    v19 := GenS2{F0: v12, F1: 226, F2: v5}
                    _ = fi16
                    _ = v19
                }
                v1 += (^(v1))
                v3 = uint32(v1)
            }
            if (false && false) {
                var v20 uint16 = v1
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
                _ = v20
            } else {
                fmt.Println(v1, v5)
                if len(v2) > 0 {
                    v2[0] = -2147483648
                }
                var v21 uint8 = 42
                _ = v21
            }
        }
        v22 := GenS3{F0: uint8(v3), F1: 54}
        if len(v2) > 0 {
            v2[0] = -1
        }
        _ = v22
    }
    v12 += (v12 >> ((954) & 63))
    v23 := GenS2{F0: v12, F1: 0, F2: v5}
    var v24 int8 = -36
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
