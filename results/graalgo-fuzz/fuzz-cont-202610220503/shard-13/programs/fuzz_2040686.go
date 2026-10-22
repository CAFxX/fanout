package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]int)
    var v3 int64 = -1
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 45, F1: 42, F2: 1}
    if (true || false) {
        v6 := GenS1{F0: 123, F1: 2, F2: 1}
        _ = v6
    }
    v7 := GenS2{F0: (32768 < 421), F1: 42, F2: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        if false {
            v9 := GenS3{F0: 255}
            v3 -= 222
            _ = v9
        }
        _ = fi8
    }
    v1 += v1
    if (!((^(v1)) != (v1 | v1))) {
        v10 := GenS2{F0: (0.1 == float64(1.5)), F1: 65535, F2: 127}
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint = 2
            var v13 int64 = (v3 & (v3 << (((v12 & 255)) & 63)))
            v14 := []int{(fi11 << ((v12) & 63)), int(v1)}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        v15 := GenS3{F0: 0}
        var v16 int32 = 0
        for fi17 := 0; fi17 < 18; fi17++ {
            v1 = (v1 ^ (v1 - (v1 / ((v1) | 1))))
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    v1 -= v1
    v3 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
