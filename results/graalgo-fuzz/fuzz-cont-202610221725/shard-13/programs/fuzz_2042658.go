package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
func genFn0(p0 uint8, p1 uint64) uint8 {
    fmt.Println(p0, p1)
    var v5 int64 = int64(p1)
    if (126 <= 0) {
        var v6 int = -769
        v7 := GenS1{F0: (v5 / ((v5) | 1))}
        _ = v6
        _ = v7
    } else {
        switch 217 {
        case 4294967295:
            v5 += (-(v5))
        default:
            v8 := []int{42, 4294967295, 62}
            _ = v8
        }
    }
    if (false && (65535 > 702)) {
        for fi9 := 0; fi9 < 10; fi9++ {
            v5 = (v5 >> ((348) & 63))
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 uint8, p1 uint64) uint16 {
    var v10 bool = (65535 != 4294967295)
    v11 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    var v3 int = -1
    v4 := make(map[int]string)
    switch (v3 + 32767) {
    case -2147483648:
        switch "a" {
        case "":
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 5; fi13++ {
                    if true {
                        var v14 uint = (v1 + 255)
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 925}
                        v3 -= v3
                        var v16 int32 = -1
                        _ = v15
                        _ = v16
                    }
                    v17 := []int{33, (^(v3))}
                    _ = fi13
                    _ = v17
                }
                v18 := []int{fi12, -205}
                v19 := GenS2{F0: 2, F1: 1, F2: 2147483647}
                _ = fi12
                _ = v18
                _ = v19
            }
        case "a":
            v3 += (v3 % (((v3 << ((807) & 63))) | 1))
        case "hello":
            v1 += (v1 & 65535)
        default:
            v20 := []int{-2147483648, (-(v3)), v3, (v3 | (v3 ^ 490))}
            _ = v20
        }
    case 476:
        if true {
            v3 += 829
            fmt.Println(v3)
        }
    default:
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(v3, v3)
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        delete(v2, 208)
        v3 += fi22
        _ = fi22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v1 = 2
        v4[v3] = "a"
        _ = fi23
    }
    v24 := GenS2{F0: 0, F1: 707, F2: (v3 & (v3 >> ((v1) & 63)))}
    delete(v4, v3)
    var v25 uint32 = 42
    v25 = (v25 % (((v25 % (((v25 % ((v25) | 1))) | 1))) | 1))
    fmt.Println(v25, v25)
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := []int{(^(v3)), 42}
        var v28 float64 = 0.1
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
