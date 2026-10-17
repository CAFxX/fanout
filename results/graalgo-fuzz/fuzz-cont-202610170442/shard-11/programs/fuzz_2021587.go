package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint32 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 bool }
func genFn0() uint {
    var v7 float64 = ((((0.0 / float64(((1e10 - (((1.5 - (1e10 * 0.0)) + 1.5) + 0.1)) - 0.0))) - (0.0 + (3.14159 * -2.25))) + 0.1) + (((((((0.0 + (3.14159 + 3.14159)) - 1.5) * 3.14159) * 0.1) + 3.14159) / -2.25) - (1e10 - (3.14159 - (-2.25 + 0.1)))))
    for fi8 := 0; fi8 < 12; fi8++ {
        switch 120 {
        case -827:
            fmt.Println(fi8)
        case -613:
            v7 += v7
        case 9223372036854775807:
            if (-909 >= 9223372036854775807) {
                v7 = ((float64(0.1) * float64(1.5)) * v7)
            }
        default:
            v9 := []int{127, (fi8 & fi8), fi8, (fi8 << ((255) & 63))}
            _ = v9
        }
        _ = fi8
    }
    var v10 float32 = -2.25
    var v11 uint16 = 255
    _ = v10
    _ = v11
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = 1e10
    var v3 int64 = -191
    v4 := make(map[uint32]bool)
    var v5 int = 228
    var v6 float64 = (v2 / v2)
    switch "\n\t" {
    case "hello":
        var v12 uint16 = 1
        _ = v12
    default:
        v13 := v4[732]
        _ = v13
    }
    var v14 bool = true
    v14 = true
    if (v3 >= v3) {
        v3 -= v3
    }
    v1 -= v1
    v15 := GenS1{F0: 1000, F1: 990, F2: (v1 ^ (v1 - v1))}
    v14 = v14
    var v16 uint64 = 2
    if (!v14) {
        v1 = 2
    } else {
        fmt.Println(v5, v5)
        var v17 int = v5
        v18 := []int{v17, (-(v17)), 42}
        _ = v17
        _ = v18
    }
    if (v2 == (v6 / v2)) {
        var v19 uint64 = (v16 + 2)
        _ = v19
    } else {
        fmt.Println(v16, v3, v3)
        v6 = float64(float32(v6))
    }
    switch (-(v5)) {
    case 4294967295:
        v6 = (((v6 + (float64(-2.25) + (1.5 * 0.0))) * v6) * (v6 / float64(0.1)))
    case 127:
        fmt.Println(v14, v14)
    }
    v20 := GenS2{F0: 2147483647}
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint32 = 2147483647
        v23 := GenS3{F0: 0, F1: -2147483648, F2: (v22 % (((v22 | v22)) | 1))}
        var v24 int64 = (v3 / ((9223372036854775807) | 1))
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS1{F0: 0, F1: 255, F2: v1}
            v5 -= fi26
            switch "hello" {
            case "foo bar":
                fmt.Println(fi25, v5)
                v4[284] = v14
                v28 := GenS4{F0: fi25, F1: -1}
                if ((v16 / (((v16 + (v16 + (v16 | v16)))) | 1)) <= v16) {
                    switch "" {
                    case "\n\t":
                        var v29 int16 = -972
                        _ = v29
                    default:
                        v3 = v3
                    }
                    v30 := GenS5{F0: 1, F1: 255, F2: (!true)}
                    var v31 float64 = v6
                    _ = v30
                    _ = v31
                }
                _ = v28
            default:
                v1 -= v1
            }
            _ = fi26
            _ = v27
        }
        var v32 int8 = int8(v3)
        _ = fi25
        _ = v32
    }
    v2 = ((3.14159 - v2) - v6)
    v3 += 0
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
