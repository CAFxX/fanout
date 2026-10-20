package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64, p1 uint32) int8 {
    v6 := GenS1{F0: 546, F1: 0, F2: p1}
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{fi7, -1, fi7}
        _ = fi7
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int = 976
    var v2 int = (v1 | -1)
    var v3 uint = 1
    var v4 uint32 = 65535
    v5 := []int{(-(v2)), v1, v1}
    for fi9 := 0; fi9 < 2; fi9++ {
        v2 = (v2 % ((fi9) | 1))
        _ = fi9
    }
    if false {
        v10 := GenS2{F0: 216, F1: uint64(v3)}
        var v11 uint = (v3 ^ v3)
        fmt.Println(v2, v11, v1)
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: (uint16(v4) != 59), F1: 42}
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 int64 = -1
        _ = fi13
        _ = v14
    }
    var v15 int16 = 0
    v16 := GenS1{F0: v3, F1: 1, F2: (v4 | (v4 / ((uint32(v3)) | 1)))}
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint16 = 261
        v4 = v4
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 127, F1: 4294967295}
    _ = len(v5)
    switch (v2 - 32768) {
    case 4294967295:
        v20 := GenS4{F0: 987, F1: 199, F2: (-972 >= -1)}
        _ = v20
    case 355:
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int64 = 9223372036854775807
            _ = len(v5)
            switch "\n\t" {
            case "hello":
                var v23 int32 = 0
                _ = v23
            }
            _ = fi21
            _ = v22
        }
    default:
        v1 = (v2 / ((v1) | 1))
    }
    if (((((((float64(float32((float64(float32(1e10)) * (1.5 * 3.14159)))) + 3.14159) + 3.14159) - 3.14159) * 3.14159) / 1.5) / (3.14159 / 1.5)) != -2.25) {
        _ = len(v5)
        var v24 int16 = 32767
        _ = v24
    }
    v4 = v4
    var v25 int8 = -10
    v26 := GenS5{F0: -1}
    for fi27 := 0; fi27 < 4; fi27++ {
        v2 += -1
        fmt.Println(v15, fi27, v4)
        _ = fi27
    }
    v4 = (^(v4))
    v28 := GenS5{F0: v15}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
