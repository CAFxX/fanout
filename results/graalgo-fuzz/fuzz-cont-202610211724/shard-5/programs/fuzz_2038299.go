package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }

func main() {
    var v1 int = -957
    var v2 uint16 = 851
    var v3 float32 = float32(-2.25)
    v4 := GenS1{F0: 0, F1: -51}
    v3 = 3.14159
    var v5 float32 = (float32(1.5) - (v3 * (((v3 / (v3 + v3)) + v3) * v3)))
    var v6 float64 = 1e10
    fmt.Println(v6, v6, v1)
    fmt.Println(v6, v1, v6)
    if (((v6 + v6) + v6) > ((v6 - ((0.0 + 0.1) / 3.14159)) / (v6 * float64(v3)))) {
        if (v1 > (v1 % ((679) | 1))) {
            fmt.Println(v5, v5, v5)
            v5 -= 0.1
        } else {
            fmt.Println(v3, v6, v5)
            var v7 int64 = 296
            _ = v7
        }
    } else {
        fmt.Println(v5, v3)
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            if (-1 >= 983) {
                if true {
                    v10 := []int{v1, fi8, v1}
                    var v11 int32 = -2147483648
                    fmt.Println(v5, fi9, v11)
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint16 = 255
                    _ = v12
                }
            }
            _ = fi9
        }
        _ = fi8
    }
    switch 783 {
    case 506:
        var v13 uint = 65535
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := []int{-2147483648}
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 2
    v2 += (v2 << ((255) & 63))
    v1 = (v1 << ((255) & 63))
    v17 := GenS1{F0: 1, F1: -1}
    v18 := GenS2{F0: 42, F1: (v16 ^ v16)}
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 -= (-(v1))
        fmt.Println(v5, v2, v2)
        _ = fi19
    }
    v20 := GenS3{F0: 127, F1: 127, F2: 378}
    v21 := []int{v1, -9223372036854775808, v1, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
