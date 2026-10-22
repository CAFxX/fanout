package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]int)
    var v3 float64 = float64(((float32(((1e10 - (3.14159 - 1e10)) / (1e10 * 3.14159))) * (float32(0.1) * ((0.1 / float32(-2.25)) - 3.14159))) - 1e10))
    var v4 bool = (v1 < v1)
    v5 := []string{"foo bar"}
    var v6 uint64 = 42
    delete(v2, (^(v1)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 -= (^(v6))
        _ = fi7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: 129, F1: 1, F2: 1}
        var v10 int32 = int32(fi8)
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS2{F0: 42}
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v13 := GenS3{F0: -1}
    switch -1 {
    case 524:
        v14 := GenS1{F0: uint8(v1), F1: 42, F2: 128}
        _ = v14
    }
    var v15 uint = 891
    var v16 uint16 = 42
    var v17 int64 = -26
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
