package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 float32 = float32(1e10)
    var v4 int = 434
    var v5 bool = false
    var v6 string = "\n\t"
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 float64 = 1e10
            v9 := v2[42]
            v6 = "\n\t"
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v4)
    }
    v10 := GenS1{F0: 598, F1: 0, F2: 175}
    v4 = v4
    v3 -= -2.25
    var v11 uint32 = 42
    v12 := GenS1{F0: 526, F1: 18446744073709551615, F2: -961}
    var v13 uint32 = (^(v11))
    var v14 int16 = 1
    v4 = v4
    switch v6 {
    case "hello":
        var v15 int64 = 0
        _ = v15
    case "\n\t":
        var v16 float64 = 0.0
        _ = v16
    default:
        v3 += (float32((1e10 + (0.1 + 0.0))) / -2.25)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float64 = 1.5
        v19 := GenS1{F0: 580, F1: 1, F2: int32(v14)}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int64 = -1
    var v21 float32 = float32(float64((float32((1e10 * -2.25)) / v3)))
    for fi22 := 0; fi22 < 16; fi22++ {
        delete(v2, v11)
        _ = fi22
    }
    var v23 int32 = 42
    v21 = float32(1.5)
    if v5 {
        fmt.Println(v23, v4)
        v24 := v2[1]
        _ = v24
    } else {
        var v25 uint8 = 255
        v26 := GenS1{F0: 2, F1: 2, F2: (^(v23))}
        v27 := GenS2{F0: 128, F1: v4, F2: 650}
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
