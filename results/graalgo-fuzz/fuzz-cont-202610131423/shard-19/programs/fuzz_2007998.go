package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64) uint8 {
    v5 := GenS1{F0: 2}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int8) int8 {
    switch ("foo bar" + "") {
    case "":
        if false {
            v6 := []int{32768, 9223372036854775807, 1}
            if false {
                var v7 float32 = (0.0 - float32(float64((((float32(1e10) - 1e10) + float32(float64(1e10))) - (0.1 + 0.0)))))
                _ = v7
            }
            _ = v6
        }
    case "a":
        v8 := []int{-9223372036854775808}
        _ = v8
    default:
        v9 := GenS1{F0: 42}
        _ = v9
    }
    if true {
        fmt.Println(p0)
        p0 = 42
    } else {
        p0 -= (p0 - (p0 ^ p0))
        fmt.Println(p0)
    }
    fmt.Println(p0)
    v10 := GenS2{F0: 65535, F1: 65535}
    _ = p0
    _ = v10
    return (-(p0))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 bool = (-30 >= 32767)
    var v4 uint32 = 0
    fmt.Println(v4, v4)
    if (-128 == 72) {
        switch 665 {
        case -2147483648:
            var v11 int = 346
            _ = v11
        case 32767:
            v12 := []int{676}
            _ = v12
        case 46:
            var v13 float64 = float64((1e10 + (1.5 / 1e10)))
            _ = v13
        }
    }
    var v14 uint8 = 1
    if (1e10 <= 1.5) {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int32 = 42
            v3 = v3
            var v17 uint32 = v4
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := v2[(v4 * (v4 ^ v4))]
        _ = v18
    }
    fmt.Println(v4, v14, v3)
    v4 = (v4 / (((^(v4))) | 1))
    if false {
        var v19 uint = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v14 += 0
        _ = fi20
    }
    v21 := []int{-9223372036854775808}
    for fi22 := 0; fi22 < 13; fi22++ {
        if ((v4 * 65535) < 2) {
            var v23 uint8 = 137
            _ = v23
        }
        _ = fi22
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
