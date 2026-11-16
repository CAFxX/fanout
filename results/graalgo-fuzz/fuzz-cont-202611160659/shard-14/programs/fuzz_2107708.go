package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint32 }

func main() {
    var v1 uint16 = 255
    var v2 float64 = 0.1
    var v3 float64 = 0.1
    if ((int8(v1) != 42) || false) {
        if true {
            if (-786 == -581) {
                v4 := GenS1{F0: 566, F1: (v1 / (((^(v1))) | 1)), F2: 2147483647}
                if ((false && (((-2.25 >= ((1e10 * -2.25) / 1.5)) || (391 < 228)) || false)) && true) {
                    var v5 int64 = 0
                    v5 += v5
                    switch "a" {
                    case "\n\t":
                        v6 := GenS1{F0: 2147483647, F1: v1, F2: 1}
                        _ = v6
                    }
                    _ = v5
                }
                _ = v4
            }
            v3 -= (((((-2.25 + v2) + (v3 - v2)) * v2) + v3) * float64((0.0 + (0.0 - 0.1))))
        } else {
            v7 := []int{-2147483648}
            v8 := GenS1{F0: 0, F1: v1, F2: 1}
            _ = v7
            _ = v8
        }
        v2 -= v3
    }
    v9 := GenS1{F0: 0, F1: 567, F2: 193}
    v1 -= (v1 & v1)
    var v10 int8 = -74
    var v11 uint64 = 1
    var v12 int16 = 1
    v13 := GenS1{F0: -140, F1: (v1 / ((42) | 1)), F2: 4294967295}
    var v14 int32 = 915
    v11 = (^(v11))
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(v2, fi15, v2)
        _ = fi15
    }
    v11 = (v11 / ((v11) | 1))
    if false {
        v16 := []int{594, 127, 993}
        v17 := GenS2{F0: (v1 >> ((65535) & 63))}
        _ = v16
        _ = v17
    }
    if true {
        fmt.Println(v11, v11)
    }
    var v18 int32 = (v14 ^ -689)
    if false {
        v19 := GenS3{F0: 65535, F1: 32767, F2: 42}
        var v20 bool = ((!(128 >= 2)) && (!(-2.25 <= ((1.5 * ((3.14159 + (1e10 * -2.25)) * 0.1)) + 1e10))))
        fmt.Println(v12, v11, v18)
        _ = v19
        _ = v20
    }
    var v21 uint64 = v11
    v22 := []int{128, 127, 197, int(v12)}
    for fi23 := 0; fi23 < 5; fi23++ {
        _ = len(v22)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
