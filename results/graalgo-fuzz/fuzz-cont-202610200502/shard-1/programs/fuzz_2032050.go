package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{45, 42, 2147483648}
    var v2 bool = true
    var v3 float64 = 0.1
    var v4 string = ""
    v3 -= (v3 * (0.1 * v3))
    v5 := GenS1{F0: -764, F1: 0}
    var v6 uint16 = 32767
    v2 = true
    var v7 uint8 = 42
    if (false || true) {
        v8 := GenS2{F0: (v6 * v6)}
        _ = v8
    }
    var v9 int = -9223372036854775808
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS2{F0: (v6 | v6)}
        _ = fi10
        _ = v11
    }
    v7 -= 1
    for fi12 := 0; fi12 < 8; fi12++ {
        switch (v9 / ((1) | 1)) {
        case -2147483648:
            v13 := GenS3{F0: 305}
            _ = v13
        default:
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS2{F0: v6}
                _ = len(v1)
                fmt.Println(v7)
                _ = fi14
                _ = v15
            }
        }
        _ = fi12
    }
    v16 := GenS1{F0: v9, F1: 2}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v7 = 175
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v2 = v2
            _ = fi18
        }
        _ = fi17
    }
    v19 := GenS2{F0: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
