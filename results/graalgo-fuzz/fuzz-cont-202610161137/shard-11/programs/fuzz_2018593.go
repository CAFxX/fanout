package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
func genFn0() int64 {
    v5 := []int{-1, -1}
    var v6 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 int8 = -1
            v8 = (-(v8))
            v6 -= 96
            _ = v8
        }
        _ = fi7
    }
    if len(v5) > 0 {
        v5[0] = -1
    }
    v9 := GenS1{F0: 0, F1: (!(0 <= 428))}
    _ = v5
    _ = v6
    _ = v9
    return 552
}

func main() {
    var v1 float64 = float64((((((float32((0.1 * (1.5 / (float64((float32(1.5) / 1.5)) * (-2.25 / (0.1 + (((1.5 * ((1.5 - 0.0) / 1.0)) / -2.25) * (float64((float32(1.5) - 0.0)) / 1e10)))))))) * 1.5) * 3.14159) - (1.5 / ((1.5 - 3.14159) * -2.25))) - 3.14159) * 1e10))
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v10 uint16 = 65535
    fmt.Println(v10)
    v11 := GenS1{F0: 9223372036854775807, F1: (42 <= genFn0())}
    v12 := GenS2{F0: 1}
    var v13 int32 = 0
    var v14 int16 = 143
    v15 := GenS3{F0: int64(v14)}
    v10 -= (v10 - 32767)
    var v16 uint32 = 143
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v14)
        var v18 uint8 = 1
        v19 := GenS4{F0: int64(v10), F1: 18446744073709551615, F2: 235}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := []int{97, 0, -122}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
