package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }

func main() {
    var v1 bool = false
    var v2 uint64 = 671
    var v3 int32 = 1
    var v4 uint32 = 562
    v2 -= 18446744073709551615
    v5 := GenS1{F0: true, F1: 2}
    v6 := GenS1{F0: (v1 || true), F1: 2}
    v3 += (-(v3))
    v7 := GenS2{F0: "", F1: v3, F2: ((v1 || ("foo bar" != "hello")) && v1)}
    if (false && false) {
        v8 := GenS3{F0: "", F1: "", F2: (v2 * 1)}
        _ = v8
    }
    v9 := GenS4{F0: v3, F1: v2, F2: 3}
    v1 = true
    var v10 uint8 = 134
    switch 42 {
    case 0:
        v10 = 190
    default:
        fmt.Println(v1)
    }
    v4 = v4
    var v11 bool = v1
    var v12 uint16 = 42
    var v13 float32 = (float32(float64(((float32(((1e10 * 0.1) / 1.0)) / 0.1) / (1.5 - (1.5 + 1e10))))) + 1.5)
    for fi14 := 0; fi14 < 7; fi14++ {
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS4{F0: (v3 + (v3 + (v3 >> ((0) & 63)))), F1: v2, F2: (^(v10))}
                v17 := []int{42}
                _ = fi15
                _ = v16
                _ = v17
            }
            switch fi14 {
            case -1:
                var v18 int8 = 100
                _ = v18
            default:
                v1 = (!false)
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                switch fi19 {
                case 128:
                    var v20 int64 = int64(v4)
                    _ = v20
                default:
                    switch fi19 {
                    case -2147483648:
                        v21 := []int{}
                        _ = v21
                    case 9223372036854775807:
                        fmt.Println(v12, v4, v3)
                    }
                }
                v22 := []int{-1, fi19, 2147483647, 2147483647}
                var v23 bool = true
                var v24 int64 = 1
                _ = fi19
                _ = v22
                _ = v23
                _ = v24
            }
        }
        _ = fi14
    }
    var v25 uint8 = 0
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
