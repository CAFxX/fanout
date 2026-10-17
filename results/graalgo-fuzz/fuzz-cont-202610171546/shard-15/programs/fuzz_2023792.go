package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    var v4 int8 = -59
    v4 = -1
    v5 := []int{32767}
    v6 := GenS1{F0: 0, F1: false, F2: (v1 >> ((18446744073709551615) & 63))}
    var v7 uint32 = 670
    var v8 uint16 = 2
    var v9 bool = (1e10 >= (1.5 - 1e10))
    v3[v7] = (v7 > v7)
    v10 := GenS2{F0: (v1 | (v1 + (v1 % (((v1 % ((v1) | 1))) | 1)))), F1: false}
    v11 := GenS3{F0: 32767}
    v8 -= (^(v8))
    var v12 uint = 0
    switch 235 {
    case 32767:
        var v13 uint8 = 2
        _ = v13
    case 42:
        v2 = ((0.0 + v2) + (1e10 + (-2.25 / (float32(3.14159) + v2))))
    case 0:
        _ = len(v5)
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        switch (-(fi14)) {
        case 4294967295:
            v4 += (v4 % ((-1) | 1))
        case 32768:
            for fi15 := 0; fi15 < 11; fi15++ {
                _ = len(v5)
                v16 := GenS3{F0: -987}
                _ = fi15
                _ = v16
            }
        default:
            v17 := GenS4{F0: (v1 >> ((v12) & 63)), F1: ((!v9) && (v9 || false))}
            _ = v17
        }
        _ = fi14
    }
    if (42 > 685) {
        fmt.Println(v9, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
