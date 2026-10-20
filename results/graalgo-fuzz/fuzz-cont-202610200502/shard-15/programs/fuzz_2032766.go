package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0(p0 int64) uint16 {
    v4 := GenS1{F0: 585}
    fmt.Println(p0)
    v5 := GenS2{F0: 42}
    v6 := GenS3{F0: 42, F1: 65535}
    var v7 uint32 = 65535
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int = 9223372036854775807
    var v3 int16 = -354
    switch (v2 % ((0) | 1)) {
    case 4294967295:
        v1 = (-(v1))
    case 1:
        v1 -= -294
    default:
        var v8 uint32 = 65535
        _ = v8
    }
    v9 := GenS2{F0: 128}
    switch (v2 ^ (^(v2))) {
    case 701:
        v2 = v2
    case -9223372036854775808:
        var v10 float64 = 0.1
        _ = v10
    }
    v1 = v1
    if true {
        var v11 int32 = -369
        if true {
            v12 := []int{v2, v1}
            _ = v12
        } else {
            var v13 int32 = (v11 / (((v11 / (((v11 ^ v11)) | 1))) | 1))
            v3 -= (v3 >> ((65535) & 63))
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            if false {
                var v15 int16 = (v3 ^ v3)
                v15 = v3
                _ = v15
            } else {
                v16 := []int{}
                var v17 float32 = ((float32(-2.25) * (-2.25 / 0.1)) * 0.0)
                _ = v16
                _ = v17
            }
            v18 := GenS4{F0: 1, F1: (1 != 0), F2: 4294967295}
            _ = fi14
            _ = v18
        }
        _ = v11
    }
    switch v1 {
    case 0:
        v2 -= (v2 | (v1 >> ((480) & 63)))
    case -784:
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int = fi19
            if (2147483647 == -2147483648) {
                v1 = -124
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 int8 = 0
                    if false {
                        fmt.Println(fi21)
                    }
                    v2 -= v20
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi19
            _ = v20
        }
    default:
        v3 -= v3
    }
    var v23 string = "\n\t"
    fmt.Println(v1, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v9
}
