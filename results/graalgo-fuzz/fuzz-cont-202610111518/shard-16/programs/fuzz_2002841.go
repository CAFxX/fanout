package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float32 = 1e10
    v2 := []int{706, -1, 128, -308}
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{65535, 1}
    var v6 float32 = v1
    if false {
        var v7 float32 = v1
        var v8 int = 32767
        if true {
            v4[1] = 2147483648
            v9 := GenS1{F0: 42, F1: -1, F2: 536}
            fmt.Println(v8)
            _ = v9
        } else {
            v10 := v4[231]
            if (!(255 < 228)) {
                v11 := GenS2{F0: 21, F1: 1, F2: 127}
                _ = v11
            } else {
                var v12 string = "a"
                if (1 < 8) {
                    v8 = (v8 * 658)
                    v13 := GenS2{F0: 127, F1: 990, F2: -1}
                    v14 := GenS3{F0: 1}
                    _ = v13
                    _ = v14
                }
                _ = v12
            }
            if (v10 == (v10 - 492)) {
                v7 = v7
                v15 := GenS4{F0: 0, F1: -9223372036854775808}
                _ = v15
            }
            _ = v10
        }
        _ = v7
        _ = v8
    }
    delete(v4, 530)
    var v16 uint8 = 127
    v6 += (((v1 * 0.0) + 3.14159) * ((((-2.25 / v1) - v1) * v6) * 3.14159))
    var v17 float32 = v6
    v18 := GenS4{F0: -128, F1: 42}
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint = 498
        fmt.Println(v16)
        v21 := GenS3{F0: 0}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float32 = (1e10 * float32((1.5 + 0.1)))
    v22 -= 0.0
    var v23 uint8 = 0
    var v24 int16 = int16(v16)
    v25 := GenS5{F0: 829, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
