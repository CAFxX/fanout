package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }

func main() {
    v1 := []int{}
    var v2 bool = (1000 <= 32767)
    var v3 int = 127
    v4 := []int{(v3 | v3), v3}
    var v5 float32 = float32((float64(-2.25) * 1e10))
    v3 = v3
    switch "hello" {
    case "hello":
        v5 += -2.25
    case "foo bar":
        v2 = v2
    case "":
        v5 += (v5 / v5)
    }
    if v2 {
        switch "hello" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v2 = ((!(true && (-9223372036854775808 == -1))) && v2)
        }
        for fi6 := 0; fi6 < 11; fi6++ {
            v2 = (fi6 < fi6)
            v7 := GenS1{F0: "\n\t", F1: 2147483648}
            _ = fi6
            _ = v7
        }
        v8 := GenS2{F0: 42, F1: v2, F2: (-1 >= 52)}
        _ = v8
    }
    _ = len(v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS2{F0: 830, F1: v2, F2: (!true)}
        _ = fi9
        _ = v10
    }
    var v11 uint = 0
    var v12 int16 = 536
    if v2 {
        v12 += (v12 | (^(v12)))
        var v13 float32 = (v5 - 0.0)
        if false {
            if (v13 >= (v5 + (0.1 + 0.1))) {
                var v14 uint32 = 0
                v15 := GenS1{F0: "\n\t", F1: (v14 * v14)}
                _ = v14
                _ = v15
            } else {
                v5 = -2.25
                v5 = ((1.5 - v13) * (v13 / v13))
                var v16 uint64 = 4294967295
                _ = v16
            }
            v4 = append(v4, v3)
            switch (v3 % (((v3 - v3)) | 1)) {
            case 32768:
                v17 := GenS3{F0: (v12 % ((0) | 1))}
                _ = v17
            }
        } else {
            switch v3 {
            case 127:
                v5 -= v5
            default:
                v5 -= ((float32((0.1 * (float64((float32(1e10) + (1.5 - v13))) / 1.0))) / v5) / v5)
            }
            v18 := GenS4{F0: 254, F1: 2147483647, F2: 2147483647}
            _ = len(v4)
            _ = v18
        }
        _ = v13
    }
    v11 = (v11 / (((v11 % (((v11 >> ((255) & 63))) | 1))) | 1))
    if v2 {
        v19 := GenS2{F0: 0, F1: ((v11 / ((v11) | 1)) == v11), F2: (-81 >= 719)}
        fmt.Println(v3, v2, v11)
        v11 += 4294967295
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
