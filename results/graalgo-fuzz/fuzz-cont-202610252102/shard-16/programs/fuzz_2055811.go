package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint64, p1 uint16) uint16 {
    v6 := []int{42, 46}
    var v7 uint8 = 127
    p0 += (p0 * 571)
    for fi8 := 0; fi8 < 5; fi8++ {
        switch ("a" + "foo bar") {
        case "hello":
            v9 := GenS1{F0: (^(p0)), F1: v7}
            _ = v9
        default:
            v7 += 128
        }
        fmt.Println(p1)
        if true {
            if true {
                v10 := GenS1{F0: 0, F1: v7}
                _ = v10
            }
        }
        _ = fi8
    }
    fmt.Println(v7, p1, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    v3 := []int{-882, 128, 32767}
    v4 := make(map[uint32]int)
    var v5 float64 = (1e10 * 1e10)
    fmt.Println(v2, v5)
    delete(v1, 748)
    for fi11 := 0; fi11 < 14; fi11++ {
        v5 += float64(0.1)
        _ = fi11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint8 = 125
        if len(v3) > 0 {
            v3[0] = fi12
        }
        _ = fi12
        _ = v13
    }
    var v14 int16 = -1
    var v15 float64 = (v5 * float64(0.1))
    if false {
        v16 := GenS1{F0: 18446744073709551615, F1: 42}
        var v17 uint64 = 214
        _ = v16
        _ = v17
    } else {
        if false {
            switch "a" {
            case "a":
                v15 += (float64(3.14159) - 0.1)
            default:
                v18 := GenS2{F0: 18446744073709551615}
                _ = v18
            }
        }
        v2 = 86
        v5 = float64((((0.0 / -2.25) / ((((0.1 - 1e10) + 3.14159) / 1.5) / -2.25)) / 1.5))
    }
    v14 -= 524
    _ = len(v3)
    var v19 int8 = v2
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS3{F0: 4294967295, F1: -1}
        fmt.Println(v15, v2, v19)
        v2 -= -31
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
