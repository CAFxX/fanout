package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int8, p1 uint64, p2 uint16) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        p1 -= (^(p1))
        p1 = 1
        fmt.Println(p1, p2, fi5)
        _ = fi5
    }
    v6 := []int{-1, 267, -9223372036854775808, -783}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 128
}
func genFn1(p0 int) int {
    p0 += (p0 * (p0 % (((p0 + 42)) | 1)))
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 string = ""
        p0 = p0
        fmt.Println(v8, p0)
        _ = fi7
        _ = v8
    }
    p0 = (p0 * p0)
    var v9 string = ""
    _ = p0
    _ = v9
    return (p0 + (p0 >> ((255) & 63)))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    v3 := make(map[int]uint32)
    var v4 int = 1
    v10 := []int{v4, (v4 - genFn1((v4 % ((4294967295) | 1)))), 128, v4}
    v1[65535] = "\n\t"
    fmt.Println(v2, v4)
    for fi11 := 0; fi11 < 4; fi11++ {
        switch -1 {
        case 868:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        default:
            _ = len(v1)
        }
        fmt.Println(v2, v4)
        _ = fi11
    }
    v3[v4] = 2147483647
    v12 := GenS1{F0: v4}
    v13 := GenS2{F0: 42, F1: 614}
    v4 += v4
    if (float32(float64(v2)) >= ((3.14159 * v2) / v2)) {
        v4 = (v4 >> ((18446744073709551615) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
