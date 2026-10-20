package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint16 = 526
    v2 := make(map[string]string)
    var v3 float64 = 1e10
    v4 := make(map[uint32]int)
    var v5 uint8 = 90
    if true {
        v6 := []int{9223372036854775807, -785, -127}
        _ = v6
    } else {
        var v7 uint16 = (v1 ^ 573)
        _ = v7
    }
    var v8 uint = 329
    if (1 < (^(v1))) {
        v1 = v1
        v4[493] = 11
    }
    v9 := GenS1{F0: v1, F1: (v5 >> (((v8 % (((^(v8))) | 1))) & 63)), F2: (v1 / (((^(v1))) | 1))}
    v3 = ((v3 + 1e10) * v3)
    v1 = 32767
    fmt.Println(v1, v5, v8)
    var v10 bool = true
    var v11 uint32 = 0
    v3 += v3
    var v12 int64 = 0
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: v1, F1: (v5 * v5), F2: (^(v1))}
        if (255 > uint16(v5)) {
            v15 := GenS2{F0: -746, F1: (v8 ^ (v8 << ((v8) & 63)))}
            fmt.Println(v5)
            _ = v15
        }
        fmt.Println(v10, v11)
        _ = fi13
        _ = v14
    }
    v1 += (^(v1))
    v3 = v3
    var v16 float64 = float64((((0.1 * 0.0) / ((0.0 - ((((1e10 / 0.1) / 3.14159) + 3.14159) * 0.1)) * ((1.5 - 1e10) + float32((v3 - v3))))) - 0.1))
    v12 += (v12 / (((v12 % ((v12) | 1))) | 1))
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS3{F0: (v12 == (^(v12)))}
        var v19 uint = (v8 + v8)
        switch "" {
        case "\n\t":
            if v10 {
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 uint = v8
                    var v22 uint16 = v1
                    var v23 uint16 = v22
                    if (!v10) {
                        v24 := GenS4{F0: 127}
                        var v25 float32 = 1e10
                        v19 = (^(v19))
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int16 = 42
                        _ = v26
                    }
                    v27 := GenS4{F0: -1}
                    v22 = v1
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v27
                }
            } else {
                v28 := GenS4{F0: 100}
                var v29 int64 = v12
                var v30 uint32 = uint32(v1)
                _ = v28
                _ = v29
                _ = v30
            }
        default:
            v31 := []int{(fi17 ^ 57)}
            _ = v31
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
