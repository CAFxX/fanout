package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []uint32{0, 0, 4294967295}
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{}
    var v5 uint32 = 100
    v6 := make(map[int]uint32)
    var v7 float32 = 0.1
    var v8 float64 = (3.14159 * 1.5)
    var v9 int8 = 127
    for fi10 := 0; fi10 < 19; fi10++ {
        v5 -= v5
        v11 := v6[4294967295]
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if true {
            fmt.Println(v8)
        }
        _ = len(v6)
        _ = fi12
    }
    v1 = append(v1, (v5 / (((v5 & v5)) | 1)))
    v1 = append(v1, 2)
    _ = len(v1)
    v13 := GenS1{F0: -1}
    fmt.Println(v9, v3)
    var v14 uint8 = 19
    v15 := v6[32768]
    if true {
        v5 = v5
        switch "\n\t" {
        case "":
            if true {
                v15 = 786
                var v16 int32 = -463
                if (32767 >= 275) {
                    fmt.Println(v3)
                } else {
                    v9 += 1
                    var v17 bool = false
                    _ = v17
                }
                _ = v16
            }
        case "hello":
            v5 -= (^(v5))
        }
        var v18 uint32 = v5
        _ = v18
    }
    switch 127 {
    case 32768:
        var v19 int64 = -9223372036854775808
        _ = v19
    case 2147483647:
        if false {
            var v20 uint8 = uint8(v3)
            v21 := v6[-1]
            v22 := GenS2{F0: (v3 >= v3), F1: 18446744073709551615, F2: 4294967295}
            _ = v20
            _ = v21
            _ = v22
        }
    default:
        v6[32767] = v5
    }
    if ((v15 < (v5 | 42)) && (1.5 != float32((v8 + ((((v8 + float64(v7)) - ((v8 / v8) - float64(((float32(-2.25) * -2.25) * ((v7 - (v7 / 1.0)) * v7))))) * ((-2.25 * (((v8 + v8) - ((0.0 + (3.14159 * float64(1e10))) / 3.14159)) - 1.5)) * 0.0)) - 1.5))))) {
        if (true || (-244 >= 32767)) {
            var v23 uint8 = v14
            var v24 int16 = 0
            _ = v23
            _ = v24
        } else {
            v9 += (v9 ^ (v9 >> ((1) & 63)))
            v25 := GenS2{F0: true, F1: v3, F2: (v15 & v5)}
            _ = v25
        }
        v26 := GenS3{F0: v9, F1: 315}
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS4{F0: false, F1: (v14 + (v14 | (v14 & v14))), F2: 2147483647}
            var v29 float64 = v8
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v26
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v9 = int8(v14)
        _ = fi30
    }
    v31 := GenS5{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
