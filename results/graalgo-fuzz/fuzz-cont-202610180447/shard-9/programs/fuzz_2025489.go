package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    v3 := []int{-2147483648, 42}
    var v4 uint16 = 65535
    v5 := make(map[uint32]uint32)
    var v6 string = "\n\t"
    v4 = 32767
    fmt.Println(v6)
    v4 = (^(v4))
    var v7 uint16 = (^(v4))
    var v8 int = 2147483647
    v6 = (v6 + (v6 + v6))
    v9 := GenS1{F0: (true || (!false))}
    v10 := v5[65535]
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 float32 = ((((float32(0.0) - -2.25) - 1.5) - 1e10) + 3.14159)
        _ = fi11
        _ = v12
    }
    v3 = append(v3, 1)
    _ = len(v5)
    v7 = (v4 / (((v4 % ((v4) | 1))) | 1))
    fmt.Println(v8)
    var v13 int8 = -128
    v4 -= v7
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint16 = 55
        switch (fi14 >> ((2) & 63)) {
        case 9223372036854775807:
            for fi16 := 0; fi16 < 2; fi16++ {
                if ((0.1 + 3.14159) != ((3.14159 - (3.14159 + ((float32(0.0) * 1e10) - 1e10))) - 0.0)) {
                    v17 := GenS2{F0: 1000, F1: false, F2: v10}
                    _ = v17
                } else {
                    var v18 uint8 = 87
                    v18 = (v18 >> ((uint(v10)) & 63))
                    _ = v18
                }
                v7 = v7
                _ = fi16
            }
        case -991:
            var v19 int64 = -513
            _ = v19
        default:
            var v20 uint16 = v4
            _ = v20
        }
        v10 += 485
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
