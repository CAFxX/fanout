package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint }
func genFn0(p0 uint, p1 uint64) int16 {
    v5 := GenS1{F0: p1, F1: -2147483648, F2: 177}
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{(fi6 ^ (fi6 >> ((p0) & 63))), 128, 0}
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "" {
            case "\n\t":
                v9 := GenS2{F0: 66, F1: -1, F2: 494}
                _ = v9
            case "a":
                if len(v7) > 0 {
                    v7[0] = fi8
                }
            case "":
                fmt.Println(p1)
            }
            fmt.Println(p0)
            var v10 bool = false
            _ = fi8
            _ = v10
        }
        var v11 string = "a"
        _ = fi6
        _ = v11
        _ = v7
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        switch 942 {
        case 32767:
            fmt.Println(p0)
        case 872:
            if true {
                v13 := GenS2{F0: 1, F1: -222, F2: p0}
                _ = v13
            }
        default:
            p0 = p0
        }
        var v14 float64 = ((float64(float32((-2.25 - 0.1))) - -2.25) + 0.1)
        p0 -= p0
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v5
    return 1000
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    v4 := []string{"a"}
    _ = len(v1)
    delete(v1, 175)
    v15 := GenS1{F0: 0, F1: v3, F2: 247}
    v3 = v3
    v3 = (v3 + v3)
    switch "\n\t" {
    case "hello":
        fmt.Println(v3)
    default:
        v16 := GenS2{F0: 127, F1: 1, F2: 574}
        _ = v16
    }
    v3 = (^(v3))
    switch "foo bar" {
    case "a":
        v3 -= v3
    case "foo bar":
        v17 := GenS2{F0: 100, F1: 0, F2: 4294967295}
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
