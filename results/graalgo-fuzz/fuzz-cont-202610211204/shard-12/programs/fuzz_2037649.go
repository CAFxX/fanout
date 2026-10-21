package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 uint32 }

func main() {
    var v1 int = -500
    v2 := make(map[uint32]string)
    v3 := []int{(v1 & v1), (v1 << ((18446744073709551615) & 63))}
    var v4 uint32 = 42
    v5 := make(map[string]int)
    v6 := GenS1{F0: -566, F1: 1, F2: -1}
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS2{F0: 238}
        _ = fi7
        _ = v8
    }
    if false {
        v1 = (^(v1))
        if false {
            switch (v1 * v1) {
            case 343:
                _ = len(v3)
            case 649:
                var v9 int8 = -19
                _ = v9
            case 4294967295:
                v3 = append(v3, (v1 / (((v1 / (((-(v1))) | 1))) | 1)))
            default:
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 uint8 = 42
                    if false {
                        v11 -= (^(v11))
                    }
                    var v12 int = -144
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            v5["\n\t"] = v1
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v5)
            if (3.14159 < 1e10) {
                var v15 uint16 = 2
                _ = v15
            }
            _ = fi14
        }
        var v16 int32 = -238
        _ = fi13
        _ = v16
    }
    var v17 int8 = 95
    v1 -= v1
    var v18 int16 = -360
    v1 += (v1 | v1)
    var v19 uint8 = 54
    v20 := GenS3{F0: -2147483648, F1: 2}
    switch "foo bar" {
    case "hello":
        if (102 == v4) {
            v21 := GenS4{F0: "hello", F1: 42, F2: (^(v4))}
            v1 -= 128
            _ = v21
        }
    case "":
        var v22 uint8 = v19
        _ = v22
    case "\n\t":
        v17 = (v17 / ((0) | 1))
    }
    v18 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 int64 = -9223372036854775808
    var v24 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
