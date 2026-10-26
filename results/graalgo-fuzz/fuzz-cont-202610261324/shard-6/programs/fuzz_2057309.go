package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }

func main() {
    var v1 uint8 = 255
    v2 := []uint32{1}
    var v3 float64 = (float64(((1e10 * 0.0) / 1.5)) * (1.5 - 1e10))
    var v4 uint8 = 42
    var v5 int8 = 83
    var v6 uint8 = (v4 ^ v4)
    if true {
        var v7 uint8 = (v4 & v6)
        switch 441 {
        case 32767:
            v5 = v5
        }
        _ = v7
    } else {
        v6 = (v1 >> ((255) & 63))
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                _ = len(v2)
                _ = fi9
            }
            v1 -= v6
            _ = fi8
        }
        v3 -= v3
    }
    switch 32767 {
    case 1:
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(v1, v4, v4)
            _ = fi10
        }
    default:
        var v11 float32 = float32((((3.14159 * (v3 / (float64(0.0) / v3))) / (1.5 - (v3 * v3))) - v3))
        _ = v11
    }
    var v12 bool = true
    v13 := GenS1{F0: 1}
    if v12 {
        v14 := GenS2{F0: 790, F1: v5, F2: v1}
        v2 = append(v2, 584)
        _ = v14
    } else {
        v5 = v5
    }
    if true {
        v1 = v4
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v15 := GenS1{F0: 839}
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            var v18 int16 = 32767
            if (3.14159 < 1e10) {
                switch fi17 {
                case 127:
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                default:
                    var v19 float64 = 0.1
                    _ = v19
                }
            } else {
                v20 := GenS3{F0: 65535, F1: ((!v12) || (v3 < v3)), F2: 42}
                _ = v20
            }
            v3 = v3
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v21 uint = 69
    var v22 int64 = int64(v5)
    v22 += v22
    v12 = (1e10 == float32(float64(float32(v3))))
    var v23 int16 = 39
    v3 = 0.1
    switch 122 {
    case 1:
        var v24 uint16 = 65535
        _ = v24
    case 710:
        var v25 int8 = (v5 ^ v5)
        _ = v25
    case -2147483648:
        if (!v12) {
            v21 = 972
        }
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v27 int16 = -32768
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
