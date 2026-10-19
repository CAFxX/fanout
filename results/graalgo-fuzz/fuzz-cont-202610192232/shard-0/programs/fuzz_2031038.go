package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }

func main() {
    var v1 float64 = ((((1e10 + 1e10) * (1e10 - (0.1 * (float64((1.5 * 0.1)) * (0.1 - 1e10))))) + -2.25) / float64(1.5))
    var v2 float32 = 0.0
    var v3 int = 128
    v4 := make(map[int]int)
    var v5 int = v3
    var v6 float32 = float32(v1)
    for fi7 := 0; fi7 < 4; fi7++ {
        if ((v1 - (0.1 * 3.14159)) <= float64(v2)) {
            v1 = v1
        } else {
            v8 := []int{}
            fmt.Println(v2, v2)
            v1 = 1e10
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = -1
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    var v11 uint16 = 42
    var v12 uint32 = 2
    v13 := []int{v5}
    v14 := v4[v5]
    fmt.Println(v2, v6, v11)
    v2 += float32(float64(v2))
    for fi15 := 0; fi15 < 19; fi15++ {
        _ = len(v13)
        _ = fi15
    }
    if len(v13) > 0 {
        v13[0] = (v5 ^ (v5 % ((v5) | 1)))
    }
    var v16 uint16 = v11
    var v17 uint16 = (v11 % ((v16) | 1))
    v18 := GenS1{F0: -1, F1: "hello", F2: 62}
    v12 -= (v12 + (^(v12)))
    v19 := GenS2{F0: true}
    if (!(!false)) {
        v20 := GenS1{F0: -2147483648, F1: "a", F2: 895}
        v21 := GenS3{F0: v12, F1: 42, F2: 42}
        fmt.Println(v16, v2, v17)
        _ = v20
        _ = v21
    } else {
        var v22 int32 = 0
        v17 = (v11 + v11)
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
