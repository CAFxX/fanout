package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint; F1 string }

func main() {
    var v1 int32 = 181
    var v2 uint16 = 42
    v3 := []uint32{968}
    var v4 float64 = 1.5
    v5 := GenS1{F0: 32768, F1: (v1 / ((v1) | 1))}
    v4 -= v4
    v6 := GenS1{F0: 722, F1: (v1 >> ((726) & 63))}
    v7 := GenS2{F0: 1000, F1: (-(v1))}
    v3 = append(v3, 2147483647)
    v3 = append(v3, 2)
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint = 4294967295
        if true {
            switch (fi9 % ((9223372036854775807) | 1)) {
            case 239:
                var v11 bool = (2 < 128)
                _ = v11
            case -134:
                v8 = (v8 & v8)
            case 621:
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := GenS1{F0: (fi12 * (fi9 | 621)), F1: 0}
                    _ = fi12
                    _ = v13
                }
            default:
                var v14 int64 = (^(v8))
                _ = v14
            }
            v15 := GenS1{F0: 467, F1: v1}
            v16 := GenS3{F0: false, F1: 18446744073709551615, F2: 127}
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi17
            }
            fmt.Println(v4)
            v4 = 0.1
        }
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = fi9
        _ = v10
    }
    var v18 bool = true
    v19 := GenS1{F0: 9223372036854775807, F1: v1}
    v20 := GenS4{F0: 737, F1: ("" + ("\n\t" + ("" + "hello")))}
    v21 := GenS2{F0: 1, F1: 1}
    var v22 uint64 = 18446744073709551615
    for fi23 := 0; fi23 < 19; fi23++ {
        v4 = 0.0
        v24 := GenS3{F0: true, F1: v22, F2: 255}
        _ = fi23
        _ = v24
    }
    var v25 uint = 255
    var v26 float32 = 1.5
    var v27 uint8 = uint8(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
