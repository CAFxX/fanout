package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int8; F2 int }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]uint32)
    var v3 int16 = -720
    v4 := make(map[uint32]int)
    var v5 uint32 = 0
    var v6 uint = 165
    if (false && true) {
        for fi7 := 0; fi7 < 12; fi7++ {
            _ = len(v4)
            v3 -= (v3 + 706)
            _ = fi7
        }
    }
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 14; fi9++ {
        delete(v2, (fi9 + fi9))
        _ = fi9
    }
    if (!false) {
        fmt.Println(v3, v5, v6)
    } else {
        var v10 uint8 = 1
        v11 := GenS2{F0: 1, F1: 23, F2: "a"}
        v12 := GenS1{F0: 9223372036854775807}
        _ = v10
        _ = v11
        _ = v12
    }
    v1 += (v3 / (((v1 * 32767)) | 1))
    switch 32767 {
    case -459:
        var v13 int32 = -1
        _ = v13
    default:
        v1 = (v3 - -282)
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v1 += 478
        var v15 float64 = -2.25
        v15 = 0.0
        _ = fi14
        _ = v15
    }
    v5 += v5
    fmt.Println(v5)
    v1 = (v3 << (((v6 + (v6 >> (((v6 & (v6 << (((v6 << (((^(v6))) & 63))) & 63)))) & 63)))) & 63))
    if true {
        v5 -= v5
        var v16 string = "hello"
        var v17 float32 = ((0.1 + ((((3.14159 + ((-2.25 + 3.14159) * 1e10)) + (3.14159 * (3.14159 + ((0.1 * (0.1 * 1.5)) + 0.0)))) - ((float32(float64(0.1)) * ((-2.25 - (0.0 + 1e10)) + 3.14159)) - 0.1)) / ((0.0 - ((-2.25 * 0.1) - 1.5)) / ((0.0 - 1e10) / 0.1)))) + 0.0)
        _ = v16
        _ = v17
    } else {
        v18 := []int{-2147483648, 42, 128}
        v3 -= (v3 << ((v6) & 63))
        _ = v18
    }
    v19 := GenS3{F0: 1}
    switch 32768 {
    case 1:
        fmt.Println(v5, v3)
    case -2147483648:
        v20 := GenS4{F0: v6, F1: -128, F2: 597}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
