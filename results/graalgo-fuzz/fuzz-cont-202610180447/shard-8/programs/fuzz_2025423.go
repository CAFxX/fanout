package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64, p1 string) int64 {
    var v5 float32 = ((0.1 * 0.1) / 3.14159)
    var v6 int64 = 1
    if ((p0 | p0) <= p0) {
        v7 := []int{int(p0)}
        var v8 uint16 = 2
        var v9 int16 = 42
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint = 437
        v10 -= v10
        if (v5 <= ((v5 + -2.25) + (v5 - (0.0 * v5)))) {
            p0 += p0
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v6 * v6)
}

func main() {
    var v1 string = ""
    var v2 uint16 = 2
    v3 := make(map[int]bool)
    var v4 string = "a"
    var v11 bool = (0 < 4294967295)
    for fi12 := 0; fi12 < 3; fi12++ {
        v2 = 2
        v11 = v11
        v13 := []int{}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v2, v1)
        _ = fi14
    }
    switch "a" {
    case "\n\t":
        v2 += (^(v2))
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS1{F0: 1, F1: -541, F2: -63}
        switch (fi15 & fi15) {
        case 10:
            v17 := []int{fi15, fi15, (fi15 / ((fi15) | 1)), fi15}
            _ = v17
        case 128:
            fmt.Println(fi15)
        }
        _ = fi15
        _ = v16
    }
    if (128 <= 255) {
        if (5 < 42) {
            v18 := GenS2{F0: 323, F1: (v2 | v2)}
            _ = v18
        } else {
            v19 := GenS3{F0: uint8(v2)}
            _ = v19
        }
    }
    v2 += 662
    var v20 uint32 = 42
    v21 := GenS1{F0: 42, F1: 0, F2: -128}
    var v22 int8 = -108
    var v23 uint8 = uint8(v2)
    switch "\n\t" {
    case "\n\t":
        v23 = (v23 | 73)
    default:
        var v24 int32 = 890
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
