package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int; F2 int }
func genFn0(p0 uint8, p1 int) int16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (-(p1)) {
        case 4294967295:
            var v8 uint8 = 127
            _ = v8
        case -2147483648:
            fmt.Println(p1)
        }
        _ = fi7
    }
    p0 -= p0
    var v9 int = 9223372036854775807
    if (65535 == 442) {
        if (false || false) {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 float64 = 3.14159
                if (496 <= 1) {
                    v12 := GenS1{F0: 9223372036854775807, F1: 42}
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS2{F0: 2147483648, F1: 2147483647, F2: (p0 | 214)}
                        var v15 int32 = 0
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v12
                } else {
                    switch fi10 {
                    case 9223372036854775807:
                        v9 -= (fi10 + v9)
                    default:
                        var v16 string = "\n\t"
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 int32 = -1
                        _ = fi17
                        _ = v18
                    }
                }
                _ = fi10
                _ = v11
            }
            switch (p1 ^ 270) {
            case -2147483648:
                var v19 uint = 0
                _ = v19
            case 0:
                v20 := []int{0, p1, v9}
                _ = v20
            default:
                fmt.Println(p1)
            }
        }
        v21 := []int{(v9 | v9)}
        _ = v21
    }
    p1 += -168
    _ = p0
    _ = p1
    _ = v9
    return 165
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "\n\t"
    var v3 uint16 = 42
    var v4 int = 32768
    var v5 uint16 = v3
    var v6 uint16 = (v3 << ((0) & 63))
    fmt.Println(v3, v2)
    var v22 uint = 615
    if (1 != 1) {
        switch v4 {
        case 527:
            if (-32768 >= 581) {
                _ = len(v1)
                v23 := GenS2{F0: 1, F1: 2147483647, F2: 76}
                _ = v23
            }
        case -2147483648:
            switch v4 {
            case 0:
                v3 = v6
            case -9223372036854775808:
                v24 := GenS3{F0: ((v4 - 0) > v4)}
                _ = v24
            case 32768:
                v2 = "\n\t"
            }
        case 42:
            v3 -= 870
        }
        var v25 float64 = 0.1
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        fmt.Println(v22, v5, v3)
        fmt.Println(v5)
        _ = fi26
    }
    if ((v4 / ((1) | 1)) != v4) {
        v27 := []int{(v4 >> ((v22) & 63)), 128, v4}
        v6 = 70
        _ = v27
    }
    fmt.Println(v22, v6, v3)
    switch 9223372036854775807 {
    case -818:
        v22 = (v22 | v22)
    }
    v28 := GenS4{F0: (v5 * v5), F1: (v4 << ((v22) & 63)), F2: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
