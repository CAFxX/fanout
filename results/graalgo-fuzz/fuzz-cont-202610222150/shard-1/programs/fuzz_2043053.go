package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0() int32 {
    var v4 uint8 = 42
    v5 := GenS1{F0: 446, F1: 42}
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v4, fi6)
        _ = fi6
    }
    _ = v4
    _ = v5
    return 750
}
func genFn1(p0 uint64) int {
    v7 := GenS1{F0: (p0 ^ (^(p0))), F1: 42}
    switch 2147483647 {
    case 42:
        p0 += 1
    case -873:
        p0 -= 316
    }
    var v8 bool = false
    var v9 int16 = 175
    p0 -= p0
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v10 := []int{-606, v1, v1, (v1 ^ v1)}
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 bool = false
        if len(v10) > 0 {
            v10[0] = fi11
        }
        _ = fi11
        _ = v12
    }
    if (true && false) {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint = 2
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 -= v1
            switch "foo bar" {
            case "a":
                v1 = (^(v1))
            case "":
                var v16 float64 = float64((0.1 / (3.14159 - ((1e10 / 3.14159) * 1e10))))
                _ = v16
            }
            _ = fi15
        }
    }
    if len(v10) > 0 {
        v10[0] = v1
    }
    if len(v10) > 0 {
        v10[0] = genFn1(1)
    }
    if true {
        v1 -= v1
        var v17 bool = false
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v1 += -9223372036854775808
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS1{F0: 18446744073709551615, F1: -324}
            _ = fi19
            _ = v20
        }
        var v21 int = v1
        _ = fi18
        _ = v21
    }
    switch "" {
    case "foo bar":
        v1 = 958
    case "":
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint8 = 18
            v24 := GenS1{F0: 42, F1: 948}
            switch fi22 {
            case 170:
                delete(v2, 981)
            case 1:
                _ = len(v10)
            case 32768:
                var v25 int16 = -106
                _ = v25
            default:
                v23 = v23
            }
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        v26 := GenS1{F0: 90, F1: -1}
        _ = v26
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
